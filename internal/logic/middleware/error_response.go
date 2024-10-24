package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"xcoder/internal/controller/utils/cerror"
	"xcoder/internal/controller/utils/response"
)

func (s *sMiddleware) ErrResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	ctx := r.Context()
	err := r.GetError()
	if err == nil {
		return
	}

	switch cErr := err.(type) {
	case *cerror.UserErrorButHttpOk:
		g.Log().Warningf(ctx, cErr.Error())
		response.JsonUserErrorButHttpOk(r, cErr)
	case *cerror.UserError:
		g.Log().Warningf(ctx, cErr.Error())
		response.JsonUserError(r, cErr)
	case *cerror.UnauthorizedError:
		response.JsonUnauthorizedError(r, cErr)
	case *cerror.ServerError:
		response.JsonInternalServerError(r, cErr)
	default:
		serverError := cerror.NewInternalServerError(err, "internal server error")
		response.JsonInternalServerError(r, serverError)
	}
}
