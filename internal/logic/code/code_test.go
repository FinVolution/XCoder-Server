package code

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"xcoder/core/llms"
	"xcoder/internal/model/input/codein"
)

// MockLLM is a mock implementation of the xCoderVLLM.LLM interface
type MockLLM struct {
	mock.Mock
}

func (m *MockLLM) Predict(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	args := m.Called(ctx, prompt, options)
	return args.String(0), args.Error(1)
}

func TestGenerateWithVLLM(t *testing.T) {
	ctx := context.Background()
	mockLLM := new(MockLLM)
	//xCoderVLLM.New = func(options ...xCoderVLLM.Option) (xCoderVLLM.LLM, error) {
	//	return mockLLM, nil
	//}

	in := &codein.CodeGenerateWithLLMRequest{
		GenerateUUID:     "test-uuid",
		CodeBeforeCursor: "before cursor",
		CodeAfterCursor:  "after cursor",
		ModelVersion:     "v1",
		MaxTokens:        100,
		Temperature:      0.7,
		TopP:             0.9,
		StopWords:        []string{"stop"},
		IsSingleLine:     true,
		ConnUrls:         []string{"http://localhost"},
	}

	t.Run("successful response", func(t *testing.T) {
		mockLLM.On("Predict", ctx, mock.Anything, mock.Anything).Return("completion", nil)

		resp, err := generateWithVLLM(ctx, in)
		assert.NoError(t, err)
		assert.Equal(t, "completion", resp.Completion)
	})

	t.Run("error in xCoderVLLM.New", func(t *testing.T) {
		//xCoderVLLM.New = func(options ...xCoderVLLM.Option) (xCoderVLLM.LLM, error) {
		//	return nil, errors.New("failed to create LLM")
		//}

		resp, err := generateWithVLLM(ctx, in)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("error in GenerateCodeLLamaPrompt", func(t *testing.T) {
		//code.GenerateCodeLLamaPrompt = func(before, after string) (string, error) {
		//	return "", errors.New("failed to generate prompt")
		//}

		resp, err := generateWithVLLM(ctx, in)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("error in Predict", func(t *testing.T) {
		//code.GenerateCodeLLamaPrompt = func(before, after string) (string, error) {
		//	return "prompt", nil
		//}
		mockLLM.On("Predict", ctx, mock.Anything, mock.Anything).Return("", errors.New("predict failed"))

		resp, err := generateWithVLLM(ctx, in)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
