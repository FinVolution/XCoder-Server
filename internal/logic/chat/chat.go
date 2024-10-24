package chat

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"xcoder/core/prompts/chat"
	"xcoder/internal/consts"
	"xcoder/internal/dao"
	"xcoder/internal/logic/common"
	"xcoder/internal/model/input/chatin"
	"xcoder/internal/service"
	"xcoder/utility/xutils"
)

type sChat struct{}

func NewChat() service.IChat {
	return &sChat{}
}

func init() {
	service.RegisterChat(NewChat())
}

func (s *sChat) SseGenerate(ctx context.Context, in *chatin.ChatSseGenerateReq, out chan *chatin.ChatResult) error {
	// 将 code language 转换为小写，并将 go 转换为 golang
	codeLanguage := xutils.CodeLanguageToLower(in.CodeLanguage)

	// 组建 prompts
	prompts, err := chat.GenerateChatPrompt(ctx, &chatin.GenerateChatPromptReq{
		RepoName:     in.GitRepo,
		CodePath:     in.CodePath,
		CodeLanguage: codeLanguage,
		Messages:     in.Message,
		Contexts:     in.Context,
		SelectedCode: in.UserCode,
	})

	// 调用 Chat LLM 服务
	err = common.LLMRun(ctx, &chatin.ChatLLMRunReq{
		Input:            in,
		Output:           out,
		ConversationType: consts.ConversationChat.String(),
		Prompt:           prompts,
		PromptVersion:    chat.PromptVersion,
	})
	if err != nil {
		g.Log().Errorf(ctx, "ConversationUUID: %s LLMRun failed: %v", in.ConversationUUID, err)
		return err
	}

	return nil
}

func (s *sChat) UpdateChatRecordsFields(ctx context.Context, in *chatin.ChatMessageUpdateMysqlReq) (
	out *chatin.ChatMessageUpdateMysqlRes, err error) {
	_, err = dao.PChatRecords.GetOneByConversationUuid(ctx, in.ConversationUUID)
	if err != nil {
		return nil, err
	}

	updatedFields := map[string]interface{}{
		"prompt_tokens":       in.PromptTokens,
		"completion_tokens":   in.CompletionTokens,
		"total_tokens":        in.TotalTokens,
		"completion_duration": in.CompletionDuration,
	}

	err = dao.PChatRecords.UpdateByConversationUuid(ctx, in.ConversationUUID, updatedFields)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
