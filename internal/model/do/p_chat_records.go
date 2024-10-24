// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PChatRecords is the golang structure of table p_chat_records for DAO operations like Where/Data.
type PChatRecords struct {
	g.Meta              `orm:"table:p_chat_records, do:true"`
	Id                  interface{} // 序号
	ConversationUuid    interface{} // 会话 uuid
	CreateUser          interface{} // 创建人
	GitRepo             interface{} // git 仓库
	GitBranch           interface{} // git 分支
	CodePath            interface{} // 代码路径
	CodeLanguage        interface{} // 代码语言
	IdeInfo             interface{} // ide 版本信息
	ProjectName         interface{} // 插件名称
	ProjectVersion      interface{} // 插件版本
	EngineName          interface{} // 模型名称
	ModelName           interface{} // 模型名称
	ModelVersion        interface{} // 模型版本
	PromptTokens        interface{} // prompt tokens
	CompletionTokens    interface{} // 生成代码 tokens
	TotalTokens         interface{} // 总 tokens
	CompletionCodeLines interface{} // 生成代码行数
	CompletionDuration  interface{} // 生成代码耗时
	FailureReason       interface{} // 会话失败原因
	AcceptStatus        interface{} // 会话是否被采纳
	Updatetime          *gtime.Time // 更新时间
	Inserttime          *gtime.Time // 插入时间
	Isactive            interface{} // 逻辑删除
	DeleteToken         interface{} // 逻辑删除标志
	ConversationType    interface{} // 会话类型
	PromptTmplVersion   interface{} // prompt 版本
	PromptTmplContent   interface{} // prompt 内容
}
