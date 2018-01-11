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
