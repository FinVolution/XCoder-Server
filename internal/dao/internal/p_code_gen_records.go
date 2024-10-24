// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PCodeGenRecordsDao is the data access object for table p_code_gen_records.
type PCodeGenRecordsDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns PCodeGenRecordsColumns // columns contains all the column names of Table for convenient usage.
}

// PCodeGenRecordsColumns defines and stores column names for table p_code_gen_records.
type PCodeGenRecordsColumns struct {
	Id                   string //
	GenerateUuid         string //
	GenerateType         string //
	IsSingleLine         string //
	CreateUser           string //
	GitRepo              string //
	GitBranch            string //
	CodePath             string //
	CodeLanguage         string //
	IdeInfo              string //
	StartCursorIdx       string //
	PrefixCodeTokens     string //
	SuffixCodeTokens     string //
	ModelName            string //
	ModelVersion         string //
	PromptTokens         string //
	CompletionCode       string //
	CompletionDuration   string //
	CompletionCodeTokens string //
	CompletionCodeLines  string //
	FinishReason         string //
	FailureReason        string //
	Updatetime           string //
	Inserttime           string //
	Isactive             string //
	DeleteToken          string //
	AcceptStatus         string //
	CodeTotalLines       string //
	CrossfileCtxNums     string //
}

// pCodeGenRecordsColumns holds the columns for table p_code_gen_records.
var pCodeGenRecordsColumns = PCodeGenRecordsColumns{
	Id:                   "id",
	GenerateUuid:         "generate_uuid",
	GenerateType:         "generate_type",
	IsSingleLine:         "is_single_line",
	CreateUser:           "create_user",
	GitRepo:              "git_repo",
	GitBranch:            "git_branch",
	CodePath:             "code_path",
	CodeLanguage:         "code_language",
	IdeInfo:              "ide_info",
	StartCursorIdx:       "start_cursor_idx",
	PrefixCodeTokens:     "prefix_code_tokens",
	SuffixCodeTokens:     "suffix_code_tokens",
	ModelName:            "model_name",
	ModelVersion:         "model_version",
	PromptTokens:         "prompt_tokens",
	CompletionCode:       "completion_code",
	CompletionDuration:   "completion_duration",
	CompletionCodeTokens: "completion_code_tokens",
	CompletionCodeLines:  "completion_code_lines",
	FinishReason:         "finish_reason",
	FailureReason:        "failure_reason",
	Updatetime:           "updatetime",
	Inserttime:           "inserttime",
	Isactive:             "isactive",
	DeleteToken:          "delete_token",
	AcceptStatus:         "accept_status",
	CodeTotalLines:       "code_total_lines",
	CrossfileCtxNums:     "crossfile_ctx_nums",
}

// NewPCodeGenRecordsDao creates and returns a new DAO object for table data access.
func NewPCodeGenRecordsDao() *PCodeGenRecordsDao {
	return &PCodeGenRecordsDao{
		group:   "default",
		table:   "p_code_gen_records",
		columns: pCodeGenRecordsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PCodeGenRecordsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PCodeGenRecordsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PCodeGenRecordsDao) Columns() PCodeGenRecordsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PCodeGenRecordsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PCodeGenRecordsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PCodeGenRecordsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
