package inline

import (
	"context"
	"net/http"

	"github.com/rur/treetop"
	"github.com/rur/treetop-demo/page"
	"github.com/rur/treetop-demo/site"
)

type contextKey int

const (
	formDataKey contextKey = iota
)

type handlerWithResources func(*FormData, *site.Env, treetop.Response, *http.Request) interface{}

func bindResources(f handlerWithResources) page.ViewHandlerWithEnv {
	return func(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
		if data, ok := req.Context().Value(formDataKey).(*FormData); ok {
			// previously loaded FormData obtained from request context
			return f(data, env, rsp, req)
		}

		if data, err := loadFormData(rsp, req); err != nil {
			env.ErrorLog.Printf("cookie store error: %s", err)
			pageErrorMessage(rsp, req, "Failed to parse form", http.StatusBadRequest)
			return nil
		} else {
			// FormData loaded from request cookie and added to context
			return f(
				data,
				env,
				rsp,
				req.WithContext(context.WithValue(req.Context(), formDataKey, data)))
		}
	}
}
