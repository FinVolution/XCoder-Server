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

import "context"

type CallOption func(*CallOptions)

type CallOptions struct {
	UUID           string                                        `json:"uuid"`
	Model          string                                        `json:"model"`
	MaxTokens      int                                           `json:"max_tokens"`
	Temperature    float64                                       `json:"temperature"`
	StopWords      []string                                      `json:"stop_words"`
	BadWords       []string                                      `json:"bad_words"`
	StreamingFunc  func(ctx context.Context, chunk []byte) error `json:"streaming_func"`
	TopK           int                                           `json:"top_k"`
	TopP           float64                                       `json:"top_p"`
	Seed           int                                           `json:"seed"`
	UserID         string                                        `json:"userid"`
	ConversationID string                                        `json:"conversation_id"`
	IsMultiLine    bool                                          `json:"is_multi_line"`
	SuffixCode     string                                        `json:"suffix_code"`
	ConnUrls       []string                                      `json:"conn_urls"`
	// MinLength is the minimum length of the generated text.
	MinLength int `json:"min_length"`
	// MaxLength is the maximum length of the generated text.
	MaxLength int `json:"max_length"`
	// N is how many chat completion choices to generate for each input message.
	N int `json:"n"`
	// RepetitionPenalty is the repetition penalty for sampling.
	RepetitionPenalty float64 `json:"repetition_penalty"`
	// FrequencyPenalty is the frequency penalty for sampling.
	FrequencyPenalty float64 `json:"frequency_penalty"`
	// PresencePenalty is the presence penalty for sampling.
	PresencePenalty float64 `json:"presence_penalty"`

	// Function defitions to include in the request.
	Functions []FunctionDefinition `json:"functions"`
	// FunctionCallBehavior is the behavior to use when calling functions.
	//
	// If a specific function should be invoked, use the format:
	// `{"name": "my_function"}`
	FunctionCallBehavior FunctionCallBehavior `json:"function_call"`
}

// FunctionDefinition is a definition of a function that can be called by the model.
type FunctionDefinition struct {
	// Name is the name of the function.
	Name string `json:"name"`
	// Description is a description of the function.
	Description string `json:"description"`
	// Parameters is a list of parameters for the function.
	Parameters any `json:"parameters"`
}

// FunctionCallBehavior is the behavior to use when calling functions.
type FunctionCallBehavior string

const (
	// FunctionCallBehaviorNone will not call any functions.
	FunctionCallBehaviorNone FunctionCallBehavior = "none"
	// FunctionCallBehaviorAuto will call functions automatically.
	FunctionCallBehaviorAuto FunctionCallBehavior = "auto"
)

func WithUUID(model string) CallOption {
	return func(o *CallOptions) {
		o.UUID = model
	}
}

func WithUserID(userID string) CallOption {
	return func(o *CallOptions) {
		o.UserID = userID
	}
}

func WithConversationID(conversationID string) CallOption {
	return func(o *CallOptions) {
		o.ConversationID = conversationID
	}
}

func WithModel(model string) CallOption {
	return func(o *CallOptions) {
		o.Model = model
	}
}

func WithMaxTokens(maxTokens int) CallOption {
	return func(o *CallOptions) {
		o.MaxTokens = maxTokens
	}
}

func WithTemperature(temperature float64) CallOption {
	return func(o *CallOptions) {
		o.Temperature = temperature
	}
}

func WithStopWords(stopWords []string) CallOption {
	return func(o *CallOptions) {
		o.StopWords = stopWords
	}
}

func WithBadWords(badWords []string) CallOption {
	return func(o *CallOptions) {
		o.BadWords = badWords
	}
}

func WithIsMultiLine(isMultiLine bool) CallOption {
	return func(o *CallOptions) {
		o.IsMultiLine = isMultiLine
	}
}

func WithSuffixCode(suffixCode string) CallOption {
	return func(o *CallOptions) {
		o.SuffixCode = suffixCode
	}
}

func WithConnUrls(connUrls []string) CallOption {
	return func(o *CallOptions) {
		o.ConnUrls = connUrls
	}
}

func WithPresencePenalty(presencePenalty float64) CallOption {
	return func(o *CallOptions) {
		o.PresencePenalty = presencePenalty
	}
}

func WithFrequencyPenalty(frequencyPenalty float64) CallOption {
	return func(o *CallOptions) {
		o.FrequencyPenalty = frequencyPenalty
	}
}

func WithTopP(topP float64) CallOption {
	return func(o *CallOptions) {
		o.TopP = topP
	}
}

func WithStreamingFunc(streamingFunc func(ctx context.Context, chunk []byte) error) CallOption {
	return func(o *CallOptions) {
		o.StreamingFunc = streamingFunc
	}
}
