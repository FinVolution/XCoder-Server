package edit

import (
	"context"
	"fmt"
	"xcoder/internal/model/input/editin"
)

const SystemMessage = `你是XCoder，一个由XinYe公司开发的AI编程助手。当被问及你的名字时，你需要回答“XCoder”。
你需要严格按照用户的要求行事。以下是你的工作方式：
1. 每个代码块都以 %[1]s 开始。
2. 你的回答使用 Markdown 格式。
3. 你的回答尽量简短且仅限于技术领域。
请注意，你是一个AI编程助手，只对软件开发相关的问题进行回答。`

const userInstruction = `你是一个代码编辑器，你的任务是按照用户的要求，参考用户给定的上下文代码片段（如果用户提供），实现编辑用户挑选的代码。你需要按照以下要求进行代码编辑：
- 逐步思考与计划：在开始编辑代码之前，仔细分析用户提供的代码，明确代码的功能和结构。
确定需要进行的更新或修改，包括哪些功能需要添加、哪些逻辑需要优化。
- 保持代码风格一致：更新后的代码必须与用户提供的代码在缩进、空白和风格上保持一致。`

func GenerateEditPrompt(ctx context.Context, in *editin.GenerateEditPromptRequest) ([]map[string]string, error) {
	tmpls := []map[string]string{
		{
			"role":    "system",
			"content": fmt.Sprintf(SystemMessage, "```", in.CodeLanguage),
		},
	}

	for _, c := range in.SharedContexts {
		tmpls = append(tmpls, map[string]string{
			"role": "user",
			"content": fmt.Sprintf(
				"这是用户的一些上下文代码片段，来自文件：`%s`:\n```%s\n%s\n```",
				c.Path, in.CodeLanguage, c.Content,
			),
		})
	}

	selectedTmpl := []map[string]string{
		{
			"role": "user",
			"content": fmt.Sprintf("这是用户从文件: `%[1]s`，挑选的 %[2]s 代码，代码内容为:\n```%[2]s\n%[3]s\n```",
				in.CodePath, in.CodeLanguage, in.SelectedCode),
		},
	}

	endTmpl := []map[string]string{
		{
			"role":    "user",
			"content": userInstruction,
		},
	}

	userTmpl := []map[string]string{
		{
			"role":    "user",
			"content": in.UserContent,
		},
	}

	tmpls = append(tmpls, selectedTmpl...)
	tmpls = append(tmpls, endTmpl...)
	tmpls = append(tmpls, userTmpl...)

	return tmpls, nil
}
