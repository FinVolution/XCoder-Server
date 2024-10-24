package callbacks

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
	"time"
	"xcoder/core/llms"
	"xcoder/internal/dao/mongodb"
	"xcoder/internal/model/input/chatin"
	"xcoder/internal/model/input/codein"
	"xcoder/internal/service"
	"xcoder/utility/xconcurrent"
)

type HandlerImpl struct{}

func (h HandlerImpl) HandleLLMStart(ctx context.Context, prompts []string) {
	fmt.Println("Entering LLM with prompts:", prompts)
}

func (h HandlerImpl) HandleLLMEnd(ctx context.Context, output llms.LLMResult) {
	fmt.Println("Implement me")
}

func (h HandlerImpl) HandleLLMPredictEnd(ctx context.Context, output llms.LLMResult) error {
	completionCode := output.Generations[0].Text
	completionCodeLines := len(strings.SplitAfter(completionCode, "\n"))

	dataToMysql := &codein.CodeUpdateGenerationInfoReq{
		GenerateUUID:         output.Generations[0].GenerationInfo["generate_uuid"].(string),
		PromptTokens:         output.Generations[0].GenerationInfo["prompt_tokens"].(int),
		CompletionCode:       completionCode,
		CompletionCodeTokens: output.Generations[0].GenerationInfo["completion_code_tokens"].(int),
		CompletionCodeLines:  completionCodeLines,
		CompletionDuration:   output.Generations[0].GenerationInfo["completion_duration"].(int),
		FinishReason:         output.Generations[0].GenerationInfo["finish_reason"].(string),
		FailureReason:        output.Generations[0].GenerationInfo["failure_reason"].(string),
	}

	dataToMongoDB := &codein.CodeGenerateUpdateMongoRequest{
		GenerateUUID:      output.Generations[0].GenerationInfo["generate_uuid"].(string),
		RawCompletionCode: output.Generations[0].GenerationInfo["completion_code"].(string),
		CompletionCode:    completionCode,
	}

	p := xconcurrent.NewBase("codeGenerateRecordUpdate")
	p.Compute(ctx, func(ctx context.Context) error {
		g.Log().Infof(ctx, "AsyncCodeGenerateUpdateDB start to update mysql")
		// 某些情况，insert 速度慢，等待1s后再次更新
		time.Sleep(1 * time.Second)

		// 请求字段保存到 mysql
		_, err := service.Code().UpdateGenerationInfo(ctx, dataToMysql)
		if err != nil {
			g.Log().Errorf(ctx, "Mysql CodeGenerateUpdate failed: %v", err)
			return err
		}

		g.Log().Infof(ctx, "AsyncCodeGenerateUpdateDB start to update mongodb")
		// 保存代码到 mongodb
		_, err = mongodb.MDao.CodeGenerateUpdate(ctx, dataToMongoDB)
		if err != nil {
			g.Log().Errorf(ctx, "Mongodb CodeGenerateUpdate failed: %v", err)
			return err
		}

		g.Log().Infof(ctx, "AsyncCodeGenerateUpdateDB finished")

		return nil
	})

	return nil
}

func getGenerateCodeFromResponse(ctx context.Context, response string) string {
	// 获取代码
	completionCode := strings.Split(response, "```")
	if len(completionCode) < 2 {
		g.Log().Infof(ctx, "getGenerateCodeFromResponse no code in completion result: %v", response)
		return ""
	}

	var result string
	lines := strings.Split(completionCode[1], "\n")
	if len(lines) > 1 {
		result = strings.Join(lines[1:], "\n")
	}

	return result
}

func (h HandlerImpl) HandleLLMChatEnd(ctx context.Context, output llms.LLMResult) error {
	dataToMysql := &chatin.ChatMessageUpdateMysqlReq{
		ConversationUUID:   output.Generations[0].GenerationInfo["conversation_uuid"].(string),
		PromptTokens:       output.Generations[0].GenerationInfo["prompt_tokens"].(int),
		CompletionTokens:   output.Generations[0].GenerationInfo["completion_tokens"].(int),
		TotalTokens:        output.Generations[0].GenerationInfo["total_tokens"].(int),
		CompletionDuration: gconv.Int(output.Generations[0].GenerationInfo["completion_duration"].(int64)),
	}

	completionCode := getGenerateCodeFromResponse(ctx, output.Generations[0].Text)

	dataToMongoDB := &chatin.ChatMessageUpdateMongoRequest{
		ConversationUUID: output.Generations[0].GenerationInfo["conversation_uuid"].(string),
		Response:         output.Generations[0].Text,
		CompletionCode:   completionCode,
	}

	// 异步更新数据到 mysql 和 mongodb
	p := xconcurrent.NewBase("chatMessageUpdate")
	p.Compute(ctx, func(ctx context.Context) error {
		g.Log().Infof(ctx, "AsyncChatMessageUpdateDB start to update mysql")
		// 请求字段保存到 mysql
		_, err := service.Chat().UpdateChatRecordsFields(ctx, dataToMysql)
		if err != nil {
			g.Log().Errorf(ctx, "Mysql ChatMessageUpdate failed: %v", err)
			return err
		}

		g.Log().Infof(ctx, "AsyncChatMessageUpdateDB start to update mongodb")
		// 保存代码到 mongodb
		_, err = mongodb.MDao.ChatMessageUpdate(ctx, dataToMongoDB)
		if err != nil {
			g.Log().Errorf(ctx, "Mongodb ChatMessageUpdate failed: %v", err)
			return err
		}

		g.Log().Infof(ctx, "AsyncChatMessageUpdateDB finished")

		return nil
	})

	return nil
}
