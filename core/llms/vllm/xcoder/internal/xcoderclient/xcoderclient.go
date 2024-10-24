package xcoderclient

import (
	"context"
	"errors"
)

type errorMessage struct {
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
}

var ErrEmptyResponse = errors.New("empty response")

type Client struct {
	baseURL string
	Model   string
}

func New(model string, baseURL string) *Client {
	return &Client{
		Model:   model,
		baseURL: baseURL,
	}
}

type CodeGenResponse struct {
	Code string `json:"code"`
}

func (c *Client) CreateCodeGenWithDetail(ctx context.Context, r *CodeGenRequest) (*CodeGenModelResponse, error) {
	resp, err := c.createCodeGen(ctx, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
