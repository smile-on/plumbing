// Package runner defines a basic Web Server to serve as access popint to execute command line script pipes.
package runner

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

// Pipe is an http handler that runs command line on request with parameters specified in the URL.
type Pipe struct {
	URL string // /url/{param1}
	Cmd string // cmd {{.param1}}
}

type runner struct {
	tmpl *template.Template // command line template
}

func newRunner(tmpl string) runner {
	t := template.Must(template.New("cmd").Parse(tmpl))
	return runner{t}
}

func (r *runner) exec(vars map[string]string) (retCode int) {
	cmd := r.format(vars)
	log.Fatal("todo exec " + cmd)
	//todo exec
	return
}

func (r *runner) format(vars map[string]string) string {
	buf := &bytes.Buffer{}
	if err := r.tmpl.Execute(buf, vars); err != nil {
		panic(fmt.Sprintf("can't format runner command, got error %s", err))
	}
	return buf.String()
}
