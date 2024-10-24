// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PCodeGenRetrievalSnippetMapDao is the data access object for table p_code_gen_retrieval_snippet_map.
type PCodeGenRetrievalSnippetMapDao struct {
	table   string                             // table is the underlying table name of the DAO.
	group   string                             // group is the database configuration group name of current DAO.
	columns PCodeGenRetrievalSnippetMapColumns // columns contains all the column names of Table for convenient usage.
}

// PCodeGenRetrievalSnippetMapColumns defines and stores column names for table p_code_gen_retrieval_snippet_map.
type PCodeGenRetrievalSnippetMapColumns struct {
	Updatetime     string //
	Inserttime     string //
	Isactive       string //
	DeleteToken    string //
	Id             string //
	GenerateUuid   string //
	GitRepo        string //
	GitBranch      string //
	SnippetUuid    string //
	SnippetScore   string //
	SnippetRepo    string //
	SnippetPath    string //
	SnippetContent string //
	ProjectName    string //
	ProjectVersion string //
	CreateUser     string //
}

// pCodeGenRetrievalSnippetMapColumns holds the columns for table p_code_gen_retrieval_snippet_map.
var pCodeGenRetrievalSnippetMapColumns = PCodeGenRetrievalSnippetMapColumns{
	Updatetime:     "updatetime",
	Inserttime:     "inserttime",
	Isactive:       "isactive",
	DeleteToken:    "delete_token",
	Id:             "id",
	GenerateUuid:   "generate_uuid",
	GitRepo:        "git_repo",
	GitBranch:      "git_branch",
	SnippetUuid:    "snippet_uuid",
	SnippetScore:   "snippet_score",
	SnippetRepo:    "snippet_repo",
	SnippetPath:    "snippet_path",
	SnippetContent: "snippet_content",
	ProjectName:    "project_name",
	ProjectVersion: "project_version",
	CreateUser:     "create_user",
}

// NewPCodeGenRetrievalSnippetMapDao creates and returns a new DAO object for table data access.
func NewPCodeGenRetrievalSnippetMapDao() *PCodeGenRetrievalSnippetMapDao {
	return &PCodeGenRetrievalSnippetMapDao{
		group:   "default",
		table:   "p_code_gen_retrieval_snippet_map",
		columns: pCodeGenRetrievalSnippetMapColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PCodeGenRetrievalSnippetMapDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PCodeGenRetrievalSnippetMapDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PCodeGenRetrievalSnippetMapDao) Columns() PCodeGenRetrievalSnippetMapColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PCodeGenRetrievalSnippetMapDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PCodeGenRetrievalSnippetMapDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PCodeGenRetrievalSnippetMapDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
