package explain

import (
	"context"
	"fmt"
	"xcoder/internal/model/input/explainin"
)

const SystemMessage = `你是XCoder，一个由XinYe公司开发的AI编程助手。当被问及你的名字时，你需要回答“XCoder”。
你需要严格按照用户的要求行事。以下是你的工作方式：
1. 每个代码块都以%[1]s开始。
2. 你总是用%[2]s语言回答。
3. 当用户要求你提供代码时，你需要以%[2]s代码块的形式回答。
4. 在答案中使用 Markdown 格式。
5. 你的答案尽量简短且仅限于技术领域。
请注意，你是一个AI编程助手，只对软件开发相关的问题进行回答。`

const userInstruction = `要求以简单易懂的方式向初级程序员解释一段代码。需要说明代码的目的、输入、输出，以及它是如何通过逻辑和算法实现目的的。
同时，要指出代码中的重要逻辑流程和数据转换，但避免过于技术性的语言。解释应该清晰、连贯、语法正确，无需逐行解释，针对整个代码块进行解释即可。`

func GenerateExplainPrompt(ctx context.Context, in *explainin.CodeExplainPromptRequest) ([]map[string]string, error) {
	tmpls := []map[string]string{
		{
			"role":    "system",
			"content": fmt.Sprintf(SystemMessage, "```", in.CodeLanguage),
		},
	}

	selectedTmpl := []map[string]string{
		{
			"role": "user",
			"content": fmt.Sprintf("我从文件: `%s`，挑选的 %s 代码，代码内容为:\n```%s\n%s\n```",
				in.CodeLanguage, in.CodePath, in.CodeLanguage, in.SelectedCode),
		},
	}

	endTmpl := []map[string]string{
		{
			"role":    "user",
			"content": userInstruction,
		},
	}
	tmpls = append(tmpls, selectedTmpl...)
	tmpls = append(tmpls, endTmpl...)

	return tmpls, nil
}
