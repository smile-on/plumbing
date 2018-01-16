package runner

import (
	"regexp"
	"strings"
)

// ParsePipes instantiates pipes as defined in settings .
// [/url/{param1}]
// cmd {{.param1}}
func ParsePipes(settings string) (pipes []Pipe) {
	sectionRegex := regexp.MustCompile(`^\[(.*)\]$`)
	nLine := -1 // line number >0 when inside section
	var p *Pipe

	for _, line := range strings.Split(settings, "\n") {
		line = strings.TrimSpace(line)
		// skip empty lines
		if line == "" {
			continue
		}
		// Skip comments
		if begins := line[0]; begins == ';' || begins == '#' {
			continue
		}
		// parsing sections
		if groups := sectionRegex.FindStringSubmatch(line); groups != nil {
			// found new section
			if nLine > 0 {
				// commit result of previos section
				pipes = append(pipes, *p)
			}
			p = new(Pipe)
			nLine = 0
			p.URL = strings.TrimSpace(groups[1])
		} else if nLine == 0 {
			// first line is a command
			p.Cmd = line
			nLine++
		}

	}
	if nLine > 0 {
		// commit result of last section
		pipes = append(pipes, *p)
	}
	return
}
