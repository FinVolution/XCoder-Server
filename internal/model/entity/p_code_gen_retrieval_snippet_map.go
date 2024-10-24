// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PCodeGenRetrievalSnippetMap is the golang structure for table p_code_gen_retrieval_snippet_map.
type PCodeGenRetrievalSnippetMap struct {
	Updatetime     *gtime.Time `json:"updatetime"     orm:"updatetime"      ` //
	Inserttime     *gtime.Time `json:"inserttime"     orm:"inserttime"      ` //
	Isactive       int         `json:"isactive"       orm:"isactive"        ` //
	DeleteToken    string      `json:"deleteToken"    orm:"delete_token"    ` //
	Id             uint        `json:"id"             orm:"id"              ` //
	GenerateUuid   string      `json:"generateUuid"   orm:"generate_uuid"   ` //
	GitRepo        string      `json:"gitRepo"        orm:"git_repo"        ` //
	GitBranch      string      `json:"gitBranch"      orm:"git_branch"      ` //
	SnippetUuid    string      `json:"snippetUuid"    orm:"snippet_uuid"    ` //
	SnippetScore   float64     `json:"snippetScore"   orm:"snippet_score"   ` //
	SnippetRepo    string      `json:"snippetRepo"    orm:"snippet_repo"    ` //
	SnippetPath    string      `json:"snippetPath"    orm:"snippet_path"    ` //
	SnippetContent string      `json:"snippetContent" orm:"snippet_content" ` //
	ProjectName    string      `json:"projectName"    orm:"project_name"    ` //
	ProjectVersion string      `json:"projectVersion" orm:"project_version" ` //
	CreateUser     string      `json:"createUser"     orm:"create_user"     ` //
}
