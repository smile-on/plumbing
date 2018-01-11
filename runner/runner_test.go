package runner

import "testing"

func TestFormat_noParams(t *testing.T) {
	r := newRunner("cmdLine {{.param1}}")
	vars := map[string]string{}
	cmd := r.format(vars)
	expected := "cmdLine "
	if cmd != expected {
		t.Errorf("got '%s' want '%s' ", cmd, expected)
	}
}

func TestFormat_noExpectedParams(t *testing.T) {
	r := newRunner("cmdLine {{.param1}}")
	vars := map[string]string{"param2": "abc"}
	cmd := r.format(vars)
	expected := "cmdLine "
	if cmd != expected {
		t.Errorf("got '%s' want '%s' ", cmd, expected)
	}
}

func TestFormat_expected(t *testing.T) {
	r := newRunner("cmdLine {{.param1}}")
	vars := map[string]string{"param1": "abc"}
	cmd := r.format(vars)
	expected := "cmdLine abc"
	if cmd != expected {
		t.Errorf("got '%s' want '%s' ", cmd, expected)
	}
}

func TestExec_simplest(t *testing.T) {
	r := newRunner("echo")
	vars := map[string]string{"param1": "ok"}
	c := r.exec(vars)
	if c != "ok" {
		t.Error("simplest command returned failed exit code:" + c)
	}
}

func TestExec_OneParam(t *testing.T) {
	r := newRunner("echo {{.param1}}")
	vars := map[string]string{"param1": "ok"}
	c := r.exec(vars)
	if c != "ok" {
		t.Error("got failed exit code:" + c)
	}
}

func TestExec_TwoParam(t *testing.T) {
	r := newRunner("echo {{.param1}} {{.p2}}")
	vars := map[string]string{"param1": "ok", "p1": "1", "p2": "2"}
	c := r.exec(vars)
	if c != "ok" {
		t.Error("got failed exit code:" + c)
	}
}
