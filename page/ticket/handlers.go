package ticket

import (
	"net/http"
	"strings"

	"github.com/rur/treetop"
	"github.com/rur/treetop-demo/site"
)

// -------------------------
// ticket Handlers
// -------------------------

// Ref: ticket-form-content
// Block: content
// Doc: Landing page for ticket wizard
func ticketFormContentHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	query := req.URL.Query()
	data := struct {
		Summary string
		Dept    string
		Form    interface{}
		Notes   interface{}
	}{
		Summary: strings.TrimSpace(query.Get("summary")),
		Form:    rsp.HandleSubView("form", req),
		Dept:    query.Get("department"),
		Notes:   rsp.HandleSubView("notes", req),
	}
	// validate department and redirect if necessary
	switch req.URL.Path {
	case "/ticket/helpdesk/new":
		data.Dept = "helpdesk"

	case "/ticket/software/new":
		data.Dept = "software"

	case "/ticket/systems/new":
		data.Dept = "systems"

	case "/ticket":
		data.Dept = ""
	}
	return data
}
