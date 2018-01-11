package runner

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// NewHTTPHandler sets up http handlers according to specifications in pipes.
func NewHTTPHandler(pipes []Pipe) http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	for _, p := range pipes {
		router.HandleFunc(p.URL, setRunnerPage(p.Cmd))
	}
	router.HandleFunc("/", setInfoPage(pipes))
	return router
}

//
func setInfoPage(pipes []Pipe) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if pipes == nil {
			fmt.Fprintf(w, "HTTPRunners registered %d", 0)
		} else {
			n := len(pipes)
			fmt.Fprintf(w, "HTTPRunners registered %d", n)
			for _, p := range pipes {
				fmt.Fprintf(w, "\n%q", p)
			}
			//  html.EscapeString(r.URL.Path)
		}
		fmt.Fprintf(w, "\n--- end ---\n")
	}
}

func setRunnerPage(template string) func(http.ResponseWriter, *http.Request) {
	runner := newRunner(template)
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		retCode := runner.exec(vars)
		fmt.Fprintf(w, "%s", retCode)
	}
}
