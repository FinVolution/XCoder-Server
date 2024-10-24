package common

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"xcoder/core/llms"
	"xcoder/core/llms/openai"
	"xcoder/core/schema"
	"xcoder/internal/consts"
	"xcoder/internal/dao"
	"xcoder/internal/dao/mongodb"
	"xcoder/internal/model/entity"
	"xcoder/internal/model/input/chatin"
	"xcoder/utility/xapollo"
	"xcoder/utility/xconcurrent"
)

func InitLLM(ctx context.Context, llmParams *chatin.ChatLLMParamsConfig) (*openai.Chat, error) {
	apiType := openai.APITypeOpenAI
	if llmParams.Config.LLMType == "azure" {
		apiType = openai.APITypeAzure
	}
	llm, err := openai.NewChat(
		openai.WithAPIType(apiType),
		openai.WithBaseURL(llmParams.Config.LLMParams.APIBase),
		openai.WithToken(llmParams.Config.LLMParams.APIKey),
		openai.WithAPIVersion(llmParams.Config.LLMParams.APIVersion),
		openai.WithModel(llmParams.Config.LLMParams.Model),
	)
	if err != nil {
		g.Log().Errorf(ctx, "openai.NewChat failed: %v", err)
		return nil, err
	}

	return llm, nil
}

func CallLLM(ctx context.Context, llm *openai.Chat, in *chatin.CallLLMReq) (err error) {
	var messages []schema.ChatMessageRequest
	err = gconv.Struct(in.Prompt, &messages)
	if err != nil {
		g.Log().Errorf(ctx, "gconv.Struct failed: %v", err)
		return err
	}

	var chatMsgs []schema.ChatMessage
	for _, msg := range messages {
		switch msg.Role {
		case schema.ChatMessageTypeSystem.String():
			chatMsgs = append(chatMsgs, schema.SystemChatMessage{
				Content: msg.Content,
			})
		case schema.ChatMessageTypeAI.String():
			chatMsgs = append(chatMsgs, schema.AIChatMessage{
				Content: msg.Content,
			})
		case schema.ChatMessageTypeHuman.String():
			chatMsgs = append(chatMsgs, schema.HumanChatMessage{
				Content: msg.Content,
			})
		default:
			return fmt.Errorf("ChatMessageType: %s is not support", msg.Role)
		}
	}

	_, err = llm.Call(ctx, chatMsgs,
		llms.WithUserID(in.Input.CreateUser),
		llms.WithConversationID(in.Input.ConversationUUID),
		llms.WithTemperature(in.LLMConfig.LLMParams.Temperature),
		llms.WithMaxTokens(in.LLMConfig.LLMParams.MaxTokens),
		llms.WithPresencePenalty(in.LLMConfig.LLMParams.PresencePenalty),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			in.Out <- &chatin.ChatResult{Error: err, Data: string(chunk)}
			return nil
		}),
	)
	if err != nil {
		return err
	}

	return nil
}

func ChatMessageInsertDB(ctx context.Context, in *chatin.ChatMessageInsertDBReq) error {
	// 组建 Mysql 存储参数
	chatModel := &entity.PChatRecords{}
	err := gconv.Struct(in.Input, chatModel)
	if err != nil {
		g.Log().Errorf(ctx, "gconv.Struct failed: %v", err)
		return err
	}

	chatModel.Isactive = 1
	chatModel.DeleteToken = "NA"
	chatModel.Inserttime = gtime.Now()
	chatModel.Updatetime = gtime.Now()
	chatModel.ConversationType = in.ConversationType
	chatModel.ProjectName = consts.ProjectName
	chatModel.EngineName = in.LLMConfig.LLMParams.Model
	chatModel.ModelName = in.LLMConfig.LLMParams.Model
	chatModel.ModelVersion = ""
	chatModel.PromptTmplVersion = in.PromptVersion
	chatModel.PromptTmplContent = in.Prompt

	// 组件 MongoDB 存储参数
	dataToMongoDB := &chatin.ChatMessageInsertMongoRequest{
		ConversationUUID: in.Input.ConversationUUID,
		CreateUser:       in.Input.CreateUser,
		SelectedCode:     in.Input.UserCode,
		Messages:         in.Input.Message,
	}

	// 异步插入数据到 mysql 和 mongodb
	p := xconcurrent.NewBase("chatRecordInsertDB")
	p.Compute(ctx, func(ctx context.Context) error {
		g.Log().Infof(ctx, "AsyncChatRecordInsertDB start to insert to mysql")
		// 请求字段保存到 mysql
		err = dao.PChatRecords.Create(ctx, chatModel)
		if err != nil {
			g.Log().Errorf(ctx, "Mysql AsyncChatRecordInsertDB Insert failed: %v", err)
			return err
		}

		g.Log().Infof(ctx, "AsyncChatMessageInsertDB start to insert to mongodb")
		// 保存代码到 mongodb
		_, err = mongodb.MDao.ChatMessageInsert(ctx, dataToMongoDB)
		if err != nil {
			g.Log().Errorf(ctx, "Mongodb AsyncChatMessageInsertDB Insert failed: %v", err)
			return err
		}

		g.Log().Infof(ctx, "AsyncChatMessageInsertDB finished")
		return nil
	})

	return nil
}

func LLMRun(ctx context.Context, in *chatin.ChatLLMRunReq) (err error) {
	// 从阿波罗获取配置信息
	chatLLMParams, err := xapollo.GetChatLLMParams()
	if err != nil {
		g.Log().Errorf(ctx, "xapollo.GetChatLLMParams failed: %v", err)
		return err
	}

	// 初始化 llm
	llm, err := InitLLM(ctx, chatLLMParams)
	if err != nil {
		g.Log().Errorf(ctx, "common.InitLLM failed: %v", err)
		return err
	}

	// 插入数据库
	if err = ChatMessageInsertDB(ctx, &chatin.ChatMessageInsertDBReq{
		Input:            in.Input,
		LLMConfig:        chatLLMParams.Config,
		ConversationType: in.ConversationType,
		Prompt:           gconv.String(in.Prompt),
		PromptVersion:    in.PromptVersion,
	}); err != nil {
		g.Log().Errorf(ctx, "insertDB failed: %v", err)
		return err
	}

	// 调用 llm
	err = CallLLM(ctx, llm, &chatin.CallLLMReq{
		Input:     in.Input,
		LLMConfig: chatLLMParams.Config,
		Prompt:    in.Prompt,
		Out:       in.Output,
	})
	if err != nil {
		g.Log().Errorf(ctx, "CallLLM failed: %v", err)
		return err
	}

	return nil
}
