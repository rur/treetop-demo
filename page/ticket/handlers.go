package ticket

import (
	"net/http"
	"net/url"
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

// Ref: get-department-form
// Block: form
// Method: GET
// Doc: redirect request to correct form or show placeholder message
func getDepartmentFormHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	var (
		redirect *url.URL
		query    = req.URL.Query()
	)
	switch dpt := query.Get("department"); dpt {
	case "helpdesk":
		redirect = mustParseURL("/ticket/helpdesk/new")

	case "software":
		redirect = mustParseURL("/ticket/software/new")

	case "systems":
		redirect = mustParseURL("/ticket/systems/new")
		if len(query["tags"]) == 0 {
			// just for the demo, make sure systems form has at least one tag
			query.Add("tags", "Example Tag 1")
			query.Add("tags", "Example Tag 2")
		}

	default:
		return nil
	}

	redirect.RawQuery = query.Encode()
	http.Redirect(rsp, req, redirect.String(), http.StatusSeeOther)
	return nil
}

// for use with hard coded urls
func mustParseURL(path string) *url.URL {
	u, err := url.Parse(path)
	if err != nil {
		panic(err)
	}
	return u
}
