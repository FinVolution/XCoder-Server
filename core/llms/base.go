// Copyright (c) Travis Cline <travis.cline@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package llms

import (
	"context"
	"xcoder/core/schema"
)

type LanguageModel interface {
	GeneratePrompt(ctx context.Context, prompts []schema.PromptValue, options ...CallOption) (LLMResult, error)
	GetNumTokens(text string) int
}

type LLM interface {
	Call(ctx context.Context, prompt string, options ...CallOption) (string, error)
	Generate(ctx context.Context, prompts []string, options ...CallOption) ([]*Generation, error)
}

// ChatLLM is a langchaingo LLM that can be used for chatting.
type ChatLLM interface {
	Call(ctx context.Context, messages []schema.ChatMessage, options ...CallOption) (*schema.AIChatMessage, error)
	Generate(ctx context.Context, messages [][]schema.ChatMessage, options ...CallOption) ([]*Generation, error)
}

type CodeLLM interface {
	Predict(ctx context.Context, prompt string, option ...CallOption) (string, error)
	PredictDo(ctx context.Context, prompts []string, options ...CallOption) ([]*Generation, error)
}

type Generation struct {
	Text string `json:"text"`
	// Message stores the potentially generated message.
	Message        *schema.AIChatMessage `json:"message"`
	GenerationInfo map[string]any        `json:"generation_info"`
}

type LLMResult struct {
	Generations []*Generation
	LLMOutput   map[string]any
}
