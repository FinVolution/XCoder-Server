package utils

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"strings"
	"time"
	"xcoder/internal/controller/utils/response"
	"xcoder/internal/model/input/chatin"
)

type AnnotatingCodeRequest struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

type AnnotatingCodeResponse struct {
	Code string `json:"code"`
}

var LanguageAnnotatingCodeFlagMap = map[string]string{
	"python":      "#",
	"go":          "//",
	"java":        "//",
	"javascript":  "//",
	"typescript":  "//",
	"kotlin":      "//",
	"groovy":      "//",
	"c":           "//",
	"c++":         "//",
	"c#":          "//",
	"ruby":        "#",
	"perl":        "#",
	"php":         "//",
	"swift":       "//",
	"scala":       "//",
	"rust":        "//",
	"objective-c": "//",
	"erlang":      "%",
	"elixir":      "%",
	"haskell":     "--",
	"lisp":        ";",
}

func AnnotatingCodeWithLanguage(ctx context.Context, req *AnnotatingCodeRequest) (*AnnotatingCodeResponse, error) {
	codeLanguage := strings.ToLower(req.Language)
	flag, ok := LanguageAnnotatingCodeFlagMap[codeLanguage]
	if !ok {
		g.Log().Errorf(ctx, "language not support: %s", req.Language)
		return &AnnotatingCodeResponse{Code: req.Code}, fmt.Errorf("language not support: %s", req.Language)
	}

	codes := strings.Split(req.Code, "\n")
	for i, code := range codes {
		if strings.TrimSpace(code) == "" {
			continue
		}
		codes[i] = flag + " " + code
	}
	req.Code = strings.Join(codes, "\n")

	return &AnnotatingCodeResponse{
		Code: req.Code,
	}, nil
}

func GetStreamingChatReq(ctx context.Context) (r *ghttp.Request) {
	r = g.RequestFromCtx(ctx)

	r.Response.Header().Set("Content-Type", "text/event-stream")
	r.Response.Header().Set("Cache-Control", "no-cache")
	r.Response.Header().Set("Connection", "keep-alive")
	r.Response.Header().Set("Access-Control-Allow-Origin", "*")
	r.Response.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	r.Response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	r.Response.Header().Set("Access-Control-Allow-Credentials", "true")
	r.Response.Header().Set("Access-Control-Max-Age", "86400")

	return
}

func ParseStreamingChatResp(ctx context.Context, r *ghttp.Request, result chan *chatin.ChatResult) {
	for {
		select {
		case res := <-result:
			if res.Error != nil {
				g.Log().Errorf(ctx, "ParseStreamingChatResp error: %v", res.Error)
				err := response.SseSendInternalServerError(ctx, r)
				if err != nil {
					g.Log().Errorf(ctx, "Chat sse response err: %v", err)
					return
				}
			} else {
				data := res.Data
				err := response.SseSendSuccess(ctx, r, data)
				if err != nil {
					g.Log().Errorf(ctx, "Chat sse response err: %v", err)
					return
				}

				r.Response.Flush()

				if data == "[DONE]" {
					g.Log().Infof(ctx, "Chat sse response done ...")
					r.Response.Request.Exit()
					return
				}
			}

		case <-ctx.Done():
			g.Log().Infof(ctx, "Chat sse response done ...")
			r.Response.Request.Exit()
			return

		case <-time.After(15 * time.Second):
			g.Log().Errorf(ctx, "Chat sse response timeout ...")
			err := response.SseSendUserError(ctx, r, 400003, "chat sse response timeout")
			if err != nil {
				g.Log().Errorf(ctx, "Chat sse response response err: %v", err)
				return
			}
			return
		}
	}
}
