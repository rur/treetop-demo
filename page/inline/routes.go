package inline

import (
	"github.com/rur/treetop"
	"github.com/rur/treetop-demo/page"
)

// Routes is the plumbing code for page endpoints, templates and handlers
func Routes(hlp page.Helper, exec treetop.ViewExecutor) {

	// Code created by go generate. You should edit the routemap.toml file; DO NOT EDIT.

	inline := treetop.NewView(
		"page/templates/base.html.tmpl",
		page.BaseHandler,
	)

	// [[content]]
	profileForm := inline.NewDefaultSubView(
		"content",
		"page/inline/templates/content/profile-form.html.tmpl",
		hlp.BindEnv(bindResources(profileFormHandler)),
	)

	// [[content.country]]
	countryField := profileForm.NewDefaultSubView(
		"country",
		"page/inline/templates/fields/select.html.tmpl",
		hlp.BindEnv(bindResources(getFormFieldHandler("country"))),
	)

	// [[content.description]]
	descriptionTestArea := profileForm.NewDefaultSubView(
		"description",
		"page/inline/templates/fields/textarea.html.tmpl",
		hlp.BindEnv(bindResources(getFormFieldHandler("description"))),
	)

	// [[content.email]]
	emailField := profileForm.NewDefaultSubView(
		"email",
		"page/inline/templates/fields/email.html.tmpl",
		hlp.BindEnv(bindResources(getFormFieldHandler("email"))),
	)

	// [[content.first-name]]
	firstNameField := profileForm.NewDefaultSubView(
		"first-name",
		"page/inline/templates/fields/input.html.tmpl",
		hlp.BindEnv(bindResources(getFormFieldHandler("firstName"))),
	)

	// [[content.surname]]
	surnameField := profileForm.NewDefaultSubView(
		"surname",
		"page/inline/templates/fields/input.html.tmpl",
		hlp.BindEnv(bindResources(getFormFieldHandler("surname"))),
	)

	// [[nav]]
	inline.NewDefaultSubView(
		"nav",
		"page/templates/nav.html.tmpl",
		treetop.Constant("inline"),
	)

	// [[scripts]]
	inline.NewDefaultSubView(
		"scripts",
		"page/inline/templates/scripts.html.tmpl",
		treetop.Noop,
	)

	// [[styles]]
	inline.NewDefaultSubView(
		"styles",
		"page/inline/templates/styles.html.tmpl",
		treetop.Noop,
	)

	hlp.Handle("/inline",
		exec.NewViewHandler(profileForm).PageOnly())
	hlp.Handle("/inline/country",
		exec.NewViewHandler(countryField).FragmentOnly())
	hlp.Handle("/inline/description",
		exec.NewViewHandler(descriptionTestArea).FragmentOnly())
	hlp.Handle("/inline/email",
		exec.NewViewHandler(emailField).FragmentOnly())
	hlp.Handle("/inline/firstName",
		exec.NewViewHandler(firstNameField).FragmentOnly())
	hlp.Handle("/inline/surname",
		exec.NewViewHandler(surnameField).FragmentOnly())

}
