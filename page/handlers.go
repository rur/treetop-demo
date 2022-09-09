package page

import (
	"net/http"

	"github.com/rur/treetop"
)

// BaseHandler can be used as a handler by multiple pages
func BaseHandler(rsp treetop.Response, req *http.Request) interface{} {
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
