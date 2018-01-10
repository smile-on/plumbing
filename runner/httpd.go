package runner

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// NewHTTPHandler sets up http handlers according to specifications in pipes.
func NewHTTPHandler(pipes []Pipe) http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/todos/{todoId}", todoShow)
	router.HandleFunc("/", setInfoPage(pipes))
	router.HandleFunc("/json", returnJSON)
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

func todoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoID)
}

func returnJSON(w http.ResponseWriter, r *http.Request) {
	type ReturnCode struct {
		ReturnCode int `json:"returnCode"`
		// output     []byte
	}
	// get input r.Body
	// execute cmd
	code := ReturnCode{-1}
	// return out
	if err := json.NewEncoder(w).Encode(code); err != nil {
		panic(err)
	}
}
