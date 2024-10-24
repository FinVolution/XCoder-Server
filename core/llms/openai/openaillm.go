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

package openai

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"time"
	"xcoder/core/callbacks"
	"xcoder/core/llms"
	"xcoder/core/llms/openai/internal/openaiclient"
	"xcoder/core/schema"
	"xcoder/utility/xcontext"
)

type Chat struct {
	CallbacksHandler callbacks.Handler
	client           *openaiclient.Client
}

// NewChat returns a new OpenAI chat LLM.
func NewChat(opts ...Option) (*Chat, error) {
	_, c, err := newClient(opts...)
	if err != nil {
		return nil, err
	}
	return &Chat{
		client:           c,
		CallbacksHandler: callbacks.HandlerImpl{},
	}, err
}

// Call requests a chat response for the given messages.
func (o *Chat) Call(ctx context.Context, messages []schema.ChatMessage, options ...llms.CallOption) (*schema.AIChatMessage, error) { // nolint: lll
	r, err := o.Generate(ctx, [][]schema.ChatMessage{messages}, options...)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, ErrEmptyResponse
	}
	return r[0].Message, nil
}

func (o *Chat) Generate(ctx context.Context, messageSets [][]schema.ChatMessage, options ...llms.CallOption) ([]*llms.Generation, error) { // nolint:lll,cyclop
	//if o.CallbacksHandler != nil {
	//	o.CallbacksHandler.HandleLLMStart(ctx, getPromptsFromMessageSets(messageSets))
	//}

	opts := llms.CallOptions{}
	for _, opt := range options {
		opt(&opts)
	}
	generations := make([]*llms.Generation, 0, len(messageSets))
	for _, messageSet := range messageSets {
		req := &openaiclient.ChatRequest{
			Model:            opts.Model,
			StopWords:        opts.StopWords,
			Messages:         messagesToClientMessages(messageSet),
			StreamingFunc:    opts.StreamingFunc,
			Temperature:      opts.Temperature,
			MaxTokens:        opts.MaxTokens,
			N:                opts.N,
			FrequencyPenalty: opts.FrequencyPenalty,
			PresencePenalty:  opts.PresencePenalty,

			FunctionCallBehavior: openaiclient.FunctionCallBehavior(opts.FunctionCallBehavior),
		}
		for _, fn := range opts.Functions {
			req.Functions = append(req.Functions, openaiclient.FunctionDefinition{
				Name:        fn.Name,
				Description: fn.Description,
				Parameters:  fn.Parameters,
			})
		}
		start := time.Now()
		result, err := o.client.CreateChat(ctx, req)
		if err != nil {
			return nil, err
		}
		if len(result.Choices) == 0 {
			return nil, ErrEmptyResponse
		}
		generationInfo := make(map[string]any)
		generationInfo["conversation_uuid"] = opts.ConversationID
		generationInfo["completion_tokens"] = result.Usage.CompletionTokens
		generationInfo["prompt_tokens"] = result.Usage.PromptTokens
		generationInfo["total_tokens"] = result.Usage.TotalTokens
		duration := time.Since(start).Milliseconds()
		generationInfo["completion_duration"] = duration
		msg := &schema.AIChatMessage{
			Content: result.Choices[0].Message.Content,
		}
		if result.Choices[0].FinishReason == "function_call" {
			msg.FunctionCall = &schema.FunctionCall{
				Name:      result.Choices[0].Message.FunctionCall.Name,
				Arguments: result.Choices[0].Message.FunctionCall.Arguments,
			}
		}
		generations = append(generations, &llms.Generation{
			Message:        msg,
			Text:           msg.Content,
			GenerationInfo: generationInfo,
		})
	}

	err := o.CallbacksHandler.HandleLLMChatEnd(xcontext.WithProtect(ctx), llms.LLMResult{Generations: generations})
	if err != nil {
		g.Log().Errorf(ctx, "HandleLLMChatEnd error: %v", err)
		return nil, err
	}

	return generations, nil
}

func messagesToClientMessages(messages []schema.ChatMessage) []*openaiclient.ChatMessage {
	msgs := make([]*openaiclient.ChatMessage, len(messages))
	for i, m := range messages {
		msg := &openaiclient.ChatMessage{
			Content: m.GetContent(),
		}
		typ := m.GetType()
		switch typ {
		case schema.ChatMessageTypeSystem:
			msg.Role = "system"
		case schema.ChatMessageTypeAI:
			msg.Role = "assistant"
		case schema.ChatMessageTypeHuman:
			msg.Role = "user"
		case schema.ChatMessageTypeGeneric:
			msg.Role = "user"
		case schema.ChatMessageTypeFunction:
			msg.Role = "function"
		}
		if n, ok := m.(schema.Named); ok {
			msg.Name = n.GetName()
		}
		msgs[i] = msg
	}

	return msgs
}
