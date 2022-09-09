package ticket

import (
	"net/http"

	"github.com/rur/treetop"
	"github.com/rur/treetop-demo/site"
)

// -------------------------
// ticket Handlers
// -------------------------

// Ref: ticket-form-content
// Block: content
// Method: GET
// Doc: Landing page for ticket wizard
func ticketFormContentHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		HandlerInfo string
		Form        interface{}
		Formnotes   interface{}
	}{
		HandlerInfo: "ticket Page ticketFormContentHandler",
		Form:        rsp.HandleSubView("form", req),
		Formnotes:   rsp.HandleSubView("formnotes", req),
	}
	return data
}
