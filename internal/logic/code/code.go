package code

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
	"time"
	codeV1 "xcoder/api/code/v1"
	"xcoder/core/llms"
	xCoderVLLM "xcoder/core/llms/vllm/xcoder"
	prompts "xcoder/core/prompts/code"
	"xcoder/core/retrieval"
	"xcoder/internal/consts"
	"xcoder/internal/controller/utils"
	"xcoder/internal/controller/utils/ccode"
	"xcoder/internal/controller/utils/cerror"
	"xcoder/internal/dao"
	"xcoder/internal/logic/common"
	"xcoder/internal/model/entity"
	"xcoder/internal/model/input/codein"
	"xcoder/internal/service"
	"xcoder/utility/xapollo"
)

type sCode struct{}

func NewCode() service.ICode {
	return &sCode{}
}

func init() {
	service.RegisterCode(NewCode())
}

// 仅 vscode 检查最后一个字符是否为换行符，如果是，则去除
func generatePostprocessRemoveNewLine(ctx context.Context, generateUUID string, completion string) string {
	g.Log().Info(ctx, "GenerateUUID: %s generate completion: %s", generateUUID, completion)
	if len(completion) > 1 && completion[len(completion)-1:] == "\n" {
		completion = completion[:len(completion)-1]
		return completion
	}
	return completion
}

// 仅 vscode 去除生成的代码与 suffixCodeFirstLine 相同的子串
func generatePostprocessRemoveDuplicatedStr(ctx context.Context, generateUUID string,
	completion string, suffixCodeFirstLine string) string {
	g.Log().Info(ctx, "GenerateUUID: %s generate completion: %s, suffixCode first line: %s",
		generateUUID, completion, suffixCodeFirstLine)

	completion = strings.TrimSpace(completion)
	suffixCodeFirstLine = strings.TrimSpace(suffixCodeFirstLine)
	// 去除生成的代码与 suffixCodeFirstLine 相同的子串
	completion = strings.Replace(completion, suffixCodeFirstLine, "", -1)

	return completion
}

func generatePostprocess(ctx context.Context, in *codein.CodeGeneratePostprocessRequest) (string, error) {
	generateUUID := in.GenerateUUID
	completion := in.Completion
	suffixCodeFirstLine := strings.Split(in.CodeAfterCursor, "\n")[0]

	// vscode 且单行生成，额外处理
	if in.IsSingleLine && strings.Contains(in.IDEInfo, "vscode") {
		// 处理单行生成时，检查生成的符号是否与 suffixCodeFirstLine 相同，如果相同，则换成生成空串
		completion = generatePostprocessRemoveDuplicatedStr(ctx, generateUUID, completion, suffixCodeFirstLine)
		// 检查生成的代码最后一个字符是否与 suffixCodeFirstLine 最后一个字符相同，且在 [')', ']', '}', "'", '"'] 中，
		// 如果不是，则加上该字符
		//completion = generatePostprocessAddCharacter(ctx, generateUUID, completion, suffixCodeFirstLine)
		// 处理单行生成时，检查最后一个字符是否为换行符，如果是，则去除
		completion = generatePostprocessRemoveNewLine(ctx, in.GenerateUUID, completion)
	}

	// 如果生成为空，或者生成的符号为<｜fim end｜>，则取消生成，记录为取消状态
	if len(completion) == 0 || completion == "<｜fim▁end｜>" {
		g.Log().Infof(ctx, "GenerateUUID: %s generate completion is empty", generateUUID)
		err := dao.PCodeGenRecords.UpdateByGenerateUuid(ctx, generateUUID, map[string]interface{}{
			"accept_status": consts.CodeAcceptStatusCancelled.String(),
		})
		if err != nil {
			errMsg := fmt.Sprintf("GenerateUUID: %s PCodeGenRecordsUpdateByUUID failed: %v", generateUUID, err)
			g.Log().Errorf(ctx, errMsg)
			return "", cerror.NewInternalServerError(err, errMsg)
		}
		return "", cerror.NewUserErrorButHttpOk(fmt.Errorf("GenerateUUID: %s gerenate completion is empty", generateUUID),
			ccode.CodeGenerateCompletionEmpty, "")
	}

	return completion, nil
}

func generateWithVLLM(ctx context.Context, in *codein.CodeGenerateWithLLMRequest) (
	*codein.CodeGenerateWithLLMResponse, error) {
	baseUrl := utils.RandSliceValue(in.ConnUrls)
	llm, err := xCoderVLLM.New(
		xCoderVLLM.WithModel(in.ModelVersion),
		xCoderVLLM.WithBaseURL(baseUrl),
	)
	if err != nil {
		g.Log().Errorf(ctx, "Model: %s, xCoderVLLM.New failed: %v", in.ModelVersion, err)
		return nil, err
	}

	prompt, err := prompts.GenerateCodeLLamaPrompt(in.CodeBeforeCursor, in.CodeAfterCursor)
	if err != nil {
		g.Log().Errorf(ctx, "GenerateUUID: %s prompts.GeneratePrompt failed: %v", in.GenerateUUID, err)
		return nil, cerror.NewUserError(err, ccode.CodeGeneratePromptError,
			ccode.CodeGeneratePromptErrorMessage)
	}
	if xapollo.GetSelectedModel() == "deepseeker" {
		prompt, err = prompts.GenerateDeepSeekerPrompt(in.CodeBeforeCursor, in.CodeAfterCursor)
		if err != nil {
			g.Log().Errorf(ctx, "GenerateUUID: %s prompts.GeneratePrompt failed: %v", in.GenerateUUID, err)
			return nil, cerror.NewUserError(err, ccode.CodeGeneratePromptError,
				ccode.CodeGeneratePromptErrorMessage)
		}
	}

	completion, err := llm.Predict(ctx, prompt,
		llms.WithUUID(in.GenerateUUID),
		llms.WithModel(in.ModelVersion),
		llms.WithMaxTokens(in.MaxTokens),
		llms.WithTemperature(in.Temperature),
		llms.WithTopP(in.TopP),
		llms.WithStopWords(in.StopWords),
		llms.WithIsMultiLine(!in.IsSingleLine),
		llms.WithSuffixCode(in.CodeAfterCursor),
		llms.WithConnUrls(in.ConnUrls),
	)
	if err != nil {
		g.Log().Errorf(ctx, "GenerateUUID: %s llm.Predict failed: %v", in.GenerateUUID, err)
		return nil, cerror.NewUserError(err, ccode.CodeGeneratePredictError,
			ccode.CodeGeneratePredictErrorMessage)
	}

	return &codein.CodeGenerateWithLLMResponse{
		Completion: completion,
	}, nil
}

func (s *sCode) Generate(ctx context.Context, in *codein.CodeGenerateRequest) (res *codeV1.CodeGenerateRes, err error) {
	// 获取代码模型参数
	llmParams, err := xapollo.GetCodeGenerateLLmParams()
	if err != nil {
		return nil, cerror.NewUserError(err, ccode.CodeUserOperatorError, "")
	}

	// 从本地配置文件获取 retrieval cfc params
	retrievalCrossFileCtxParams, err := xapollo.GetRetrievalCrossFileContextParams()
	if err != nil {
		return nil, cerror.NewUserError(err, ccode.CodeUserOperatorError, "")
	}

	// 组建参数
	isSingleLine := true
	generateType := in.GenerateType
	topP := llmParams.SingleLineTopP
	stopWords := llmParams.SingleLineStopWords
	maxTokens := llmParams.SingleLineMaxTokens
	temperature := llmParams.SingeLineTemperature
	if utils.StringIsInSmallSlice(generateType, consts.MultiLineGenerateType) {
		isSingleLine = false
		topP = llmParams.MultiLineTopP
		stopWords = llmParams.MultiLineStopWords
		maxTokens = llmParams.MultiLineMaxTokens
		temperature = llmParams.MultiLineTemperature
	}

	codeBeforeCursor := in.CodeBeforeCursor
	codeAfterCursor := in.CodeAfterCursor
	codeBeforeWithContext := in.CodeBeforeCursor

	totalMaxTokens := llmParams.TotalMaxTokens
	if (len([]rune(codeBeforeCursor)) + len([]rune(codeAfterCursor))) > totalMaxTokens*4 {
		if len(codeBeforeCursor) >= gconv.Int(float64(totalMaxTokens)*0.85*4) {
			// codeBeforeCursor 太长，需要从下往上截断，将上面多余部分丢掉
			codeBeforeCursor = codeBeforeCursor[len(codeBeforeCursor)-gconv.Int(float64(totalMaxTokens)*0.85*4):]
			// 去掉第一行
			codeBeforeCursorLines := strings.Split(codeBeforeCursor, "\n")
			if len(codeBeforeCursorLines) > 0 {
				codeBeforeCursor = strings.Join(codeBeforeCursorLines[1:], "\n")
			}
		}

		if len(codeAfterCursor) >= gconv.Int(float64(totalMaxTokens)*0.15*4) {
			// codeAfterCursor 太长，需要从上往下截断，将下面多余部分丢掉
			codeAfterCursor = codeAfterCursor[:gconv.Int(float64(totalMaxTokens)*0.15*4)]
			// 去掉最后一行
			codeAfterCursorLines := strings.Split(codeAfterCursor, "\n")
			if len(codeAfterCursorLines) > 0 {
				codeAfterCursor = strings.Join(codeAfterCursorLines[:len(codeAfterCursorLines)-1], "\n")
			}
		}
	}

	codeBeforeCursorLines := strings.Split(codeBeforeCursor, "\n")
	st := time.Now()
	cfcNums := 0
	chunkSize := retrievalCrossFileCtxParams.ChunkSize
	supportCodeLangs := retrievalCrossFileCtxParams.SupportCodeLangs
	contextsNum := len(in.Contexts)

	// 如果 contexts 数量大于 maxCrossFileNums 则需要使用前 maxCrossFileNums 个 context
	if contextsNum > retrievalCrossFileCtxParams.MaxCrossFileNums {
		in.Contexts = in.Contexts[:retrievalCrossFileCtxParams.MaxCrossFileNums]
	}

	// 使用 cfc 的条件：
	// 1. 需要有 context 信息
	// 2. 当前编辑的代码语言支持
	// 3. codeBeforeCursor 行数大于 chunkSize
	codeGenSnippetModels := make([]*entity.PCodeGenRetrievalSnippetMap, 0)
	if contextsNum > 0 && utils.StringIsInSmallSlice(strings.ToLower(in.CodeLanguage),
		supportCodeLangs) && len(codeBeforeCursorLines) > chunkSize {
		slideWindowSize := retrievalCrossFileCtxParams.SlideWindowSize
		minScore := retrievalCrossFileCtxParams.MinScore

		// 取 codeBeforeCursor 最后10行作为queryDoc
		queryDoc := strings.Join(codeBeforeCursorLines[len(codeBeforeCursorLines)-chunkSize:], "\n")
		reg := retrieval.NewMatcher(in.CodeLanguage, in.CodePath, chunkSize, slideWindowSize, minScore, queryDoc)
		snippets, err := reg.RetrieveAllSnippets(ctx, in.Contexts)
		if err != nil {
			g.Log().Errorf(ctx, "failed to retrieve snippets: %v", err)
		} else {
			cfcNums = len(snippets)

			ctxContent := ""
			ctxTitleFormatter := "Compare this snippet from %s:"
			for _, item := range snippets {
				g.Log().Infof(ctx, "item.Path: %s, item.Snippet Score: %f", item.Path, item.Score)
				ctxTitle := fmt.Sprintf(ctxTitleFormatter, item.Path)
				ctxContent += fmt.Sprintf("%s\n%s\n\n", ctxTitle, item.Snippet)

				codeGenSnippetModels = append(codeGenSnippetModels, &entity.PCodeGenRetrievalSnippetMap{
					Isactive:       1,
					DeleteToken:    "NA",
					Inserttime:     gtime.Now(),
					Updatetime:     gtime.Now(),
					GenerateUuid:   in.GenerateUUID,
					GitRepo:        in.GitRepo,
					GitBranch:      in.GitBranch,
					SnippetUuid:    utils.GenerateMD5(item.Snippet),
					SnippetRepo:    in.GitRepo,
					SnippetPath:    item.Path,
					SnippetScore:   item.Score,
					SnippetContent: item.Snippet,
					ProjectName:    consts.ProjectName,
					ProjectVersion: consts.ProjectVersion,
					CreateUser:     in.CreateUser,
				})
			}

			// 将 ctxContent 注释
			annotCode, err := utils.AnnotatingCodeWithLanguage(ctx,
				&utils.AnnotatingCodeRequest{
					Language: in.CodeLanguage,
					Code:     ctxContent,
				})
			if err != nil {
				g.Log().Errorf(ctx, "failed to annotating code with language: %v", err)
				return nil, err
			}

			annotCode.Code = fmt.Sprintf("// Path: %s\n%s", in.CodePath, annotCode.Code)
			codeBeforeCursor = fmt.Sprintf("%s%s", annotCode.Code, codeBeforeCursor)
			codeBeforeWithContext = codeBeforeCursor
		}
	}
	g.Log().Infof(ctx, "=== retrieve cost: %v ===", time.Since(st))

	// 插入数据库
	if err = common.CodeGenerateRecordInsertDB(ctx, &codein.CodeGenerateRecordInsertDBRequest{
		Input:                 in,
		LLMParams:             llmParams,
		GenerateType:          generateType,
		IsSingleLine:          isSingleLine,
		CfcNums:               cfcNums,
		CodeBeforeCursor:      codeBeforeCursor,
		CodeAfterCursor:       codeAfterCursor,
		CodeBeforeWithContext: codeBeforeWithContext,
		CodeGenSnippetModels:  codeGenSnippetModels,
	}); err != nil {
		return nil, err
	}

	// 开始调用模型，生成代码
	var completion string
	req := &codein.CodeGenerateWithLLMRequest{
		GenerateUUID:     in.GenerateUUID,
		CodeBeforeCursor: codeBeforeCursor,
		CodeAfterCursor:  codeAfterCursor,
		ModelVersion:     llmParams.ModelVersion,
		MaxTokens:        maxTokens,
		Temperature:      temperature,
		TopP:             topP,
		StopWords:        stopWords,
		IsSingleLine:     isSingleLine,
		ConnUrls:         llmParams.ConnUrls,
	}
	result, err := generateWithVLLM(ctx, req)
	if err != nil {
		g.Log().Errorf(ctx, "generateWithCodeLlama failed: %v", err)
		return nil, err
	}
	completion = result.Completion

	// 生成代码后处理
	generatePostprocessReq := &codein.CodeGeneratePostprocessRequest{
		GenerateUUID:    in.GenerateUUID,
		CodeAfterCursor: codeAfterCursor,
		IDEInfo:         in.IDEInfo,
		IsSingleLine:    isSingleLine,
		GenerateType:    in.GenerateType,
		Completion:      completion,
	}
	completion, err = generatePostprocess(ctx, generatePostprocessReq)
	if err != nil {
		return nil, err
	}

	return &codeV1.CodeGenerateRes{
		Code: completion,
	}, nil
}

func (s *sCode) UpdateGenerationInfo(ctx context.Context, in *codein.CodeUpdateGenerationInfoReq) (
	out *codein.CodeUpdateGenerationInfoRes, err error) {
	_, err = dao.PCodeGenRecords.GetOneByGenerateUuid(ctx, in.GenerateUUID)
	if err != nil {
		g.Log().Errorf(ctx, "failed to get code gen records by generateUUID: %s, error: %v",
			in.GenerateUUID, err)
		return nil, err
	}

	updatedFields := g.Map{
		"prompt_tokens":          in.PromptTokens,
		"completion_code":        in.CompletionCode,
		"completion_code_tokens": in.CompletionCodeTokens,
		"completion_code_lines":  in.CompletionCodeLines,
		"completion_duration":    in.CompletionDuration,
		"finish_reason":          in.FinishReason,
	}

	err = dao.PCodeGenRecords.UpdateByGenerateUuid(ctx, in.GenerateUUID, updatedFields)
	if err != nil {
		g.Log().Errorf(ctx, "failed to update code gen records by generateUUID: %s, error: %v",
			in.GenerateUUID, err)
		return nil, err
	}

	return nil, nil
}
