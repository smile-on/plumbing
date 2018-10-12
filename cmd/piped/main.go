package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/smile-on/plumbing/runner"
)

func main() {
	// parse command line arguments
	logFile := flag.String("log", "", "log file name")
	listen := flag.String("listen", ":8080", "listener bind host:port address")
	iniFile := flag.String("ini", "piped.ini", "definitions of command line services")
	flag.Parse()

	//todo setup logging
	// "" = disabled; 	stdout;		file
	if *logFile != "" {
		f, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			usage("can't open log file " + *logFile)
		}
		defer f.Close()
		log.SetOutput(f)
	}

	// read pipe's settings
	settings, err := ioutil.ReadFile(*iniFile)
	if err != nil {
		usage("can't read settings in " + *iniFile)
	}
	log.Println("started with settings in " + *iniFile)
	pipes := runner.ParsePipes(string(settings))

	// run server
	server.Handler = runner.NewHTTPHandler(pipes)
	server.Addr = *listen
	if err := server.ListenAndServe(); err != nil {
		log.Printf("stopped by %s\n", err)
	}
	log.Printf("done\n")
}

func usage(failure string) {
	fmt.Println(failure)
	flag.PrintDefaults()
	os.Exit(2)
}

var server http.Server

func shutdownServer() {
	log.Printf("stopping HTTP server")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Print("failure in shutting down the server gracefully.", err)
		server.Close()
	}
}
