package unit_test

import (
	"context"
	"fmt"
	"xcoder/internal/consts"
	"xcoder/internal/model/input/unit_testin"
	"xcoder/utility/xutils"
)

const SystemMessage = `你是XCoder，一个由XinYe公司开发的AI编程助手。当被问及你的名字时，你需要回答“XCoder”。
你需要严格按照用户的要求行事。以下是你的工作方式：
1. 每个代码块都以%[1]s开始。
2. 你总是用%[2]s语言回答。
3. 当用户要求你提供代码时，你需要以%[2]s代码块的形式回答。
4. 在答案中使用 Markdown 格式。
5. 你的答案尽量简短且仅限于技术领域。
请注意，你是一个AI编程助手，只对软件开发相关的问题进行回答。`

const UserInstruction = `请为用户挑选的代码生成单元测试，生成过程严格遵循下面的每一个要求：
- 以上有一些用户给定参考的代码片段。包括但不限于有以下几类：
  1. 用户选择的代码所在的上下文文件； 2. 用户选择的代码对应的现有单元测试文件；3. 其他示例的单元测试文件； 
- 每个代码片段都是以 %[1]s 以及 %[2]s FILEPATH 开始的。
- 需要符合一下几点单元测试的编写规范：
  1. 使用简单而完整的断言来验证关键功能；2. 无论何时何地运行测试，结果都应该是一致的；3. 测试应该是独立的，不应该依赖于任何外部状态。
- 如果用户提供的参考单元测试文件中，存在对用户挑选的代码的测试，则需要在测试中添加其他新的测试，而不是修改现有的测试。
- 新的测试应该验证预期的功能，并涵盖用户代码的所有必需导入的边缘情况，包括导入被测试的函数。不要重复现有的测试。
- 新的测试应该放置在一个单独的 markdown %[1]s 代码块中。
- 如果 %[3]s 不为空，则使用 %[3]s 框架的单元测试规范。
- 如果 %[3]s 为空，则使用符合 %[1]s 语言框架的单元测试规范。
- 使用中文进行总结测试的覆盖范围以及局限性。%[4]s`

func GenerateUnitTestPrompt(ctx context.Context, in *unit_testin.GenerateUTPromptReq) ([]map[string]string, error) {
	codeLanguageAnnotationFlag := xutils.CodeLanguageAnnotationFlagGet(in.CodeLanguage)

	tmpls := []map[string]string{
		{
			"role":    "system",
			"content": fmt.Sprintf(SystemMessage, "```", in.CodeLanguage),
		},
	}

	// 加入用户上下文信息
	for _, cContext := range in.SharedContexts {
		ctxContext := cContext.Content
		var content string
		switch cContext.Type {
		case consts.ContextFileLocal.String():
			content = fmt.Sprintf("\n\n这是测试代码所在的本地上下文文件：\n\n\n```%s\n%s FILEPATH: %s\n\n%s\n```\n\n因为这是一个本地上下文文件，通常情况下，除非是为了验证预期功能或覆盖边缘情况，否则没有必要引用这个文件。",
				in.CodeLanguage, codeLanguageAnnotationFlag, cContext.Path, ctxContext)
		case consts.ContextFileLocalTest.String():
			content = fmt.Sprintf("\n\n以下是现有测试文件的摘录：\n\n\n```%s\n%s FILEPATH: %s\n\n%s\n```\n\n因为存在一个现有测试文件：\n- 不要生成前言部分，比如 import、copyright 等。\n- 要生成可以追加到现有测试文件中的代码。",
				in.CodeLanguage, codeLanguageAnnotationFlag, cContext.Path, ctxContext)
		case consts.ContextFileOtherTest.String():
			content = fmt.Sprintf("这是一个示例测试文件：\n\n\n```%s\n%s FILEPATH: %s\n\n%s\n```",
				in.CodeLanguage, codeLanguageAnnotationFlag, cContext.Path, ctxContext)
		default:
			content = fmt.Sprintf("这是一个示例测试文件：\n\n\n```%s\n%s FILEPATH: %s\n\n%s\n```",
				in.CodeLanguage, codeLanguageAnnotationFlag, cContext.Path, ctxContext)
		}

		tmpls = append(tmpls, map[string]string{
			"role":    "user",
			"content": content,
		})
	}

	// 加入用户选择的代码
	selectedTmpl := []map[string]string{
		{
			"role": "user",
			"content": fmt.Sprintf("这是用户从文件: `%[1]s`，挑选的 %[2]s 代码，代码内容为:\n```%[2]s\n%[3]s\n```",
				in.CodePath, in.CodeLanguage, in.SelectedCode),
		},
	}

	// 加入用户的指令
	UserUTInstruction := fmt.Sprintf(UserInstruction, in.CodeLanguage, codeLanguageAnnotationFlag, in.Framework, in.UserInstruction)
	endTmpl := []map[string]string{
		{
			"role":    "user",
			"content": UserUTInstruction,
		},
	}

	tmpls = append(tmpls, selectedTmpl...)
	tmpls = append(tmpls, endTmpl...)

	return tmpls, nil
}
