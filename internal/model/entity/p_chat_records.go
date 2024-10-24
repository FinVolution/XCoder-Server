// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PChatRecords is the golang structure for table p_chat_records.
type PChatRecords struct {
	Id                  uint        `json:"id"                  orm:"id"                    ` // 序号
	ConversationUuid    string      `json:"conversationUuid"    orm:"conversation_uuid"     ` // 会话 uuid
	CreateUser          string      `json:"createUser"          orm:"create_user"           ` // 创建人
	GitRepo             string      `json:"gitRepo"             orm:"git_repo"              ` // git 仓库
	GitBranch           string      `json:"gitBranch"           orm:"git_branch"            ` // git 分支
	CodePath            string      `json:"codePath"            orm:"code_path"             ` // 代码路径
	CodeLanguage        string      `json:"codeLanguage"        orm:"code_language"         ` // 代码语言
	IdeInfo             string      `json:"ideInfo"             orm:"ide_info"              ` // ide 版本信息
	ProjectName         string      `json:"projectName"         orm:"project_name"          ` // 插件名称
	ProjectVersion      string      `json:"projectVersion"      orm:"project_version"       ` // 插件版本
	EngineName          string      `json:"engineName"          orm:"engine_name"           ` // 模型名称
	ModelName           string      `json:"modelName"           orm:"model_name"            ` // 模型名称
	ModelVersion        string      `json:"modelVersion"        orm:"model_version"         ` // 模型版本
	PromptTokens        uint        `json:"promptTokens"        orm:"prompt_tokens"         ` // prompt tokens
	CompletionTokens    uint        `json:"completionTokens"    orm:"completion_tokens"     ` // 生成代码 tokens
	TotalTokens         uint        `json:"totalTokens"         orm:"total_tokens"          ` // 总 tokens
	CompletionCodeLines uint        `json:"completionCodeLines" orm:"completion_code_lines" ` // 生成代码行数
	CompletionDuration  uint        `json:"completionDuration"  orm:"completion_duration"   ` // 生成代码耗时
	FailureReason       string      `json:"failureReason"       orm:"failure_reason"        ` // 会话失败原因
	AcceptStatus        string      `json:"acceptStatus"        orm:"accept_status"         ` // 会话是否被采纳
	Updatetime          *gtime.Time `json:"updatetime"          orm:"updatetime"            ` // 更新时间
	Inserttime          *gtime.Time `json:"inserttime"          orm:"inserttime"            ` // 插入时间
	Isactive            int         `json:"isactive"            orm:"isactive"              ` // 逻辑删除
	DeleteToken         string      `json:"deleteToken"         orm:"delete_token"          ` // 逻辑删除标志
	ConversationType    string      `json:"conversationType"    orm:"conversation_type"     ` // 会话类型
	PromptTmplVersion   string      `json:"promptTmplVersion"   orm:"prompt_tmpl_version"   ` // prompt 版本
	PromptTmplContent   string      `json:"promptTmplContent"   orm:"prompt_tmpl_content"   ` // prompt 内容
}
