package chatin

import (
	"xcoder/core/schema"
)

type Context struct {
	Type    string `json:"type" binding:"omitempty"`
	Path    string `json:"path" binding:"omitempty"`
	Content string `json:"content" binding:"omitempty"`
}

type ChatSseGenerateReq struct {
	CreateUser       string `json:"createUser" binding:"omitempty"`
	ConversationUUID string `json:"conversationUUID" binding:"required"`
	// git 信息
	GitRepo   string `json:"gitRepo" binding:"omitempty"`
	GitBranch string `json:"gitBranch" binding:"omitempty"`
	// code 信息
	CodePath     string `json:"codePath" binding:"omitempty"`
	CodeLanguage string `json:"codeLanguage" binding:"omitempty"`
	// IDE 信息
	IdeInfo        string `json:"ideInfo" binding:"required"`
	ProjectVersion string `json:"projectVersion" binding:"required"`
	// 聊天信息
	Message []schema.ChatMessageRequest `json:"message"`
	// 上下文信息
	UserCode string    `json:"userCode" binding:"omitempty"`
	Context  []Context `json:"context"`
}

type ChatSseGenerateRes struct{}

type ChatResult struct {
	Error error  `json:"error"`
	Data  string `json:"data"`
}

type ChatLLMRunReq struct {
	Input            *ChatSseGenerateReq
	Output           chan *ChatResult
	ConversationType string
	Prompt           []map[string]string
	PromptVersion    string
}

type ChatMessageInsertDBReq struct {
	Input            *ChatSseGenerateReq
	LLMConfig        *ChatLLMParams
	Out              chan *ChatResult
	ConversationType string
	Prompt           string
	PromptVersion    string
}

type CallLLMReq struct {
	Input     *ChatSseGenerateReq
	LLMConfig *ChatLLMParams
	Prompt    []map[string]string
	Out       chan *ChatResult
}

type ChatMessageUpdateMysqlReq struct {
	ConversationUUID   string `json:"conversationUUID" binding:"required"`
	PromptTokens       int    `json:"prompt_tokens" binding:"required"`
	CompletionTokens   int    `json:"completion_tokens" binding:"required"`
	TotalTokens        int    `json:"total_tokens" binding:"required"`
	CompletionDuration int    `json:"completion_duration" binding:"required"`
}

type ChatMessageUpdateMysqlRes struct{}

type ChatLLMParamsConfig struct {
	Config *ChatLLMParams `json:"config"`
}

type ChatLLMParams struct {
	LLMType   string     `json:"llmType"`
	LLMParams *LLMParams `json:"llmParams"`
}

type LLMParams struct {
	APIBase          string  `json:"apiBase"`
	APIKey           string  `json:"apiKey"`
	APIVersion       string  `json:"apiVersion"`
	Model            string  `json:"model"`
	Temperature      float64 `json:"temperature"`
	MaxTokens        int     `json:"maxTokens"`
	TopP             float64 `json:"topP"`
	PresencePenalty  float64 `json:"presencePenalty"`
	FrequencyPenalty float64 `json:"frequencyPenalty"`
}

type GenerateChatPromptReq struct {
	RepoName     string                      `json:"repoName" binding:"required"`
	CodePath     string                      `json:"codePath" binding:"required"`
	CodeLanguage string                      `json:"codeLanguage" binding:"required"`
	Contexts     []Context                   `json:"Contexts" binding:"omitempty"`
	SelectedCode string                      `json:"selectedCode" binding:"required"`
	Messages     []schema.ChatMessageRequest `json:"messages"`
}

type ChatMessageInsertMongoRequest struct {
	ConversationUUID string                      `json:"conversationUUID" binding:"required"`
	CreateUser       string                      `json:"createUser" binding:"required"`
	SelectedCode     string                      `json:"selectedCode" binding:"required"`
	Messages         []schema.ChatMessageRequest `json:"messages"`
}

type ChatMessageUpdateMongoRequest struct {
	ConversationUUID string `json:"conversationUUID" binding:"required"`
	Response         string `json:"response" binding:"required"`
	CompletionCode   string `json:"completion_code" binding:"required"`
}
