package xcoder

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"os"
	"time"
	"xcoder/core/callbacks"
	"xcoder/core/llms"
	"xcoder/core/llms/vllm/xcoder/internal/xcoderclient"
	"xcoder/utility/xcontext"
	"xcoder/utility/xutils"
)

type LLM struct {
	CallBacksHandler callbacks.Handler
	client           *xcoderclient.Client
}

func New(opts ...Option) (*LLM, error) {
	c, err := newClient(opts...)
	return &LLM{
		CallBacksHandler: callbacks.HandlerImpl{},
		client:           c,
	}, err
}

var (
	ErrEmptyResponse = errors.New("empty response")
)

func newClient(opts ...Option) (*xcoderclient.Client, error) {
	options := &options{
		model:   os.Getenv(modelEnvVarName),
		baseURL: os.Getenv(baseURLEnvVarName),
	}
	for _, opt := range opts {
		opt(options)
	}

	return xcoderclient.New(options.model, options.baseURL), nil
}

func (c *LLM) Predict(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	r, err := c.Generate(ctx, []string{prompt}, options...)
	if err != nil {
		return "", err
	}
	return r[0].Text, nil
}

func (c *LLM) Generate(ctx context.Context, prompts []string, options ...llms.CallOption) ([]*llms.Generation, error) {
	opts := llms.CallOptions{}
	for _, opt := range options {
		opt(&opts)
	}

	generations := make([]*llms.Generation, 0, len(prompts))
	for _, p := range prompts {
		startTime := time.Now()
		result, err := c.client.CreateCodeGenWithDetail(ctx, &xcoderclient.CodeGenRequest{
			Model:            opts.Model,
			Prompt:           p,
			MaxTokens:        opts.MaxTokens,
			StopWords:        opts.StopWords,
			Temperature:      opts.Temperature,
			FrequencyPenalty: opts.FrequencyPenalty,
			PresencePenalty:  opts.PresencePenalty,
			TopP:             opts.TopP,
			StreamingFunc:    opts.StreamingFunc,
		})
		if err != nil {
			return nil, err
		}
		if len(result.Choices) == 0 {
			return nil, ErrEmptyResponse
		}

		elapsedTime := time.Since(startTime)
		generationInfo := make(map[string]any)
		generationInfo["completion_code_tokens"] = result.Usage.CompletionTokens
		generationInfo["prompt_tokens"] = result.Usage.PromptTokens
		generationInfo["total_tokens"] = result.Usage.TotalTokens
		generationInfo["finish_reason"] = result.Choices[0].FinishReason
		generationInfo["failure_reason"] = ""
		generationInfo["completion_duration"] = int(elapsedTime.Milliseconds())
		generationInfo["generate_uuid"] = opts.UUID
		generationInfo["completion_code"] = result.Choices[0].Text

		generateText := result.Choices[0].Text
		// 如果停止原因是 length，则将最后一行的代码去除
		if generationInfo["finish_reason"] == "length" {
			generateText = xutils.HandleGenerateTextWithLengthStop(ctx, generateText)
		}
		// 如果是多行生成，去除重复的代码
		if opts.IsMultiLine {
			generateText = xutils.HandleRemoveDuplicateCodeForMultiLine(ctx, generateText, opts.SuffixCode)
		}
		// 如果是单行生成，去除尾部的空格
		if !opts.IsMultiLine {
			generateText = xutils.HandleRemoveTailSpaceIfExistForSingleLine(ctx, generateText)
		}

		generations = append(generations, &llms.Generation{
			Text:           generateText,
			GenerationInfo: generationInfo,
		})
	}

	g.Log().Infof(ctx, "HandleLLMPredictEnd start")
	err := c.CallBacksHandler.HandleLLMPredictEnd(xcontext.WithProtect(ctx), llms.LLMResult{
		Generations: generations,
	})
	if err != nil {
		g.Log().Errorf(ctx, "HandleLLMPredictEnd error: %v", err)
		return nil, err
	}

	return generations, nil
}
