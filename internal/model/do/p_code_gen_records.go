// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PCodeGenRecords is the golang structure of table p_code_gen_records for DAO operations like Where/Data.
type PCodeGenRecords struct {
	g.Meta               `orm:"table:p_code_gen_records, do:true"`
	Id                   interface{} //
	GenerateUuid         interface{} //
	GenerateType         interface{} //
	IsSingleLine         interface{} //
	CreateUser           interface{} //
	GitRepo              interface{} //
	GitBranch            interface{} //
	CodePath             interface{} //
	CodeLanguage         interface{} //
	IdeInfo              interface{} //
	StartCursorIdx       interface{} //
	PrefixCodeTokens     interface{} //
	SuffixCodeTokens     interface{} //
	ModelName            interface{} //
	ModelVersion         interface{} //
	PromptTokens         interface{} //
	CompletionCode       interface{} //
	CompletionDuration   interface{} //
	CompletionCodeTokens interface{} //
	CompletionCodeLines  interface{} //
	FinishReason         interface{} //
	FailureReason        interface{} //
	Updatetime           *gtime.Time //
	Inserttime           *gtime.Time //
	Isactive             interface{} //
	DeleteToken          interface{} //
	AcceptStatus         interface{} //
	CodeTotalLines       interface{} //
	CrossfileCtxNums     interface{} //
}
