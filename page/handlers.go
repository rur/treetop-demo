package page

import (
	"net/http"

	"github.com/rur/treetop"
	"github.com/rur/treetop-demo/site"
)

// ExampleSharedHandler can be used as a handler by multiple pages
func ExampleSharedHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	return "Example"
}
