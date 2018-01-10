package runner

import (
	"bytes"
	"html/template"
)

// ParsePipes instantiates pipes as defined in settings .
// [/url/{param1}]
// cmd {{.param1}}
func ParsePipes(settings string) (pipes []Pipe) {
	//todo implement
	p := []Pipe{{"/url1", "cmd1"}, {"/url2/{param1}", "/usr/bin/cmd2 {{.param1}}"}}
	return p
}

/*
m := map[string]interface{}{
	"name": "John", "age": 47,
}

tmplText :=  `Do {{.todoId}}.`
*/

/* usage renderTemplate
tmpl := Tprintf(tmplText, m)
fmt.Println(tmpl)
*/

// Tprintf passed template string is formatted usign its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func Tprintf(tmpl string, data map[string]interface{}) string {
	t := template.Must(template.New("sql").Parse(tmpl))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		return ""
	}
	return buf.String()
}
