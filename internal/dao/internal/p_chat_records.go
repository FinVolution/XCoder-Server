// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PChatRecordsDao is the data access object for table p_chat_records.
type PChatRecordsDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns PChatRecordsColumns // columns contains all the column names of Table for convenient usage.
}

// PChatRecordsColumns defines and stores column names for table p_chat_records.
type PChatRecordsColumns struct {
	Id                  string // 序号
	ConversationUuid    string // 会话 uuid
	CreateUser          string // 创建人
	GitRepo             string // git 仓库
	GitBranch           string // git 分支
	CodePath            string // 代码路径
	CodeLanguage        string // 代码语言
	IdeInfo             string // ide 版本信息
	ProjectName         string // 插件名称
	ProjectVersion      string // 插件版本
	EngineName          string // 模型名称
	ModelName           string // 模型名称
	ModelVersion        string // 模型版本
	PromptTokens        string // prompt tokens
	CompletionTokens    string // 生成代码 tokens
	TotalTokens         string // 总 tokens
	CompletionCodeLines string // 生成代码行数
	CompletionDuration  string // 生成代码耗时
	FailureReason       string // 会话失败原因
	AcceptStatus        string // 会话是否被采纳
	Updatetime          string // 更新时间
	Inserttime          string // 插入时间
	Isactive            string // 逻辑删除
	DeleteToken         string // 逻辑删除标志
	ConversationType    string // 会话类型
	PromptTmplVersion   string // prompt 版本
	PromptTmplContent   string // prompt 内容
}

// pChatRecordsColumns holds the columns for table p_chat_records.
var pChatRecordsColumns = PChatRecordsColumns{
	Id:                  "id",
	ConversationUuid:    "conversation_uuid",
	CreateUser:          "create_user",
	GitRepo:             "git_repo",
	GitBranch:           "git_branch",
	CodePath:            "code_path",
	CodeLanguage:        "code_language",
	IdeInfo:             "ide_info",
	ProjectName:         "project_name",
	ProjectVersion:      "project_version",
	EngineName:          "engine_name",
	ModelName:           "model_name",
	ModelVersion:        "model_version",
	PromptTokens:        "prompt_tokens",
	CompletionTokens:    "completion_tokens",
	TotalTokens:         "total_tokens",
	CompletionCodeLines: "completion_code_lines",
	CompletionDuration:  "completion_duration",
	FailureReason:       "failure_reason",
	AcceptStatus:        "accept_status",
	Updatetime:          "updatetime",
	Inserttime:          "inserttime",
	Isactive:            "isactive",
	DeleteToken:         "delete_token",
	ConversationType:    "conversation_type",
	PromptTmplVersion:   "prompt_tmpl_version",
	PromptTmplContent:   "prompt_tmpl_content",
}

// NewPChatRecordsDao creates and returns a new DAO object for table data access.
func NewPChatRecordsDao() *PChatRecordsDao {
	return &PChatRecordsDao{
		group:   "default",
		table:   "p_chat_records",
		columns: pChatRecordsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PChatRecordsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PChatRecordsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PChatRecordsDao) Columns() PChatRecordsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PChatRecordsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PChatRecordsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PChatRecordsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
