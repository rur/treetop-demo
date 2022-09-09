package ticket

import (
	"github.com/rur/treetop"
	"github.com/rur/treetop-demo/page"
)

// Routes is the plumbing code for page endpoints, templates and handlers
func Routes(hlp page.Helper, exec treetop.ViewExecutor) {

	// Code created by go generate. You should edit the routemap.toml file; DO NOT EDIT.

	ticket := treetop.NewView(
		"page/templates/base.html.tmpl",
		page.BaseHandler,
	)

	// [[content]]
	ticketMainContent := ticket.NewDefaultSubView(
		"content",
		"page/ticket/templates/content/ticket-main-content.html.tmpl",
		hlp.BindEnv(ticketMainContentHandler),
	)

	// [[nav]]
	ticket.NewDefaultSubView(
		"nav",
		"page/templates/nav.html.tmpl",
		treetop.Constant("ticket"),
	)

	// [[scripts]]
	ticket.NewDefaultSubView(
		"scripts",
		"page/ticket/templates/scripts.html.tmpl",
		treetop.Noop,
	)

	// [[styles]]
	ticket.NewDefaultSubView(
		"styles",
		"page/ticket/templates/styles.html.tmpl",
		treetop.Noop,
	)

	hlp.Handle("/ticket",
		exec.NewViewHandler(ticketMainContent).PageOnly())

}
