package cerror

import (
	"fmt"
	"xcoder/internal/controller/utils/ccode"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	RawErr  error
}

type UserError struct {
	CustomError
}

type UserErrorButHttpOk struct {
	CustomError
}

type ServerError struct {
	CustomError
}

type UnauthorizedError struct {
	CustomError
}

func (err CustomError) Error() string {
	return err.Message
}

func (err CustomError) RawError() error {
	return err.RawErr
}

func (e *UserError) Error() string {
	return fmt.Sprintf("%v: %v", e.Message, e.RawErr.Error())
}

func NewInternalServerError(err error, message string) *ServerError {
	return &ServerError{
		CustomError{
			Code:    ccode.CodeInternalServerError,
			Message: message,
			RawErr:  err,
		},
	}
}

func NewUserError(err error, code int, message string) *UserError {
	return &UserError{
		CustomError{
			Code:    code,
			Message: message,
			RawErr:  err,
		},
	}
}

func NewUnauthorizedError(err error, code int, message string) *UnauthorizedError {
	return &UnauthorizedError{CustomError{
		Code:    code,
		Message: message,
		RawErr:  err,
	}}
}

func NewUserErrorButHttpOk(err error, code int, message string) *UserErrorButHttpOk {
	return &UserErrorButHttpOk{
		CustomError{
			Code:    code,
			Message: message,
			RawErr:  err,
		},
	}
}
