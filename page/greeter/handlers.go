package greeter

import (
	"net/http"

	"github.com/rur/treetop"
	"github.com/rur/treetop-demo/site"
)

// -------------------------
// greeter Handlers
// -------------------------

// Ref: greeter-form-content
// Block: content
// Doc: static form layout
func greeterFormContentHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		Value   string
		Message interface{}
		Notes   interface{}
	}{
		Value:   req.URL.Query().Get("name"),
		Message: rsp.HandleSubView("message", req),
		Notes:   rsp.HandleSubView("notes", req),
	}
	return data
}

// Ref: greeting-message
// Block: message
// Method: GET
// Doc: show the greeting message
func greetingMessageHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	query := req.URL.Query()
	who := query.Get("name")
	if who == "" {
		who = "World"
	}
	return who
}
