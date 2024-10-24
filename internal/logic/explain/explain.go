package explain

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"xcoder/core/prompts/explain"
	"xcoder/core/schema"
	"xcoder/internal/consts"
	"xcoder/internal/controller/utils/ccode"
	"xcoder/internal/controller/utils/cerror"
	"xcoder/internal/logic/common"
	"xcoder/internal/model/input/chatin"
	"xcoder/internal/model/input/explainin"
	"xcoder/internal/service"
	"xcoder/utility/xutils"
)

type sExplain struct{}

func NewExplain() service.IExplain {
	return &sExplain{}
}

func init() {
	service.RegisterExplain(NewExplain())
}

func (s *sExplain) SseGenerate(ctx context.Context, in *explainin.ExplainSseGenerateRequest, out chan *chatin.ChatResult) error {
	// 将 code language 转换为小写，并将 go 转换为 golang
	codeLanguage := xutils.CodeLanguageToLower(in.CodeLanguage)

	// 组建 prompts
	prompts, err := explain.GenerateExplainPrompt(ctx,
		&explainin.CodeExplainPromptRequest{
			RepoName:     in.GitRepo,
			CodePath:     in.CodePath,
			CodeLanguage: codeLanguage,
			SelectedCode: in.UserCode,
		})
	if err != nil {
		g.Log().Errorf(ctx, "ConversationUUID: %s, explain.GenerateExplainPrompt failed: %v",
			in.ConversationUUID, err)
		return cerror.NewUserError(err, ccode.CodeGeneratePromptError,
			ccode.CodeGeneratePromptErrorMessage)
	}

	var messages []schema.ChatMessageRequest
	err = gconv.Struct(prompts, &messages)
	if err != nil {
		g.Log().Errorf(ctx, "ConversationUUID: %s, gconv.Struct failed: %v", in.ConversationUUID, err)
		return err
	}

	// 调用 Chat LLM 服务
	chatInput := &chatin.ChatSseGenerateReq{}
	if err = gconv.Struct(in, chatInput); err != nil {
		g.Log().Errorf(ctx, "gconv.Struct failed: %v", err)
		return err
	}
	var codeContexts []chatin.Context
	if err = gconv.SliceStruct(in.UserContext, &codeContexts); err != nil {
		g.Log().Errorf(ctx, "gconv.SliceStruct failed: %v", err)
		return err
	}
	chatInput.Message = messages
	chatInput.Context = codeContexts
	err = common.LLMRun(ctx, &chatin.ChatLLMRunReq{
		Input:            chatInput,
		Output:           out,
		ConversationType: consts.ConversationExplain.String(),
		Prompt:           prompts,
		PromptVersion:    explain.PromptVersion,
	})
	if err != nil {
		g.Log().Errorf(ctx, "ConversationUUID: %s LLMRun failed: %v", in.ConversationUUID, err)
		return err
	}

	return nil
}
