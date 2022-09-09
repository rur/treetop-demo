package greeter

import (
	"github.com/rur/treetop"
	"github.com/rur/treetop-demo/page"
)

// Routes is the plumbing code for page endpoints, templates and handlers
func Routes(hlp page.Helper, exec treetop.ViewExecutor) {

	// Code created by go generate. You should edit the routemap.toml file; DO NOT EDIT.

	greeter := treetop.NewView(
		"page/templates/base.html.tmpl",
		page.BaseHandler,
	)

	// [[content]]
	greeterFormContent := greeter.NewDefaultSubView(
		"content",
		"page/greeter/templates/content/greeter-form-content.html.tmpl",
		hlp.BindEnv(greeterFormContentHandler),
	)

	// [[content.message]]
	greeterLandingScreen := greeterFormContent.NewDefaultSubView(
		"message",
		"page/greeter/templates/content/message/greeter-landing-screen.html.tmpl",
		treetop.Noop,
	)
	greetingMessage := greeterFormContent.NewSubView(
		"message",
		"page/greeter/templates/content/message/greeting-message.html.tmpl",
		hlp.BindEnv(greetingMessageHandler),
	)

	// [[content.notes]]
	hideNotesFragment := greeterFormContent.NewDefaultSubView(
		"notes",
		"page/greeter/templates/content/notes/hide-notes-fragment.html.tmpl",
		treetop.Noop,
	)
	greeterDemoNotes := greeterFormContent.NewSubView(
		"notes",
		"page/greeter/templates/content/notes/greeter-demo-notes.html.tmpl",
		treetop.Noop,
	)

	// [[nav]]
	greeter.NewDefaultSubView(
		"nav",
		"page/templates/nav.html.tmpl",
		treetop.Constant("greeter"),
	)

	// [[scripts]]
	greeter.NewDefaultSubView(
		"scripts",
		"page/greeter/templates/scripts.html.tmpl",
		treetop.Noop,
	)

	// [[styles]]
	greeter.NewDefaultSubView(
		"styles",
		"page/greeter/templates/styles.html.tmpl",
		treetop.Noop,
	)

	hlp.HandleGET("/greeter",
		exec.NewViewHandler(
			greeterLandingScreen,
			hideNotesFragment,
		))
	hlp.HandleGET("/greeter/greet",
		exec.NewViewHandler(greetingMessage))
	hlp.HandleGET("/greeter/notes",
		exec.NewViewHandler(greeterDemoNotes).FragmentOnly())

}
