package comment

import (
	"context"
	"fmt"
	"xcoder/internal/model/input/commentin"
)

const SystemMessage = `你是XCoder，一个由XinYe公司开发的AI编程助手。当被问及你的名字时，你需要回答“XCoder”。
你需要严格按照用户的要求行事。以下是你的工作方式：
1. 每个代码块都以%[1]s开始。
2. 你总是用%[2]s语言回答。
3. 当用户要求你提供代码时，你需要以%[2]s代码块的形式回答。
4. 在答案中使用 Markdown 格式。
5. 你的答案尽量简短且仅限于技术领域。
请注意，你是一个AI编程助手，只对软件开发相关的问题进行回答。`

const userInstruction = `1. 请给选中代码增加函数文档以及代码注释，以便让代码更易于理解。
2. 如果选中代码中存在代码注释，则将其作为示例使用，遵循其相同的风格（生成语言，行数等）返回结果，不要修改已有的代码注释。
3. 仅仅是给选中代码添加注释，不要修改代码逻辑。
4. 生成的函数文档应该根据%s语言特性，生成在合适的位置。
5. 生成的代码注释必须都放在代码所在行的上面一行。
6. 生成的函数文档及代码注释使用中文返回。`

func GenerateCommentPrompt(ctx context.Context, in *commentin.GenerateCommentPromptRequest) ([]map[string]string, error) {
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
			"content": fmt.Sprintf(userInstruction, in.CodeLanguage),
		},
	}
	tmpls = append(tmpls, selectedTmpl...)
	tmpls = append(tmpls, endTmpl...)

	return tmpls, nil
}
