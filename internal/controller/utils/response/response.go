package response

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"net/http"
	"xcoder/api/common"
	"xcoder/internal/controller/utils/ccode"
	"xcoder/internal/controller/utils/cerror"
)

func JsonSuccess(r *ghttp.Request, data interface{}) {
	r.Response.WriteHeader(http.StatusOK)
	r.Response.WriteJson(&common.Response{
		Code:    ccode.CodeSuccess,
		Message: ccode.CodeSuccessMessage,
		Data:    data,
	})
}

func JsonUserError(r *ghttp.Request, err *cerror.UserError) {
	r.Response.WriteHeader(http.StatusBadRequest)
	r.Response.WriteJsonExit(&common.Response{
		Code:    err.Code,
		Message: fmt.Sprintf("err msg: %s, traceID: %v", err.Error(), gctx.CtxId(r.Context())),
		Data:    nil,
	})
}

func JsonUserErrorButHttpOk(r *ghttp.Request, err *cerror.UserErrorButHttpOk) {
	r.Response.WriteHeader(http.StatusOK)
	r.Response.WriteJsonExit(&common.Response{
		Code:    err.Code,
		Message: fmt.Sprintf("err msg: %s, traceID: %v", err.Error(), gctx.CtxId(r.Context())),
		Data:    nil,
	})
}

func JsonInternalServerError(r *ghttp.Request, err *cerror.ServerError) {
	r.Response.WriteHeader(http.StatusInternalServerError)
	r.Response.WriteJsonExit(&common.Response{
		Code:    err.Code,
		Message: fmt.Sprintf("err msg: %s, traceID: %v", "internal server error", gctx.CtxId(r.Context())),
		Data:    nil,
	})
}

func JsonUnauthorizedError(r *ghttp.Request, err *cerror.UnauthorizedError) {
	r.Response.WriteHeader(http.StatusUnauthorized)
	r.Response.WriteJsonExit(&common.Response{
		Code:    err.Code,
		Message: fmt.Sprintf("err msg: %s, traceId: %v", err.Error(), gctx.CtxId(r.Context())),
		Data:    nil,
	})
}

func SseSendSuccess(ctx context.Context, r *ghttp.Request, data string) error {
	resp := &common.Response{
		Code:    0,
		Message: "",
		Data:    data,
	}
	respJson, err := json.Marshal(resp)
	if err != nil {
		g.Log().Errorf(ctx, "SseSendSuccess json.Marshal err: %v", err)
		return err
	}
	message := fmt.Sprintf("data: %s\n\n", respJson)

	r.Response.WriteHeader(http.StatusOK)
	r.Response.Write([]byte(message))

	return nil
}

func SseSendInternalServerError(ctx context.Context, r *ghttp.Request) error {
	resp := &common.Response{
		Code:    500000,
		Message: "internal server error",
		Data:    "",
	}
	respJson, err := json.Marshal(resp)
	if err != nil {
		g.Log().Errorf(ctx, "SseSendInternalServerError json.Marshal err: %v", err)
		return err
	}

	message := fmt.Sprintf("data: %s\n\n", respJson)
	r.Response.WriteHeader(http.StatusInternalServerError)
	r.Response.WriteExit([]byte(message))

	return nil
}

func SseSendUserError(ctx context.Context, r *ghttp.Request, code int, errMsg string) error {
	resp := &common.Response{
		Code:    code,
		Message: errMsg,
		Data:    "",
	}
	respJson, err := json.Marshal(resp)
	if err != nil {
		g.Log().Errorf(ctx, "SseSendUserError json.Marshal err: %v", err)
		return err
	}
	message := fmt.Sprintf("data: %s\n\n", respJson)

	r.Response.WriteHeader(http.StatusBadRequest)
	r.Response.WriteExit([]byte(message))

	return nil
}
