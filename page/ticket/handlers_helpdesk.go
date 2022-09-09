package ticket

import (
	"net/http"
	"strings"

	"github.com/rur/treetop"
	"github.com/rur/treetop-demo/page/ticket/inputs"
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
	if treetop.IsTemplateRequest(req) {
		// replace existing browser history entry with current URL
		rsp.ReplacePageURL(req.URL.String())
	}
	query := req.URL.Query()
	data := struct {
		ReportedBy     interface{}
		AttachmentList interface{}
		FormMessage    interface{}
		Notes          interface{}
		Description    string
		Urgency        string
	}{
		ReportedBy:     rsp.HandleSubView("reported-by", req),
		AttachmentList: rsp.HandleSubView("attachment-list", req),
		FormMessage:    rsp.HandleSubView("form-message", req),
		Description:    query.Get("description"),
		Urgency:        query.Get("urgency"),
		Notes:          rsp.HandleSubView("notes", req),
	}

	// for mock purposes
	if data.Description == "" {
		data.Description = strings.Join([]string{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do",
			"eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad",
			"minim veniam, quis nostrud exercitation ullamco laboris nisi ut",
			"aliquip ex ea commodo consequat.",
		}, "\n")
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
	query := req.URL.Query()
	data := struct {
		ReportedBy         string
		ReportedByUser     string
		ReportedByCustomer string
		FindUser           interface{}
		CustomerList       []string
		CustomerContact    string
	}{
		ReportedBy: query.Get("reported-by"),
	}
	// Use allow-list for input validation when possible
	switch data.ReportedBy {
	case "team-member":
		// Now parse extra input for this setting
		data.ReportedByUser = query.Get("reported-by-user")
		data.FindUser = rsp.HandleSubView("find-user", req)
	case "customer":
		// Would otherwise be loaded from a customer database
		data.CustomerList = []string{
			"Example Customer 0",
			"Example Customer 1",
			"Example Customer 2",
			"Example Customer 3",
			"Example Customer A",
			"Example Customer B",
			"Example Customer C",
			"Example Customer D",
			"Example Customer E",
		}
		if rBC := query.Get("reported-by-customer"); rBC != "" {
			for _, cst := range data.CustomerList {
				if cst == rBC {
					// accept the input when it matches a known customer
					data.ReportedByCustomer = rBC
					break
				}
			}
		}
		data.CustomerContact = query.Get("customer-contact")
	case "myself":
		// no additional information required
	default:
		// default for empty or unrecognized input
		data.ReportedBy = ""
	}
	return data
}

// Ref: find-helpdesk-reported-by
// Block: find-user
// Method: GET
// Doc: query string to find a user to select
func findHelpdeskReportedByHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	query := req.URL.Query()
	return struct {
		Results     []string
		QueryString string
	}{
		QueryString: query.Get("search-query"),
		Results:     inputs.SearchForUser(query.Get("search-query")),
	}
}
