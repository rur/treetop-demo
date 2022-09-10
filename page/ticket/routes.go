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
	ticketFormContent := ticket.NewDefaultSubView(
		"content",
		"page/ticket/templates/content/ticket-form-content.html.tmpl",
		hlp.BindEnv(ticketFormContentHandler),
	)

	// [[content.form]]
	getDepartmentForm := ticketFormContent.NewDefaultSubView(
		"form",
		"page/ticket/templates/content/form/choose-department-message.html.tmpl",
		hlp.BindEnv(getDepartmentFormHandler),
	)
	newHelpdeskTicket := ticketFormContent.NewSubView(
		"form",
		"page/ticket/templates/content/form/new-helpdesk-ticket.html.tmpl",
		hlp.BindEnv(newHelpdeskTicketHandler),
	)

	// [[content.form.attachment-list]]
	newHelpdeskTicket.NewDefaultSubView(
		"attachment-list",
		"page/ticket/templates/content/form/attachment-list/attachment-file-list.html.tmpl",
		hlp.BindEnv(attachmentFileListHandler),
	)
	uploadedHelpdeskFiles := newHelpdeskTicket.NewSubView(
		"attachment-list",
		"page/ticket/templates/content/form/attachment-list/attachment-file-list.html.tmpl",
		hlp.BindEnv(uploadedFilesHandler),
	)

	// [[content.form.form-message]]
	submitHelpDeskTicket := newHelpdeskTicket.NewSubView(
		"form-message",
		"page/ticket/templates/content/form/form-message/submit-help-desk-ticket.html.tmpl",
		hlp.BindEnv(submitHelpDeskTicketHandler),
	)

	// [[content.form.notes]]
	newHelpdeskTicket.NewDefaultSubView(
		"notes",
		"page/ticket/templates/content/form/notes/new-helpdesk-notes.html.tmpl",
		treetop.Noop,
	)

	// [[content.form.reported-by]]
	helpdeskReportedBy := newHelpdeskTicket.NewDefaultSubView(
		"reported-by",
		"page/ticket/templates/content/form/reported-by/helpdesk-reported-by.html.tmpl",
		hlp.BindEnv(helpdeskReportedByHandler),
	)

	// [[content.form.reported-by.find-user]]
	findHelpdeskReportedBy := helpdeskReportedBy.NewDefaultSubView(
		"find-user",
		"page/ticket/templates/content/form/reported-by/find-user/find-helpdesk-reported-by.html.tmpl",
		hlp.BindEnv(findHelpdeskReportedByHandler),
	)

	// [[content.form]]
	newSoftwareTicket := ticketFormContent.NewSubView(
		"form",
		"page/ticket/templates/content/form/new-software-ticket.html.tmpl",
		hlp.BindEnv(newSoftwareTicketHandler),
	)

	// [[content.form.assignee]]
	viewSoftwareAssignee := newSoftwareTicket.NewDefaultSubView(
		"assignee",
		"page/ticket/templates/content/form/assignee/view-software-assignee.html.tmpl",
		hlp.BindEnv(viewSoftwareAssigneeHandler),
	)

	// [[content.form.assignee.find-user]]
	findSoftwareAssignee := viewSoftwareAssignee.NewDefaultSubView(
		"find-user",
		"page/ticket/templates/content/form/assignee/find-user/find-software-assignee.html.tmpl",
		hlp.BindEnv(findSoftwareAssigneeHandler),
	)

	// [[content.form.attachment-list]]
	newSoftwareTicket.NewDefaultSubView(
		"attachment-list",
		"page/ticket/templates/content/form/attachment-list/attachment-file-list.html.tmpl",
		hlp.BindEnv(attachmentFileListHandler),
	)
	uploadedSoftwareFiles := newSoftwareTicket.NewSubView(
		"attachment-list",
		"page/ticket/templates/content/form/attachment-list/attachment-file-list.html.tmpl",
		hlp.BindEnv(uploadedFilesHandler),
	)

	// [[content.form.form-message]]
	submitSoftwareTicket := newSoftwareTicket.NewSubView(
		"form-message",
		"page/ticket/templates/content/form/form-message/submit-software-ticket.html.tmpl",
		hlp.BindEnv(submitSoftwareTicketHandler),
	)

	// [[content.form.notes]]
	newSoftwareTicket.NewDefaultSubView(
		"notes",
		"page/ticket/templates/content/form/notes/new-software-notes.html.tmpl",
		treetop.Noop,
	)

	// [[content]]
	issuePreview := ticket.NewSubView(
		"content",
		"page/ticket/templates/content/issue-preview.html.tmpl",
		hlp.BindEnv(issuePreviewHandler),
	)

	// [[content.notes]]
	issuePreview.NewDefaultSubView(
		"notes",
		"page/ticket/templates/content/notes/preview-notes.html.tmpl",
		treetop.Noop,
	)

	// [[content.preview]]
	previewSoftwareTicket := issuePreview.NewSubView(
		"preview",
		"page/ticket/templates/content/preview/preview-software-ticket.html.tmpl",
		hlp.BindEnv(previewSoftwareTicketHandler),
	)
	previewHelpdeskTicket := issuePreview.NewSubView(
		"preview",
		"page/ticket/templates/content/preview/preview-helpdesk-ticket.html.tmpl",
		hlp.BindEnv(previewHelpdeskTicketHandler),
	)
	previewSystemsTicket := issuePreview.NewSubView(
		"preview",
		"page/ticket/templates/content/preview/preview-systems-ticket.html.tmpl",
		hlp.BindEnv(previewSystemsTicketHandler),
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

	hlp.HandleGET("/ticket",
		exec.NewViewHandler(ticketFormContent))
	hlp.HandleGET("/ticket/get-form",
		exec.NewViewHandler(getDepartmentForm).FragmentOnly())
	hlp.HandleGET("/ticket/helpdesk/new",
		exec.NewViewHandler(newHelpdeskTicket))
	hlp.HandlePOST("/ticket/helpdesk/upload-attachment",
		exec.NewViewHandler(uploadedHelpdeskFiles).FragmentOnly())
	hlp.HandlePOST("/ticket/helpdesk/submit",
		exec.NewViewHandler(submitHelpDeskTicket).FragmentOnly())
	hlp.HandleGET("/ticket/helpdesk/update-reported-by",
		exec.NewViewHandler(helpdeskReportedBy).FragmentOnly())
	hlp.HandleGET("/ticket/helpdesk/find-reported-by",
		exec.NewViewHandler(findHelpdeskReportedBy).FragmentOnly())
	hlp.HandleGET("/ticket/software/new",
		exec.NewViewHandler(newSoftwareTicket))
	hlp.Handle("/ticket/software/update-assignee",
		exec.NewViewHandler(viewSoftwareAssignee).FragmentOnly())
	hlp.HandleGET("/ticket/software/find-assignee",
		exec.NewViewHandler(findSoftwareAssignee).FragmentOnly())
	hlp.HandlePOST("/ticket/software/upload-attachment",
		exec.NewViewHandler(uploadedSoftwareFiles).FragmentOnly())
	hlp.HandlePOST("/ticket/software/submit",
		exec.NewViewHandler(submitSoftwareTicket).FragmentOnly())
	hlp.HandleGET("/ticket/software/preview",
		exec.NewViewHandler(previewSoftwareTicket).PageOnly())
	hlp.HandleGET("/ticket/helpdesk/preview",
		exec.NewViewHandler(previewHelpdeskTicket).PageOnly())
	hlp.HandleGET("/ticket/systems/preview",
		exec.NewViewHandler(previewSystemsTicket).PageOnly())

}
