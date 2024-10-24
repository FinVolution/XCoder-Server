package middleware

import "xcoder/internal/service"

type sMiddleware struct{}

func NewMiddleware() service.IMiddleware {
	return &sMiddleware{}
}

func init() {
	service.RegisterMiddleware(NewMiddleware())
}
