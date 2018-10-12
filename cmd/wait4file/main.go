// wait4file checks if specified file exists. If file found returns 0.
// If not waits for t seconds and checks again. This check repeats n times.
// In case wait exceeds timeout and file still not exists returns 1.
package main

import (
	"flag"
	"os"
	"time"
)

const (
	// OK signals specified file exists.
	OK = 0
	// NoFile signals specified file does NOT exist.
	NoFile = 1
	// Err signals on execution error.
	Err = 2
)

func main() {
	waitSeconds := flag.Int("w", 1, "seconds to wait before next iteration.")
	trys := flag.Int("n", 1, "number of iterations it trys before reporting file not exist.")
	fileName := flag.String("f", "", "file name to check it exists.")
	flag.Parse()
	if flag.NFlag() == 0 || *fileName == "" {
		flag.PrintDefaults()
		os.Exit(Err)
	}

	for i := 0; i < *trys; i++ {
		if isExist(*fileName) {
			os.Exit(OK)
		}
		time.Sleep(time.Duration(*waitSeconds) * time.Second)
	}
	os.Exit(NoFile)
}

func isExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false // file does not exist
	}
	return true
}
