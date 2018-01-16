// Package runner defines a basic Web Server to serve as access popint to execute command line script pipes.
package runner

import (
	"bytes"
	"fmt"
	"html/template"
	"os/exec"
	"strings"
)

// Pipe is an instruction for the http handler that runs command line on request with parameters as specified in the URL.
type Pipe struct {
	URL string // /url/{param1}
	Cmd string // cmd {{.param1}}
}

type runner struct {
	tmpl  *template.Template // command line template 'cmd {{.param1}}'
	queue chan string
	code  chan string
}

func newRunner(tmpl string) *runner {
	t := template.Must(template.New("cmd").Parse(tmpl))
	q := make(chan string)
	c := make(chan string)
	r := &runner{t, q, c}
	go r.runningLoop()
	return r
}

func (r *runner) runningLoop() {
	for cmd := range r.queue {
		if len(cmd) == 0 {
			r.code <- "Empty command line."
		}
		binary, params := cmdLineSplit(cmd)
		err := exec.Command(binary, params...).Run()
		if err != nil {
			r.code <- err.Error()
		}
		r.code <- "ok"
	}
}

func (r *runner) executeWith(vars map[string]string) string {
	cmd := r.formatCommand(vars)
	r.queue <- cmd
	return <-r.code
}

func (r *runner) formatCommand(vars map[string]string) string {
	buf := &bytes.Buffer{}
	if err := r.tmpl.Execute(buf, vars); err != nil {
		panic(fmt.Sprintf("can't format runner command, got error %s", err))
	}
	return buf.String()
}

func cmdLineSplit(cmdLine string) (cmd string, args []string) {
	// todo split command and params should respect escaping space in quoted strings.
	// look at "github.com/mattn/go-shellwords"
	line := strings.SplitN(cmdLine, " ", 2)
	cmd = line[0]
	if len(line) > 1 {
		args = strings.Split(line[1], " ")
	}
	return
}
