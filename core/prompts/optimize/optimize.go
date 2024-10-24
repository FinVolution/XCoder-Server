package optimize

import (
	"context"
	"fmt"
	"xcoder/internal/model/input/optimizein"
)

const SystemMessage = `你是XCoder，一个由XinYe公司开发的AI编程助手。当被问及你的名字时，你需要回答“XCoder”。
你需要严格按照用户的要求行事。以下是你的工作方式：
1. 每个代码块都以%[1]s开始。
2. 你总是用%[2]s语言回答。
3. 当用户要求你提供代码时，你需要以%[2]s代码块的形式回答。
4. 在答案中使用 Markdown 格式。
5. 你的答案尽量简短且仅限于技术领域。
请注意，你是一个AI编程助手，只对软件开发相关的问题进行回答。`

const userInstruction = `优化用户选中代码。你可以按照以下步骤进行代码优化：
1. 理解用户代码：仔细阅读用户提供的代码。
2. 问题检测：分析代码逻辑、性能、安全性等方面，寻找潜在的缺陷或改进点。
3. 提供解决方案：如果发现问题，提出优化建议并提供相应的代码示例；如果未检测到问题，提醒用户准确描述需要修复的内容。
4. 逐步解释：在给出代码之前，先阐述编写计划，以便用户了解优化的思路和步骤。
5. 代码输出：将优化后的代码放入代码块中，确保清晰易读。
注意：除了代码部分，其他总结性语句必须用中文返回。`

func GenerateOptimizePrompt(ctx context.Context, in *optimizein.GenerateCodeOptimizePromptRequest) ([]map[string]string, error) {
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

	userInstructionTmpl := []map[string]string{
		{
			"role":    "user",
			"content": userInstruction,
		},
	}

	tmpls = append(tmpls, selectedTmpl...)
	tmpls = append(tmpls, userInstructionTmpl...)

	return tmpls, nil
}
