package chat

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"xcoder/internal/model/input/chatin"
)

const SystemMessage = `你是XCoder，一个由XinYe公司开发的AI编程助手。当被问及你的名字时，你需要回答“XCoder”。
你需要严格按照用户的要求行事。以下是你的工作方式：
1. 每个代码块都以 %[1]s 开始。
2. 你的回答使用 Markdown 格式。
3. 你的回答尽量简短且仅限于技术领域。
请注意，你是一个AI编程助手，只对软件开发相关的问题进行回答。`

const userInstruction = `在回答用户问题时，请遵循以下格式：
在输出代码之前，首先对代码的目的、逻辑和实现步骤进行清晰的描述。这将帮助用户理解代码的背景和用途。
使用适当的代码块语法，并在代码块的开头标明使用的编程语言。确保文件名和路径用单引号括起，便于识别。
在代码部分结束后，使用中文总结性语句，简要概括代码的功能或注意事项。
注意：如果用户问题中指定了编程语言，则按照用户指定的回答；否则，可以使用 %s 代码语言回答用户问题`

func GenerateChatPrompt(ctx context.Context, in *chatin.GenerateChatPromptReq) ([]map[string]string, error) {
	// 系统消息
	tmpls := []map[string]string{
		{
			"role":    "system",
			"content": fmt.Sprintf(SystemMessage, "```", in.CodeLanguage),
		},
	}

	// 加入历史聊天记录
	idx := len(in.Messages) - 1
	lastMsg := in.Messages[idx]
	beforeMsgs := in.Messages[:idx]
	if len(beforeMsgs) > 0 {
		historyStr := gconv.String(beforeMsgs)
		historyMsg := map[string]string{
			"role":    "user",
			"content": fmt.Sprintf("这些是用户聊天记录的会话历史:\n%s\n", historyStr),
		}
		tmpls = append(tmpls, historyMsg)
	}

	// 加入用户选择的代码
	if len(in.SelectedCode) > 0 {
		selectedMsg := map[string]string{
			"role": "user",
			"content": fmt.Sprintf("这个代码片段是用户挑选的代码，希望你针对这段代码进行回答:\n```%s\n%s\n```",
				in.CodeLanguage, in.SelectedCode),
		}
		tmpls = append(tmpls, selectedMsg)
	}

	// 加入用户指令
	userInstructionMsg := map[string]string{
		"role":    "user",
		"content": fmt.Sprintf(userInstruction, in.CodeLanguage),
	}
	tmpls = append(tmpls, userInstructionMsg)

	// 加入用户最新的问题
	lastMsgMsg := map[string]string{
		"role":    "user",
		"content": lastMsg.Content,
	}
	tmpls = append(tmpls, lastMsgMsg)

	return tmpls, nil
}
