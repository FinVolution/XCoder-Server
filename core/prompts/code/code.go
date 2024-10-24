package code

import (
	"github.com/tmc/langchaingo/prompts"
)

func GenerateCodeLLamaPrompt(prefix string, suffix string) (string, error) {
	template := `{{.before_code}}<FILL_ME>{{.after_code}}`

	promptTemplate := prompts.NewPromptTemplate(template, []string{"before_code", "after_code"})
	prompt, err := promptTemplate.FormatPrompt(map[string]any{
		"before_code": prefix,
		"after_code":  suffix,
	})
	if err != nil {
		return "", err
	}

	return prompt.String(), nil
}

func GenerateDeepSeekerPrompt(prefix string, suffix string) (string, error) {
	template := `<｜fim▁begin｜>{{.before_code}}<｜fim▁hole｜>{{.after_code}}<｜fim▁end｜>`

	promptTemplate := prompts.NewPromptTemplate(template, []string{"before_code", "after_code"})
	prompt, err := promptTemplate.FormatPrompt(map[string]any{
		"before_code": prefix,
		"after_code":  suffix,
	})
	if err != nil {
		return "", err
	}

	return prompt.String(), nil
}
