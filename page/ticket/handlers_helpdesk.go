package ticket

import (
	"net/http"

	"github.com/rur/treetop"
	"github.com/rur/treetop-demo/site"
)

// -------------------------
// ticket Handlers
// -------------------------

// Ref: new-helpdesk-ticket
// Block: form
// Method: GET
// Doc: Form for creating helpdesk tickets specifically
func newHelpdeskTicketHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		HandlerInfo    string
		AttachmentList interface{}
		FormMessage    interface{}
		ReportedBy     interface{}
		Notes          interface{}
	}{
		HandlerInfo:    "ticket Page newHelpdeskTicketHandler",
		AttachmentList: rsp.HandleSubView("attachment-list", req),
		FormMessage:    rsp.HandleSubView("form-message", req),
		ReportedBy:     rsp.HandleSubView("reported-by", req),
		Notes:          rsp.HandleSubView("notes", req),
	}
	return data
}

// Ref: helpdesk-attachment-file-list
// Block: attachment-list
// Doc: Default helpdesk attachment file list template handler, parse file info from query string
func helpdeskAttachmentFileListHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		HandlerInfo string
	}{
		HandlerInfo: "ticket Page helpdeskAttachmentFileListHandler",
	}
	return data
}

// Ref: uploaded-helpdesk-files
// Block: attachment-list
// Method: POST
// Doc: Load a list of uploaded files, save to storage and return metadata to the form
func uploadedHelpdeskFilesHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		HandlerInfo string
	}{
		HandlerInfo: "ticket Page uploadedHelpdeskFilesHandler",
	}
	return data
}

// Ref: submit-help-desk-ticket
// Block: form-message
// Method: POST
// Doc: process creation of a new help desk ticket
func submitHelpDeskTicketHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		HandlerInfo string
	}{
		HandlerInfo: "ticket Page submitHelpDeskTicketHandler",
	}
	return data
}

// Ref: helpdesk-reported-by
// Block: reported-by
// Method: GET
// Doc: Options for notifying help desk of who reported the issue
func helpdeskReportedByHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		HandlerInfo string
		FindUser    interface{}
	}{
		HandlerInfo: "ticket Page helpdeskReportedByHandler",
		FindUser:    rsp.HandleSubView("find-user", req),
	}
	return data
}

// Ref: find-helpdesk-reported-by
// Block: find-user
// Method: GET
// Doc: query string to find a user to select
func findHelpdeskReportedByHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		HandlerInfo string
	}{
		HandlerInfo: "ticket Page findHelpdeskReportedByHandler",
	}
	return data
}
