package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/smile-on/plumbing/runner"
)

func main() {
	// parse command line arguments
	logFile := flag.String("log", "", "log file name")
	listen := flag.String("listen", ":8080", "listener bind host:port address")
	iniFile := flag.String("ini", "runnerd.ini", "definitions of command line services")
	flag.Parse()
	if *logFile != "" {
		f, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			usage("can't open log file " + *logFile)
		}
		defer f.Close()
		log.SetOutput(f)
	}

	//todo setup logging

	// read pipe's settings
	settings, err := ioutil.ReadFile(*iniFile)
	if err != nil {
		usage("can't read settings in " + *iniFile)
	}
	log.Println("started with settings in " + *iniFile)
	pipes := runner.ParsePipes(string(settings))

	service := runner.NewHTTPHandler(pipes)
	log.Fatal(http.ListenAndServe(*listen, service))
}

func usage(failure string) {
	fmt.Println(failure)
	flag.PrintDefaults()
	os.Exit(2)
}
