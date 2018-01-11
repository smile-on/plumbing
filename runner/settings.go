package runner

// ParsePipes instantiates pipes as defined in settings .
// [/url/{param1}]
// cmd {{.param1}}
func ParsePipes(settings string) (pipes []Pipe) {
	//todo implement
	p := []Pipe{{"/url1", "cmd1"}, {"/url2/{param1}", "/usr/bin/cmd2 {{.param1}}"}}
	return p
}
