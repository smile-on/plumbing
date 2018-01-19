package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

const timeout = 10 * time.Millisecond

func TestAcceptance_noArguments(t *testing.T) {

	// get http server up with defaults
	log.Printf("test: starting HTTP server")
	go main()
	defer shutdownServer()
	log.Printf("test: waiting for HTTP server")
	time.Sleep(timeout)
	client := http.Client{}

	// http client sends GET echo page
	resp, err := client.Get("http://localhost:8080/echo")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	// Check the status code is what we expect.
	if respondStatus := resp.StatusCode; respondStatus != http.StatusOK {
		t.Errorf("server returned wrong status code: got %v want %v", respondStatus, http.StatusOK)
	}
	// Check the content type is what we expect.
	content := "text/plain; charset=utf-8"
	if respondType := resp.Header.Get("Content-Type"); content != respondType {
		t.Errorf("server returned wrong content type: got '%s' want '%s'", respondType, content)
	}
	// check respond body
	respondBody, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if string(respondBody) != "ok" {
		t.Error("unexpected respond from Piped server:" + string(respondBody))
	}

	// http client sends GET incorrect echo page
	resp, err = client.Get("http://localhost:8080/echo/abc")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	// Check the status code is what we expect.
	if respondStatus := resp.StatusCode; respondStatus != http.StatusNotFound {
		t.Errorf("server returned wrong status code: got %v want %v", respondStatus, http.StatusOK)
	}
	// Check the content type is what we expect.
	content = "text/plain; charset=utf-8"
	if respondType := resp.Header.Get("Content-Type"); content != respondType {
		t.Errorf("server returned wrong content type: got '%s' want '%s'", respondType, content)
	}
	// check respond body
	respondBody, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if string(respondBody) != "404 page not found\n" {
		t.Error("unexpected respond from Piped server:'" + string(respondBody) + "'")
	}
}
