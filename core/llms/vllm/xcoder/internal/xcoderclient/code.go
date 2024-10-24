package xcoderclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	"net/http"
)

type CodeGenRequest struct {
	Model            string   `json:"model"`
	Prompt           string   `json:"prompt"`
	Temperature      float64  `json:"temperature"`
	MaxTokens        int      `json:"max_tokens"`
	FrequencyPenalty float64  `json:"frequency_penalty,omitempty"`
	PresencePenalty  float64  `json:"presence_penalty,omitempty"`
	TopP             float64  `json:"top_p,omitempty"`
	StopWords        []string `json:"stop,omitempty"`
	Stream           bool     `json:"stream,omitempty"`

	// StreamingFunc is a function to be called for each chunk of a streaming response.
	// Return an error to stop streaming early.
	StreamingFunc func(ctx context.Context, chunk []byte) error `json:"-"`
}

type CodeGenModelResponse struct {
	ID      string  `json:"id,omitempty"`
	Created float64 `json:"created,omitempty"`
	Choices []struct {
		FinishReason string      `json:"finish_reason,omitempty"`
		Index        float64     `json:"index,omitempty"`
		Logprobs     interface{} `json:"logprobs,omitempty"`
		Text         string      `json:"text,omitempty"`
	} `json:"choices,omitempty"`
	Model  string `json:"model,omitempty"`
	Object string `json:"object,omitempty"`
	Usage  struct {
		CompletionTokens int `json:"completion_tokens,omitempty"`
		PromptTokens     int `json:"prompt_tokens,omitempty"`
		TotalTokens      int `json:"total_tokens,omitempty"`
	} `json:"usage,omitempty"`
}

func (c *Client) setCodeGenDefaults(ctx context.Context, payload *CodeGenRequest) {
	if payload.MaxTokens == 0 {
		payload.MaxTokens = 160
	}

	if len(payload.StopWords) == 0 {
		payload.StopWords = nil
	}
}

func (c *Client) createCodeGen(ctx context.Context, payload *CodeGenRequest) (*CodeGenModelResponse, error) {
	g.Log().Infof(ctx, "=== Model: %s, baseUrl: %s ===", c.Model, c.baseURL)
	c.setCodeGenDefaults(ctx, payload)
	return c.createCodeGenDo(ctx, &CodeGenRequest{
		Model:            payload.Model,
		Prompt:           payload.Prompt,
		Temperature:      payload.Temperature,
		TopP:             payload.TopP,
		MaxTokens:        payload.MaxTokens,
		StopWords:        payload.StopWords,
		FrequencyPenalty: payload.FrequencyPenalty,
		PresencePenalty:  payload.PresencePenalty,
		StreamingFunc:    payload.StreamingFunc,
	})
}

func (c *Client) createCodeGenDo(ctx context.Context, payload *CodeGenRequest) (*CodeGenModelResponse, error) {
	if payload.StreamingFunc != nil {
		payload.Stream = true
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	g.Log().Infof(ctx, "=== payloadBytes: %s === ", string(payloadBytes))

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/v1/completions", body)
	if err != nil {
		return nil, err
	}

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("API returned unexpected status code: %d", r.StatusCode)

		var errResp errorMessage
		if err := json.NewDecoder(r.Body).Decode(&errResp); err != nil {
			return nil, errors.New(msg)
		}

		return nil, fmt.Errorf("%s: %s", msg, errResp.Error.Message)
	}

	content, err := io.ReadAll(r.Body)

	if payload.StreamingFunc != nil {
		parseStreamingCodeGenResponse(ctx, r, payload)
	}

	var response CodeGenModelResponse
	err = gconv.Struct(content, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func parseStreamingCodeGenResponse(ctx context.Context, r *http.Response, payload *CodeGenRequest) {
	fmt.Println("Implement me")
}
