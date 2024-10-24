// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PCodeGenRetrievalSnippetMap is the golang structure of table p_code_gen_retrieval_snippet_map for DAO operations like Where/Data.
type PCodeGenRetrievalSnippetMap struct {
	g.Meta         `orm:"table:p_code_gen_retrieval_snippet_map, do:true"`
	Updatetime     *gtime.Time //
	Inserttime     *gtime.Time //
	Isactive       interface{} //
	DeleteToken    interface{} //
	Id             interface{} //
	GenerateUuid   interface{} //
	GitRepo        interface{} //
	GitBranch      interface{} //
	SnippetUuid    interface{} //
	SnippetScore   interface{} //
	SnippetRepo    interface{} //
	SnippetPath    interface{} //
	SnippetContent interface{} //
	ProjectName    interface{} //
	ProjectVersion interface{} //
	CreateUser     interface{} //
}
