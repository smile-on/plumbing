package runner

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testCase struct {
	t       *testing.T
	handler http.Handler
}

func (c *testCase) testGetOK(url, expected string) {
	c.testHTTP("GET", url, http.StatusOK, "text/plain; charset=utf-8", expected)
}

func (c *testCase) testPostOK(url, expected string) {
	c.testHTTP("POST", url, http.StatusOK, "text/plain; charset=utf-8", expected)
}

func (c *testCase) testHTTP(method, url string, status int, content, body string) {
	req := httptest.NewRequest(method, url, nil)
	w := httptest.NewRecorder()
	c.handler.ServeHTTP(w, req)
	resp := w.Result()
	respondBody, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// Check the status code is what we expect.
	if respondStatus := resp.StatusCode; respondStatus != status {
		c.t.Errorf("handler returned wrong status code: got %v want %v", respondStatus, status)
	}
	// Check the content type is what we expect.
	if respondType := resp.Header.Get("Content-Type"); content != respondType {
		c.t.Errorf("handler returned wrong content type: got '%s' want '%s'", content, respondType)
	}
	// Check the response body is what we expect.
	if string(respondBody) != body {
		c.t.Errorf("handler returned unexpected body: got '%s' want '%s'", respondBody, body)
	}
}

func TestUnknownPage(t *testing.T) {
	c := testCase{t,
		NewHTTPHandler(nil), // no pipeline
	}
	c.testHTTP("GET", "/no-page", http.StatusNotFound, "text/plain; charset=utf-8", `404 page not found
`)
}

func TestInfoPage_Void(t *testing.T) {
	c := testCase{t,
		NewHTTPHandler(nil), // no pipeline
	}
	c.testGetOK("/", `HTTPRunners registered 0
--- end ---
`)
}

func TestInfoPage_Empty(t *testing.T) {
	empty := []Pipe{}
	c := testCase{t,
		NewHTTPHandler(empty),
	}
	c.testGetOK("/", `HTTPRunners registered 0
--- end ---
`)
}

func TestInfoPage_Small(t *testing.T) {
	p := []Pipe{{"/url1", "cmd1"}, {"/url2/{param1}", "/usr/bin/cmd2 {{.param1}}"}}
	c := testCase{t,
		NewHTTPHandler(p),
	}
	c.testGetOK("/", `HTTPRunners registered 2
{"/url1" "cmd1"}
{"/url2/{param1}" "/usr/bin/cmd2 {{.param1}}"}
--- end ---
`)
}
