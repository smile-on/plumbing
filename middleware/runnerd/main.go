package main

import (
	"log"
	"net/http"

	"github.com/smile-on/plumbing/runner"
)

func main() {
	//todo read pipeline settings from file specified as argument 1
	settings := ""
	// settings, _ := ioutil.ReadFile(filename)
	pipes := runner.ParsePipes(settings)
	service := runner.NewHTTPHandler(pipes)
	log.Fatal(http.ListenAndServe(":8080", service))
	//todo main test http://localhost:8080/abc?q=w&j=p
}
