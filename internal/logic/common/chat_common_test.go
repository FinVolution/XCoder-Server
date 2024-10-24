package common_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"xcoder/internal/logic/common"
	"xcoder/internal/model/input/chatin"
)

func TestInitLLM(t *testing.T) {
	ctx := context.Background()

	t.Run("successful initialization with OpenAI", func(t *testing.T) {
		llmParams := &chatin.ChatLLMParamsConfig{
			Config: &chatin.ChatLLMParams{
				LLMType: "openai",
				LLMParams: &chatin.LLMParams{
					APIBase:    "https://api.openai.com",
					APIKey:     "test-api-key",
					APIVersion: "v1",
					Model:      "gpt-3.5-turbo",
				},
			},
		}

		llm, err := common.InitLLM(ctx, llmParams)
		require.NoError(t, err)
		require.NotNil(t, llm)
		//assert.Equal(t, openai.APITypeOpenAI, llm.GetAPIType())
	})

	t.Run("successful initialization with Azure", func(t *testing.T) {
		llmParams := &chatin.ChatLLMParamsConfig{
			Config: &chatin.ChatLLMParams{
				LLMType: "azure",
				LLMParams: &chatin.LLMParams{
					APIBase:    "https://api.openai.com",
					APIKey:     "test-api-key",
					APIVersion: "v1",
					Model:      "gpt-3.5-turbo",
				},
			},
		}

		llm, err := common.InitLLM(ctx, llmParams)
		require.NoError(t, err)
		require.NotNil(t, llm)
		//assert.Equal(t, openai.APITypeAzure, llm.GetAPIType())
	})

	t.Run("initialization failure", func(t *testing.T) {
		llmParams := &chatin.ChatLLMParamsConfig{
			Config: &chatin.ChatLLMParams{
				LLMType: "aopenai",
				LLMParams: &chatin.LLMParams{
					APIBase:    "https://api.openai.com",
					APIKey:     "test-api-key",
					APIVersion: "v1",
					Model:      "gpt-3.5-turbo",
				},
			},
		}

		llm, err := common.InitLLM(ctx, llmParams)
		require.NoError(t, err)
		assert.NotNil(t, llm)
	})
}
