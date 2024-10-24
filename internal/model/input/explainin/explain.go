package explainin

type UserContext struct {
	Type    string `json:"type" binding:"omitempty"`
	Path    string `json:"path" binding:"omitempty"`
	Content string `json:"content" binding:"omitempty"`
}

type ExplainSseGenerateRequest struct {
	CreateUser       string `json:"createUser" binding:"required"`
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
	// 代码信息
	UserCode    string         `json:"userCode" binding:"required"`
	UserContext []*UserContext `json:"userContext" binding:"omitempty"`
}

type ExplainSseGenerateResponse struct{}

type CodeExplainPromptRequest struct {
	RepoName     string `json:"repoName" binding:"required"`
	CodePath     string `json:"codePath" binding:"required"`
	CodeLanguage string `json:"codeLanguage" binding:"required"`
	SelectedCode string `json:"selectCode" binding:"required"`
}

type CodeFuncAcceptStatusUpdateRequest struct {
	ConversationUUID string `json:"conversationUUID" binding:"required"`
	AcceptStatus     string `json:"acceptStatus" binding:"required"`
}

type CodeFuncAcceptStatusUpdateResponse struct {
	ConversationUUID string `json:"conversationUUID" binding:"required"`
}
