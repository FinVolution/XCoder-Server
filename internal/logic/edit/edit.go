package edit

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"xcoder/core/prompts/edit"
	"xcoder/core/schema"
	"xcoder/internal/consts"
	"xcoder/internal/controller/utils/ccode"
	"xcoder/internal/controller/utils/cerror"
	"xcoder/internal/logic/common"
	"xcoder/internal/model/input/chatin"
	"xcoder/internal/model/input/editin"
	"xcoder/internal/service"
	"xcoder/utility/xutils"
)

type sEdit struct{}

func NewEdit() service.IEdit {
	return &sEdit{}
}

func init() {
	service.RegisterEdit(NewEdit())
}

func (s *sEdit) SseGenerate(ctx context.Context, in *editin.EditSseGenerateRequest, out chan *chatin.ChatResult) error {
	// 将 code language 转换为小写，并将 go 转换为 golang
	codeLanguage := xutils.CodeLanguageToLower(in.CodeLanguage)

	// 组建 prompts
	prompts, err := edit.GenerateEditPrompt(ctx,
		&editin.GenerateEditPromptRequest{
			RepoName:       in.GitRepo,
			CodePath:       in.CodePath,
			CodeLanguage:   codeLanguage,
			SelectedCode:   in.UserCode,
			UserContent:    in.UserText,
			SharedContexts: in.UserContext,
		})
	if err != nil {
		g.Log().Errorf(ctx, "ConversationUUID: %s, explain.GenerateExplainPrompt failed: %v", in.ConversationUUID, err)
		return cerror.NewUserError(err, ccode.CodeGeneratePromptError, ccode.CodeGeneratePromptErrorMessage)
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
		ConversationType: consts.ConversationEdit.String(),
		Prompt:           prompts,
		PromptVersion:    edit.PromptVersion,
	})
	if err != nil {
		g.Log().Errorf(ctx, "ConversationUUID: %s LLMRun failed: %v", in.ConversationUUID, err)
		return err
	}

	return nil
}
