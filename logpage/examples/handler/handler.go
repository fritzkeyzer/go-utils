package main

import (
	"fmt"
	"github.com/fritzkeyzer/go-utils/logpage"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	lp := logpage.New()

	// write to logger like this:
	_, _ = lp.Write([]byte("hello world from lp.Write"))

	// or like this:
	// logpage implements io.Writer interface:
	_, _ = fmt.Fprint(lp, "message from fmt.Fprintf")

	// or like this:
	// by redirecting all logs to stdErr and logpage
	writer := io.MultiWriter(os.Stderr, lp)
	log.SetOutput(writer)
	log.Println("this came from log.Println")

	// logpage implements http.Handler and can be served like this
	// open http://localhost/ in your browser
	http.ListenAndServe(":80", lp)
}
