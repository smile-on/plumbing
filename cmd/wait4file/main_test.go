package main

import (
	"errors"
	"os"
	"testing"
	"time"
)

func TestIsExist_yes(t *testing.T) {
	ok, _ := isExist("main_test.go")
	if !ok {
		t.Error("file exists test false.")
	}
}

func TestIsExist_no(t *testing.T) {
	ok, _ := isExist("no_such.file")
	if ok {
		t.Error("file not exist test false.")
	}
}

func TestCheckFile_yes(t *testing.T) {
	guard := time.NewTimer(2 * time.Second)
	defer guard.Stop()
	done := make(chan int)
	defer close(done)

	start := time.Now()
	go func() {
		done <- checkFileExist(1, 1, "main_test.go")
	}()

	select {
	case <-guard.C:
		t.Error("checkFileExist test timeout.")
	case code := <-done:
		if code != OK {
			t.Error("checkFileExist on existed file failed.")
		}
		waitTime := time.Since(start).Seconds()
		if waitTime > 0.01 {
			t.Errorf("checkFileExist on existed file waits for %f expected 0", waitTime)
		}
	}
}

func TestCheckFile_no(t *testing.T) {
	// long test may take few seconds
	guard := time.NewTimer(2 * time.Second)
	defer guard.Stop()
	done := make(chan int)
	defer close(done)

	start := time.Now()
	go func() {
		done <- checkFileExist(1, 1, "no_such.file")
	}()

	select {
	case <-guard.C:
		t.Error("checkFileExist test timeout.")
	case code := <-done:
		if code != NoFile {
			t.Error("checkFileExist on absent file failed.")
		}
		waitTime := time.Since(start).Seconds()
		if waitTime > 1.1 || waitTime < 0.9 {
			t.Errorf("checkFileExist on absent file waits for %f expected 1 second", waitTime)
		}
	}

}

func TestCheckFile_err(t *testing.T) {
	guard := time.NewTimer(2 * time.Second)
	defer guard.Stop()
	done := make(chan int)
	defer close(done)
	stat = statError

	start := time.Now()
	go func() {
		done <- checkFileExist(1, 1, "no_such.file")
	}()

	select {
	case <-guard.C:
		t.Error("checkFileExist test timeout.")
	case code := <-done:
		if code != Exception {
			t.Error("checkFileExist with IO error failed.")
		}
		waitTime := time.Since(start).Seconds()
		if waitTime > 0.01 {
			t.Errorf("checkFileExist with IO error waits for %f expected 0", waitTime)
		}
	}
	stat = os.Stat
}

func statError(name string) (os.FileInfo, error) {
	return nil, errors.New("raised by test mock")
}
