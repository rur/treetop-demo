package ticket

import (
	"net/http"

	"github.com/rur/treetop"
	"github.com/rur/treetop-demo/site"
)

// -------------------------
// ticket Handlers
// -------------------------

// Ref: new-systems-ticket
// Block: form
// Method: GET
// Doc: Form for creating systems tickets specifically
func newSystemsTicketHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		HandlerInfo    string
		AttachmentList interface{}
		ComponentTags  interface{}
		FormMessage    interface{}
		Notes          interface{}
	}{
		HandlerInfo:    "ticket Page newSystemsTicketHandler",
		AttachmentList: rsp.HandleSubView("attachment-list", req),
		ComponentTags:  rsp.HandleSubView("component-tags", req),
		FormMessage:    rsp.HandleSubView("form-message", req),
		Notes:          rsp.HandleSubView("notes", req),
	}
	return data
}

// Ref: systems-component-tags-input-group
// Block: component-tags
// Doc: Load form input group for the component tags selector
func systemsComponentTagsInputGroupHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		HandlerInfo string
		TagSearch   interface{}
	}{
		HandlerInfo: "ticket Page systemsComponentTagsInputGroupHandler",
		TagSearch:   rsp.HandleSubView("tag-search", req),
	}
	return data
}

// Ref: systems-component-tag-search
// Block: tag-search
// Doc: fuzzy match query to available systems component tags
func systemsComponentTagSearchHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		HandlerInfo string
	}{
		HandlerInfo: "ticket Page systemsComponentTagSearchHandler",
	}
	return data
}

// Ref: submit-systems-ticket
// Block: form-message
// Method: POST
// Doc: process creation of a new systems department ticket
func submitSystemsTicketHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		HandlerInfo string
	}{
		HandlerInfo: "ticket Page submitSystemsTicketHandler",
	}
	return data
}
