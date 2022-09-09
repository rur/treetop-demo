package intro

import (
	"github.com/rur/treetop"
	"github.com/rur/treetop-demo/page"
)

// Routes is the plumbing code for page endpoints, templates and handlers
func Routes(hlp page.Helper, exec treetop.ViewExecutor) {

	// Code created by go generate. You should edit the routemap.toml file; DO NOT EDIT.

	intro := treetop.NewView(
		"page/templates/base.html.tmpl",
		hlp.BindEnv(introHandler),
	)

	// [[content]]
	pageLanding := intro.NewDefaultSubView(
		"content",
		"page/intro/templates/content/page-landing.html.tmpl",
		treetop.Noop,
	)

	// [[nav]]
	intro.NewDefaultSubView(
		"nav",
		"page/templates/nav.html.tmpl",
		treetop.Constant("intro"),
	)

	hlp.HandleGET("/",
		exec.NewViewHandler(pageLanding).PageOnly())

}
