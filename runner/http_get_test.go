// Tests covering http GET request to runner pages.
package runner

import "testing"

func TestRunnerPage_smallNoParam(t *testing.T) {
	p := []Pipe{{"/url1", "echo abc"}}
	c := testCase{t,
		NewHTTPHandler(p),
	}
	c.testGetOK("/url1", `abc`)
}

func TestRunnerPage_smallOneParam(t *testing.T) {
	p := []Pipe{{"/url1/{p1}", "echo {{.p1}}"}}
	c := testCase{t,
		NewHTTPHandler(p),
	}
	c.testGetOK("/url1/abc", `abc`)
}
