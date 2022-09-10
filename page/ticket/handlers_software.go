package ticket

import (
	"net/http"

	"github.com/rur/treetop"
	"github.com/rur/treetop-demo/site"
)

// -------------------------
// ticket Handlers
// -------------------------

// Ref: new-software-ticket
// Block: form
// Method: GET
// Doc: Form for creating software tickets specifically
func newSoftwareTicketHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	if treetop.IsTemplateRequest(req) {
		// replace existing browser history entry with current URL
		rsp.ReplacePageURL(req.URL.String())
	}
	query := req.URL.Query()
	data := struct {
		AttachmentList interface{}
		FormMessage    interface{}
		Notes          interface{}
		Assignee       interface{}
		Description    string
		IssueType      string
	}{
		AttachmentList: rsp.HandleSubView("attachment-list", req),
		FormMessage:    rsp.HandleSubView("form-message", req),
		Description:    query.Get("description"),
		IssueType:      query.Get("issue-type"),
		Notes:          rsp.HandleSubView("notes", req),
		Assignee:       rsp.HandleSubView("assignee", req),
	}
	return data
}

// Ref: view-software-assignee
// Block: assignee
// Doc: Select multiple users as assignees
func viewSoftwareAssigneeHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		HandlerInfo string
		FindUser    interface{}
	}{
		HandlerInfo: "ticket Page viewSoftwareAssigneeHandler",
		FindUser:    rsp.HandleSubView("find-user", req),
	}
	return data
}

// Ref: find-software-assignee
// Block: find-user
// Method: GET
// Doc: query string to find a user to select
func findSoftwareAssigneeHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		HandlerInfo string
	}{
		HandlerInfo: "ticket Page findSoftwareAssigneeHandler",
	}
	return data
}

// Ref: submit-software-ticket
// Block: form-message
// Method: POST
// Doc: process creation of a new software department ticket
func submitSoftwareTicketHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		HandlerInfo string
	}{
		HandlerInfo: "ticket Page submitSoftwareTicketHandler",
	}
	return data
}
