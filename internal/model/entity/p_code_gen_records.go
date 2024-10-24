// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PCodeGenRecords is the golang structure for table p_code_gen_records.
type PCodeGenRecords struct {
	Id                   uint        `json:"id"                   orm:"id"                     ` //
	GenerateUuid         string      `json:"generateUuid"         orm:"generate_uuid"          ` //
	GenerateType         string      `json:"generateType"         orm:"generate_type"          ` //
	IsSingleLine         int         `json:"isSingleLine"         orm:"is_single_line"         ` //
	CreateUser           string      `json:"createUser"           orm:"create_user"            ` //
	GitRepo              string      `json:"gitRepo"              orm:"git_repo"               ` //
	GitBranch            string      `json:"gitBranch"            orm:"git_branch"             ` //
	CodePath             string      `json:"codePath"             orm:"code_path"              ` //
	CodeLanguage         string      `json:"codeLanguage"         orm:"code_language"          ` //
	IdeInfo              string      `json:"ideInfo"              orm:"ide_info"               ` //
	StartCursorIdx       uint        `json:"startCursorIdx"       orm:"start_cursor_idx"       ` //
	PrefixCodeTokens     uint        `json:"prefixCodeTokens"     orm:"prefix_code_tokens"     ` //
	SuffixCodeTokens     uint        `json:"suffixCodeTokens"     orm:"suffix_code_tokens"     ` //
	ModelName            string      `json:"modelName"            orm:"model_name"             ` //
	ModelVersion         string      `json:"modelVersion"         orm:"model_version"          ` //
	PromptTokens         uint        `json:"promptTokens"         orm:"prompt_tokens"          ` //
	CompletionCode       string      `json:"completionCode"       orm:"completion_code"        ` //
	CompletionDuration   uint        `json:"completionDuration"   orm:"completion_duration"    ` //
	CompletionCodeTokens uint        `json:"completionCodeTokens" orm:"completion_code_tokens" ` //
	CompletionCodeLines  uint        `json:"completionCodeLines"  orm:"completion_code_lines"  ` //
	FinishReason         string      `json:"finishReason"         orm:"finish_reason"          ` //
	FailureReason        string      `json:"failureReason"        orm:"failure_reason"         ` //
	Updatetime           *gtime.Time `json:"updatetime"           orm:"updatetime"             ` //
	Inserttime           *gtime.Time `json:"inserttime"           orm:"inserttime"             ` //
	Isactive             int         `json:"isactive"             orm:"isactive"               ` //
	DeleteToken          string      `json:"deleteToken"          orm:"delete_token"           ` //
	AcceptStatus         string      `json:"acceptStatus"         orm:"accept_status"          ` //
	CodeTotalLines       int         `json:"codeTotalLines"       orm:"code_total_lines"       ` //
	CrossfileCtxNums     int         `json:"crossfileCtxNums"     orm:"crossfile_ctx_nums"     ` //
}
