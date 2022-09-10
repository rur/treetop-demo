package ticket

import (
	"fmt"
	"net/http"
	"net/url"
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

const (
	formMessageInfo = iota
	formMessageWarning
	formMessageError
)

// Ref: submit-help-desk-ticket
// Block: form-message
// Method: POST
// Doc: process creation of a new help desk ticket
func submitHelpDeskTicketHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	var redirected bool
	defer func() {
		if !redirected && treetop.IsTemplateRequest(req) && len(req.PostForm) > 0 {
			// quality-of-life improvement, replace browser URL to include latest
			// form state so that a refresh will preserve inputs
			newURL, _ := url.Parse("/ticket/helpdesk/new")
			q := req.PostForm
			q.Del("file-upload")
			newURL.RawQuery = q.Encode()
			newURL.Fragment = "form-message"
			// replace existing browser history entry with current URL
			rsp.ReplacePageURL(newURL.String())
		}
	}()

	// If all inputs are valid this handler will redirect the web browser
	// either to the newly created ticket or to a blank form.
	//
	// If creation cannot proceed for any reason, this endpoint will render
	// a form message HTML framgent with an alert level: info, warning or error
	data := struct {
		Level           int
		Message         string
		Title           string
		ConfirmCritical bool
	}{
		Level: formMessageInfo,
	}

	if err := req.ParseForm(); err != nil {
		rsp.Status(http.StatusBadRequest)
		data.Level = formMessageError
		data.Title = "Request Error"
		data.Message = "Failed to read form data, try again or contact support."
		return data
	}
	ticket := inputs.HelpdeskTicketFromQuery(req.PostForm)

	// validation rules for creating a new Help Desk ticket
	// NOTE, Do not take client-side validation for granted
	if ticket.Summary == "" {
		data.Level = formMessageWarning
		data.Title = "Missing input"
		data.Message = "Ticket title is required"
		return data
	}
	switch ticket.ReportedBy {
	case "team-member":
		if ticket.ReportedByUser == "" {
			rsp.Status(http.StatusBadRequest)
			data.Level = formMessageWarning
			data.Title = "Missing input"
			data.Message = "Please sepecify which user reported the issue"
			return data
		}
	case "customer":
		if ticket.ReportedByCustomer == "" {
			rsp.Status(http.StatusBadRequest)
			data.Level = formMessageWarning
			data.Title = "Missing input"
			data.Message = "Please sepecify which customer reported the issue"
			return data
		}
	case "":
		rsp.Status(http.StatusBadRequest)
		data.Level = formMessageWarning
		data.Title = "Missing input"
		data.Message = "Please sepecify for whom this issue is being reported"
		return data
	}
	if ticket.Urgency == "" {
		rsp.Status(http.StatusBadRequest)
		data.Level = formMessageWarning
		data.Title = "Invalid input"
		data.Message = fmt.Sprintf("Invalid ticket urgency value '%s'",
			req.PostForm.Get("urgency"))
		return data
	}

	if ticket.Urgency == "critical" && req.PostForm.Get("confirm-critical") != "yes" {
		data.ConfirmCritical = true
		return data
	}

	// ticket is valid redirect to preview endpoint
	previewURL := url.URL{
		Path:     "/ticket/helpdesk/preview",
		RawQuery: ticket.RawQuery(),
	}
	treetop.Redirect(rsp, req, previewURL.String(), http.StatusSeeOther)
	redirected = true
	return nil
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
