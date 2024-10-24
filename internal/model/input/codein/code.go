package codein

import "xcoder/internal/model/entity"

type Context struct {
	Type    string `json:"type" binding:"omitempty"`
	Path    string `json:"path" binding:"omitempty"`
	Content string `json:"content" binding:"omitempty"`
}

type CodeGenerateRequest struct {
	GenerateUUID string `json:"generateUUID" binding:"required"`
	GenerateType string `json:"generateType" binding:"required"`
	CreateUser   string `json:"createUser" binding:"omitempty"`
	// git 信息
	GitRepo   string `json:"gitRepo" binding:"omitempty"`
	GitBranch string `json:"gitBranch" binding:"omitempty"`
	// code 信息
	CodePath         string `json:"codePath" binding:"required"`
	CodeLanguage     string `json:"codeLanguage" binding:"required"`
	CodeBeforeCursor string `json:"codeBeforeCursor" binding:"required"`
	CodeAfterCursor  string `json:"codeAfterCursor" binding:"omitempty"`
	CodeTotalLines   int    `json:"codeTotalLines" binding:"omitempty"`
	// IDE 信息
	IDEInfo        string `json:"ideInfo" binding:"required"`
	StartCursorIdx int    `json:"cursorStartIdx" binding:"required"`
	// 模型超参
	ModelName   string   `json:"modelName" binding:"omitempty"`
	Temperature float64  `json:"temperature" binding:"omitempty"`
	TopP        float64  `json:"topP" binding:"omitempty"`
	StopWords   []string `json:"stopWords" binding:"omitempty"`
	Stream      bool     `json:"stream" binding:"omitempty"`
	MaxTokens   int      `json:"maxTokens" binding:"omitempty"`
	// 其他参数
	NextIndent        int                    `json:"nextIndent" binding:"omitempty"`
	TrimByIndentation bool                   `json:"trimByIndentation" binding:"omitempty"`
	Extra             map[string]interface{} `json:"extra" binding:"omitempty"`
	// 上下文信息
	Contexts []*Context `json:"contexts" binding:"omitempty"`
}

type CodeGenerateResponse struct {
	// CodeGenerateResponse
	Code string `json:"code"`
}

type CodeUpdateGenerationInfoReq struct {
	GenerateUUID         string `json:"generate_uuid" binding:"required"`
	PromptTokens         int    `json:"prompt_tokens" binding:"required"`
	CompletionCode       string `json:"completion_code" binding:"required"`
	CompletionCodeTokens int    `json:"completion_code_tokens" binding:"required"`
	CompletionCodeLines  int    `json:"completion_code_lines" binding:"required"`
	CompletionDuration   int    `json:"completion_duration" binding:"required"`
	FinishReason         string `json:"finish_reason" binding:"required"`
	FailureReason        string `json:"failure_reason" binding:"omitempty"`
}

type CodeUpdateGenerationInfoRes struct {
}

type CodeGenerateInsertMongoRequest struct {
	GenerateUUID          string `json:"generateUUID" binding:"required"`
	CodeBeforeCursor      string `json:"codeBeforeCursor" binding:"required"`
	CodeAfterCursor       string `json:"codeAfterCursor" binding:"required"`
	CodeBeforeWithContext string `json:"codeBeforeWithContext" binding:"omitempty"`
}

type CodeGenerateUpdateMongoRequest struct {
	GenerateUUID      string `json:"GenerateUUID" binding:"required"`
	RawCompletionCode string `json:"RawCompletionCode" binding:"required"`
	CompletionCode    string `json:"CompletionCode" binding:"required"`
}

type CodeCreateAcceptStatusRequest struct {
	GenerateUUID string `json:"generateUUID" binding:"required"`
	AcceptStatus string `json:"acceptStatus" binding:"required"`
}

type CodeCreateAcceptStatusResponse struct {
	GenerateUUID string `json:"generateUUID"`
}

type CodeLLmModelDefaultParams struct {
	ModelName            string   `json:"modelName"`
	ModelVersion         string   `json:"modelVersion"`
	TotalMaxTokens       int      `json:"totalMaxTokens"`
	SingleLineMaxTokens  int      `json:"singleLineMaxTokens"`
	MultiLineMaxTokens   int      `json:"multiLineMaxTokens"`
	SingeLineTemperature float64  `json:"singeLineTemperature"`
	MultiLineTemperature float64  `json:"multiLineTemperature"`
	SingleLineTopP       float64  `json:"singleLineTopP"`
	MultiLineTopP        float64  `json:"multiLineTopP"`
	SingleLineStopWords  []string `json:"singleLineStopWords"`
	MultiLineStopWords   []string `json:"multiLineStopWords"`
	IsStream             bool     `json:"isStream"`
	ConnUrls             []string `json:"connUrl"`
}

type ContextFileWithClass struct {
	FilePath           string            `json:"filePath"`
	ClassNamesWithBody map[string]string `json:"classNamesWithBody"`
}

type MaskFuncBodyInput struct {
	LanguageName string
	Context      string
}

type ReGenerateRequest struct {
	GenerateUUID string `json:"generateUUID" binding:"required"`
	GenerateRes  string `json:"generateRes" binding:"required"`
}

type CodeGenerateGetRequest struct {
	GenerateUUID string `json:"generateUUID" binding:"required"`
}

type CodeGenerateGetResponse struct {
	GenerateUUID     string `json:"generateUUID" binding:"required"`
	CodeBeforeCursor string `json:"codeBeforeCursor" binding:"required"`
	CodeAfterCursor  string `json:"codeAfterCursor" binding:"required"`
}

type CodeGenerateWithLLMRequest struct {
	GenerateUUID     string   `json:"generateUUID" binding:"required"`
	CodeBeforeCursor string   `json:"codeBeforeCursor" binding:"required"`
	CodeAfterCursor  string   `json:"codeAfterCursor" binding:"omitempty"`
	ModelVersion     string   `json:"modelVersion"`
	MaxTokens        int      `json:"maxTokens"`
	Temperature      float64  `json:"temperature"`
	TopP             float64  `json:"topP"`
	StopWords        []string `json:"stopWords"`
	IsSingleLine     bool     `json:"isSingleLine"`
	ConnUrls         []string `json:"connUrls"`
	Prompt           string   `json:"prompt"`
}

type CodeGenerateWithLLMResponse struct {
	Completion string `json:"completion"`
}

type RetrievalCrossFileCtxParams struct {
	FlowPercent      int      `json:"flowPercent"`
	MaxCrossFileNums int      `json:"maxCrossFileNums"`
	MinScore         float64  `json:"minScore"`
	ChunkSize        int      `json:"chunkSize"`
	SlideWindowSize  int      `json:"slideWindowSize"`
	SupportCodeLangs []string `json:"supportCodeLangs"`
}

type CodeGeneratePostprocessRequest struct {
	GenerateUUID    string `json:"generateUUID" binding:"required"`
	CodeAfterCursor string `json:"codeAfterCursor" binding:"required"`
	IDEInfo         string `json:"ideInfo" binding:"required"`
	IsSingleLine    bool   `json:"isSingleLine" binding:"required"`
	GenerateType    string `json:"generateType" binding:"required"`
	Completion      string `json:"completion" binding:"required"`
}

type CodeGenerateLLMParams struct {
	ModelName            string   `json:"modelName"`
	ModelVersion         string   `json:"modelVersion"`
	TotalMaxTokens       int      `json:"totalMaxTokens"`
	SingleLineMaxTokens  int      `json:"singleLineMaxTokens"`
	MultiLineMaxTokens   int      `json:"multiLineMaxTokens"`
	SingeLineTemperature float64  `json:"singeLineTemperature"`
	MultiLineTemperature float64  `json:"multiLineTemperature"`
	SingleLineTopP       float64  `json:"singleLineTopP"`
	MultiLineTopP        float64  `json:"multiLineTopP"`
	SingleLineStopWords  []string `json:"singleLineStopWords"`
	MultiLineStopWords   []string `json:"multiLineStopWords"`
	IsStream             bool     `json:"isStream"`
	ConnUrls             []string `json:"connUrls"`
}

type RetrievalCrossFileContextParams struct {
	FlowPercent      int      `json:"flowPercent"`
	MaxCrossFileNums int      `json:"maxCrossFileNums"`
	MinScore         float64  `json:"minScore"`
	ChunkSize        int      `json:"chunkSize"`
	SlideWindowSize  int      `json:"slideWindowSize"`
	SupportCodeLangs []string `json:"supportCodeLangs"`
}

type CodeGenerateRecordInsertDBRequest struct {
	Input                 *CodeGenerateRequest
	LLMParams             *CodeGenerateLLMParams
	GenerateType          string
	IsSingleLine          bool
	CfcNums               int
	CodeBeforeCursor      string
	CodeAfterCursor       string
	CodeBeforeWithContext string
	CodeGenSnippetModels  []*entity.PCodeGenRetrievalSnippetMap
}
