package unit_testin

type UserContext struct {
	Type    string `json:"type" binding:"omitempty"`
	Path    string `json:"path" binding:"omitempty"`
	Content string `json:"content" binding:"omitempty"`
}

type UnitTestSseGenerateReq struct {
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
	// 测试用例信息
	Framework   string         `json:"unitTestFramework" binding:"omitempty"`
	UserText    string         `json:"userText" binding:"omitempty"`
	UserCode    string         `json:"userCode" binding:"required"`
	UserContext []*UserContext `json:"userContext" binding:"omitempty"`
}

type UnitTestSseGenerateResp struct{}

type GenerateUTPromptReq struct {
	RepoName        string         `json:"repoName" binding:"required"`
	CodePath        string         `json:"codePath" binding:"required"`
	CodeLanguage    string         `json:"codeLanguage" binding:"required"`
	SharedContexts  []*UserContext `json:"sharedContexts" binding:"required"`
	SelectedCode    string         `json:"selectedCode" binding:"required"`
	UserInstruction string         `json:"userInstruction" binding:"omitempty"`
	Framework       string         `json:"framework" binding:"required"`
}
