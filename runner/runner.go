// Package runner defines a basic Web Server to serve as access popint to execute command line script pipes.
package runner

// Pipe is an http handler that runs command line on request with parameters specified in the URL.
type Pipe struct {
	URL string // /url/{param1}
	Cmd string // cmd {{.param1}}
}
