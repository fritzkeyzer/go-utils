package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/fritzkeyzer/go-utils/logpage"
)

func main() {
	// create and start new logpage
	lp := logpage.New()
	go lp.Host(80, "logs")

	// send all logs to stdErr and logpage
	writer := io.MultiWriter(os.Stderr, lp)
	log.SetOutput(writer)

	log.Println("setup complete. use log.xxx in your application")
	log.Println("open http://localhost/logs in your browser")

	// for example purposes, loop here and log some messages...
	for {
		log.Print("hello")
		log.Print("world")
		log.Print("WARN: ", "spooky warning")
		log.Print("ERROR: ", "scary error")

		time.Sleep(5 * time.Second)
	}
}
