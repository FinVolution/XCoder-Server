package ccode

var (
	// CodeSuccess 成功，默认 Http Status Code 为 200
	CodeSuccess        = 100000
	CodeSuccessMessage = "success"

	DBModelAlreadyExistsError         = 400001
	DBModelRecordNotFoundError        = 400003
	DBModelRecordNotFoundErrorMessage = "record not found"
	CodeInternalServerError           = 500000
	CodeUserOperatorError             = 500001

	CodeGenerateCompletionEmpty     = 402001
	CodeGeneratePromptError         = 402002
	CodeGeneratePromptErrorMessage  = "generate prompt error"
	CodeGeneratePredictError        = 402003
	CodeGeneratePredictErrorMessage = "predict error"
)
