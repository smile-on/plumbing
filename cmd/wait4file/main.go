// wait4file checks if specified file exists. If the file is found returns 0.
// If not waits for w seconds and checks again. This check repeats n times.
// In case wait exceeds timeout and file still not exists returns 1.
// In case of IO error or interuption returns 2.
package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	// OK signals specified file exists.
	OK = 0
	// NoFile signals specified file does NOT exist.
	NoFile = 1
	// Exception signals on execution error.
	Exception = 2
)

func main() {
	waitSeconds := flag.Int("w", 1, "seconds to wait before next iteration.")
	trys := flag.Int("n", 1, "number of iterations it trys before reporting file not exist.")
	fileName := flag.String("f", "", "file name to check it exists.")
	flag.Parse()
	if flag.NFlag() == 0 || *fileName == "" {
		flag.PrintDefaults()
		os.Exit(Exception)
	}
	if *waitSeconds <= 0 || *trys <= 0 {
		fmt.Fprintf(os.Stderr, "wait time and number of iterations must be positive integer numbers.\n")
		os.Exit(Exception)
	}

	retCode := checkFileExist(*waitSeconds, *trys, *fileName)
	os.Exit(retCode)
}

func checkFileExist(waitSeconds, trys int, fileName string) int {
	pause := time.Duration(waitSeconds) * time.Second
	for i := 0; ; i++ {
		ok, err := isExist(fileName)
		if ok {
			return OK
		}
		if i == trys {
			if err != nil {
				return Exception
			}
			return NoFile
		}
		time.Sleep(pause)
	}
}

// func Stat(name string) (os.FileInfo, error)
var stat = os.Stat

func isExist(filename string) (bool, error) {
	_, err := stat(filename)
	if err == nil {
		return true, nil // file does exist
	}
	if os.IsNotExist(err) {
		return false, nil // file does not exist
	}
	return false, err

}
