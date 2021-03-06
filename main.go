package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/rur/treetop"

	"os"

	"github.com/rur/treetop-demo/assets"
	"github.com/rur/treetop-demo/greeter"
	"github.com/rur/treetop-demo/inline"
	"github.com/rur/treetop-demo/ticket"
)

var portReg = regexp.MustCompile(`^\d+$`)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "https://golang.org/favicon.ico", http.StatusSeeOther)
	})
	mux.Handle("/treetop.js", treetop.ServeClientLibrary)

	devMode := true
	port := "3000"
	for i, arg := range os.Args[1:] {
		if arg == "--prod" {
			devMode = false
		} else if portReg.MatchString(arg) {
			port = arg
		} else {
			log.Fatalf("Invalid argument [%d] '%s'", i, arg)
		}
	}

	// Register routes for demo apps
	greeter.Setup(mux, devMode)
	inline.Setup(mux, devMode)
	ticket.Setup(mux, devMode)

	infoSetup(mux, devMode)

	fmt.Printf("serving on http://0.0.0.0:%s/\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

// infoSetup will create template, handlers and bind to routes for the demo landing page
func infoSetup(mux *http.ServeMux, devMode bool) {
	// define handler for home page
	var exec treetop.ViewExecutor = &treetop.FileExecutor{
		KeyedString: map[string]string{
			"local://base.html":     assets.BaseHTML,
			"local://nav-home.html": assets.NavHTML(assets.IntroNav),
			"local://nav-more.html": assets.NavHTML(assets.MoreNav),
			"local://nav.html":      assets.NavHTML(assets.NoPage),
			"local://notfound.html": `
				<div class="text-center">
					<hr class="mb-3">
					<h3>404 Not Found</h3>
					<p>No page was found for this path.</p>
					<p>You can file a (real) issue here <a href="https://github.com/rur/treetop/issues">treetop/issues</a></p>
				</div>
				`,
		},
	}
	if devMode {
		exec = &treetop.DeveloperExecutor{
			ViewExecutor: exec,
		}
	}

	base := treetop.
		NewView("local://base.html", treetop.Noop)

	mux.Handle("/more", exec.NewViewHandler(
		base.NewSubView("content", "more.html", treetop.Noop),
		base.NewSubView("nav", "local://nav-more.html", treetop.Noop),
	).PageOnly())

	homeHandler := exec.NewViewHandler(
		base.NewSubView("content", "intro.html", treetop.Noop),
		base.NewSubView("nav", "local://nav-home.html", treetop.Noop),
	).PageOnly()

	notFoundHandler := exec.NewViewHandler(
		base.NewSubView("content", "local://notfound.html", treetop.Noop),
		base.NewSubView("nav", "local://nav.html", treetop.Noop),
	).PageOnly()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		switch req.URL.Path {
		case "/", "":
			homeHandler.ServeHTTP(w, req)
		default:
			notFoundHandler.ServeHTTP(w, req)
		}
	})

	if errs := exec.FlushErrors(); len(errs) > 0 {
		log.Fatalf("Error(s) loading demo templates:\n%s", errs)
	}
}
