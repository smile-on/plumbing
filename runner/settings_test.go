package runner

import (
	"testing"
)

func TestParsePipes_ok(t *testing.T) {
	settings := `
# settings
[/echo]
echo

[/echo1/{p1}]
echo {{.p1}}
[/echo2/{p1}/{p2}]
echo {{.p2}}

`
	expected := ""
	pipes := ParsePipes(settings)
	n := len(pipes)
	if n != 3 {
		t.Errorf("want 3 pipes got %d", n)
	}
	if n < 1 {
		return
	}
	expected = "/echo"
	if url := pipes[0].URL; url != expected {
		t.Errorf("url want %s got %s", expected, url)
	}
	expected = "echo"
	if cmd := pipes[0].Cmd; cmd != expected {
		t.Errorf("cmd want %s got %s", expected, cmd)
	}
	if n < 2 {
		return
	}
	expected = "/echo1/{p1}"
	if url := pipes[1].URL; url != expected {
		t.Errorf("url want %s got %s", expected, url)
	}
	expected = "echo {{.p1}}"
	if cmd := pipes[1].Cmd; cmd != expected {
		t.Errorf("cmd want %s got %s", expected, cmd)
	}
	if n < 3 {
		return
	}
	expected = "/echo2/{p1}/{p2}"
	if url := pipes[2].URL; url != expected {
		t.Errorf("url want %s got %s", expected, url)
	}
	expected = "echo {{.p2}}"
	if cmd := pipes[2].Cmd; cmd != expected {
		t.Errorf("cmd want %s got %s", expected, cmd)
	}
}
