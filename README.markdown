# Treetop Demo

A demo space for mini web apps that use the [Treetop library](https://github.com/rur/treetop) with HTML template
requests.

#### Links

- [Online Demo](https://treetop-demo.fly.dev/) - hosted by [Fly](https://treetop-demo.fly.dev/)
- Client library for template requests, [treetop client](https://github.com/rur/treetop-client)
- Web app setup [The Good Scaffold](https://github.com/rur/good)
- Styling for components, [Bootstrap 4](https://getbootstrap.com/docs/4.0)

## Greeter App

Given a name, the server responds with a greeting message which is routed to the "message" div on the page.
Not a real scenario, but it's nice to be nice!

It demonstrates: XHR form submit, browser history control and multi-fragment template response.

Review source: [page/greeter/greeter.go](page/greeter/greeter.go)

#### Template Hierarchy

Full page view hierarchy for the `"/greeter/greet"` path.

```
- View("page/templates/base.html.tmpl", github.com/rur/treetop-demo/page.BaseHandler)
  |- content: SubView("content", "page/greeter/templat……greeter-form-content.html.tmpl", github.com/rur/treet……demo/page.Helper.BindEnv.func1)
  |  |- message: SubView("message", "page/greeter/templat……eeter-landing-screen.html.tmpl", github.com/rur/treetop.Noop)
  |  '- notes: SubView("notes", "page/greeter/templat……/hide-notes-fragment.html.tmpl", github.com/rur/treetop.Noop)
  |
  |- nav: SubView("nav", "page/templates/nav.html.tmpl", github.com/rur/treet……demo/page/greeter.Routes.func1)
  |- scripts: SubView("scripts", "page/greeter/templates/scripts.html.tmpl", github.com/rur/treetop.Noop)
  '- styles: SubView("styles", "page/greeter/templates/styles.html.tmpl", github.com/rur/treetop.Noop)
```

## Inline Edit App

Inline editing is commonly used when many different elements can be modified independently.
A user profile is the scenario chosen for the app. The goal is to eliminate the
need for client code that 'knows what's going on'.

Review source: [page/inline/setup.go](page/inline/setup.go)

#### Template Hierarchy

Full page view hierarchy for user profile page

```
- View("page/templates/base.html.tmpl", github.com/rur/treetop-demo/page.BaseHandler)
  |- content: SubView("content", "page/inline/templates/content/profile-form.html.tmpl", github.com/rur/treet……demo/page.Helper.BindEnv.func1)
  | |- country: SubView("country", "page/inline/templates/fields/select.html.tmpl", github.com/rur/treet……demo/page.Helper.BindEnv.func1)
  | |- description: SubView("description", "page/inline/templates/fields/textarea.html.tmpl", github.com/rur/treet……demo/page.Helper.BindEnv.func1)
  | |- email: SubView("email", "page/inline/templates/fields/email.html.tmpl", github.com/rur/treet……demo/page.Helper.BindEnv.func1)
  | |- first-name: SubView("first-name", "page/inline/templates/fields/input.html.tmpl", github.com/rur/treet……demo/page.Helper.BindEnv.func1)
  | '- surname: SubView("surname", "page/inline/templates/fields/input.html.tmpl", github.com/rur/treet……demo/page.Helper.BindEnv.func1)
  |
  |- nav: SubView("nav", "page/templates/nav.html.tmpl", github.com/rur/treetop-demo/page/inline.Routes.func1)
  |- scripts: SubView("scripts", "page/inline/templates/scripts.html.tmpl", github.com/rur/treetop.Noop)
  '- styles: SubView("styles", "page/inline/templates/styles.html.tmpl", github.com/rur/treetop.Noop)
```

## Ticket Wizard App

A multi-stage workflow with branch points and conditions based upon user input.
The forms includes several input components that require server IO:

- temporarily store files,
- search the backend for input values,
- show new options conditioned on a previous one,
- flash messages, and
- redirects.

This app makes use of a 'route map' tool\* to generate the router setup code.

Review source: [page/ticket/routemap.toml](page/ticket/routemap.toml), is used to generate, [page/ticket/routes.go](page/ticket/routes.go)

#### Template Hierarchy

The full page hierarchy for the `"/ticket/helpdesk/new"` endpoint.

```
- View("page/templates/base.html.tmpl", github.com/rur/treetop-demo/page.BaseHandler)
  |- content: SubView("content", "page/ticket/template……/ticket-form-content.html.tmpl", github.com/rur/treet……demo/page.Helper.BindEnv.func1)
  |  '- form: SubView("form", "page/ticket/template……/new-helpdesk-ticket.html.tmpl", github.com/rur/treet……demo/page.Helper.BindEnv.func1)
  |     |- attachment-list: SubView("attachment-list", "page/ticket/template……attachment-file-list.html.tmpl", github.com/rur/treet……demo/page.Helper.BindEnv.func1)
  |     |- form-message: nil
  |     |- notes: SubView("notes", "page/ticket/template……s/new-helpdesk-notes.html.tmpl", github.com/rur/treetop.Noop)
  |     '- reported-by: SubView("reported-by", "page/ticket/template……helpdesk-reported-by.html.tmpl", github.com/rur/treet……demo/page.Helper.BindEnv.func1)
  |        '- find-user: SubView("find-user", "page/ticket/template……helpdesk-reported-by.html.tmpl", github.com/rur/treet……demo/page.Helper.BindEnv.func1)
  |
  |- nav: SubView("nav", "page/templates/nav.html.tmpl", github.com/rur/treetop-demo/page/ticket.Routes.func1)
  |- scripts: SubView("scripts", "page/ticket/templates/scripts.html.tmpl", github.com/rur/treetop.Noop)
  '- styles: SubView("styles", "page/ticket/templates/styles.html.tmpl", github.com/rur/treetop.Noop)
```
