package intro

import (
	"net/http"

	"github.com/rur/treetop"
	"github.com/rur/treetop-demo/site"
)

// -------------------------
// intro Handlers
// -------------------------

// Ref: intro
// Doc: Base HTML template for intro page
func introHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	return struct {
		Content interface{}
		Nav     interface{}
		Scripts interface{}
		Styles  interface{}
	}{
		Content: rsp.HandleSubView("content", req),
		Nav:     rsp.HandleSubView("nav", req),
		Scripts: rsp.HandleSubView("scripts", req),
		Styles:  rsp.HandleSubView("styles", req),
	}
}
