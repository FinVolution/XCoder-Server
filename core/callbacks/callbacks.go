package callbacks

import (
	"context"
	"xcoder/core/llms"
)

type Handler interface {
	HandleLLMStart(ctx context.Context, prompts []string)
	HandleLLMEnd(ctx context.Context, output llms.LLMResult)
	HandleLLMPredictEnd(ctx context.Context, output llms.LLMResult) error
	HandleLLMChatEnd(ctx context.Context, output llms.LLMResult) error
}
