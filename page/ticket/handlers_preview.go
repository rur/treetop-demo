package ticket

import (
	"net/http"
	"net/url"

	"github.com/rur/treetop"
	"github.com/rur/treetop-demo/page/ticket/inputs"
	"github.com/rur/treetop-demo/site"
)

// -------------------------
// ticket Handlers
// -------------------------

// Ref: issue-preview
// Block: content
// Doc: Content wrapper around preview for different ticket type
func issuePreviewHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		Preview interface{}
		Notes   interface{}
	}{
		Preview: rsp.HandleSubView("preview", req),
		Notes:   rsp.HandleSubView("notes", req),
	}
	return data
}

// Ref: preview-software-ticket
// Block: preview
// Method: GET
// Doc: Show preview of software ticket, no database so take details form query params
func previewSoftwareTicketHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		EditLink string
		Ticket   *inputs.SoftwareTicket
	}{
		// generally this would be loaded from a database but for the demo
		// we are only previewing from URL parameters
		Ticket: inputs.SoftwareTicketFromQuery(req.URL.Query()),
	}
	formURL := url.URL{}
	formURL.Path = "/ticket/software/new"
	formURL.RawQuery = req.URL.Query().Encode()
	data.EditLink = formURL.String()
	return data
}

// Ref: preview-helpdesk-ticket
// Block: preview
// Method: GET
// Doc: Show preview of help desk ticket, no database so take details form query params
func previewHelpdeskTicketHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		EditLink string
		Ticket   *inputs.HelpDeskTicket
	}{
		// generally this would be loaded from a database but for the demo
		// we are only previewing from URL parameters
		Ticket: inputs.HelpdeskTicketFromQuery(req.URL.Query()),
	}
	formURL := url.URL{}
	formURL.Path = "/ticket/helpdesk/new"
	formURL.RawQuery = req.URL.Query().Encode()
	data.EditLink = formURL.String()
	return data
}

// Ref: preview-systems-ticket
// Block: preview
// Method: GET
// Doc: Show preview of systems ticket, no database so take details form query params
func previewSystemsTicketHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		HandlerInfo string
	}{
		HandlerInfo: "ticket Page previewSystemsTicketHandler",
	}
	return data
}
