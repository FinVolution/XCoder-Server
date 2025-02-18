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
	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkoukk/tiktoken-go"
)

const (
	_tokenApproximation = 4
)

const (
	_gpt35TurboContextSize   = 4096
	_gpt432KContextSize      = 32768
	_gpt4ContextSize         = 8192
	_textDavinci3ContextSize = 4097
	_textBabbage1ContextSize = 2048
	_textAda1ContextSize     = 2048
	_textCurie1ContextSize   = 2048
	_codeDavinci2ContextSize = 8000
	_codeCushman1ContextSize = 2048
	_textBisonContextSize    = 2048
	_chatBisonContextSize    = 2048
	_defaultContextSize      = 2048
)

// nolint:gochecknoglobals
var modelToContextSize = map[string]int{
	"gpt-3.5-turbo":    _gpt35TurboContextSize,
	"gpt-4-32k":        _gpt432KContextSize,
	"gpt-4":            _gpt4ContextSize,
	"text-davinci-003": _textDavinci3ContextSize,
	"text-curie-001":   _textCurie1ContextSize,
	"text-babbage-001": _textBabbage1ContextSize,
	"text-ada-001":     _textBabbage1ContextSize,
	"code-davinci-002": _codeDavinci2ContextSize,
	"code-cushman-001": _codeCushman1ContextSize,
}

// GetModelContextSize gets the max number of tokens for a language model. If the model
// name isn't recognized the default value 2048 is returned.
func GetModelContextSize(model string) int {
	contextSize, ok := modelToContextSize[model]
	if !ok {
		return _defaultContextSize
	}
	return contextSize
}

// CountTokens gets the number of tokens the text contains.
func CountTokens(model, text string) int {
	e, err := tiktoken.EncodingForModel(model)
	if err != nil {
		e, err = tiktoken.GetEncoding("gpt2")
		if err != nil {
			g.Log().Infof(context.Background(), "[INFO] Can not use model to calculate number of tokens, "+
				"falling back to approximate count")
			return len([]rune(text)) / _tokenApproximation
		}
	}
	return len(e.Encode(text, nil, nil))
}

// CalculateMaxTokens calculates the max number of tokens that could be added to a text.
func CalculateMaxTokens(model, text string) int {
	return GetModelContextSize(model) - CountTokens(model, text)
}
