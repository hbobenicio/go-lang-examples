package root

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// Handler is the handler of the index route ("/")
type Handler struct {
	template *template.Template
}

type templateContext struct {
	ShowHeader bool
}

// NewHandler creates a new Root Handler
func NewHandler() Handler {
	return Handler{
		template: template.Must(template.ParseFiles(
			"./root/index.go.html",
			"./root/header.go.html",
			"./root/footer.go.html",
		)),
	}
}

// ServeHTTP implements the http.Handler interface for this Root Handler
func (rh Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	showHeader := r.FormValue("header") != "false"

	tc := templateContext{
		ShowHeader: showHeader,
	}

	if err := rh.template.Execute(w, tc); err != nil {
		fmt.Fprintf(os.Stderr, "error while executing template: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
