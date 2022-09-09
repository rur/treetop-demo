package ticket

import (
	"net/http"

	"github.com/rur/treetop"
	"github.com/rur/treetop-demo/site"
)

// -------------------------
// ticket Handlers
// -------------------------

// Ref: ticket-main-content
// Block: content
// Doc: ticket page content handler, TODO: add rsp.HandleSubView as needed
func ticketMainContentHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		HandlerInfo string
	}{
		HandlerInfo: "ticket Page ticketMainContentHandler",
	}
	return data
}
