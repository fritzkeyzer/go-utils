package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/fritzkeyzer/go-utils/logpage"
)

func main() {
	lp := logpage.NewLogPage(80)
	writer := io.MultiWriter(os.Stderr, lp)
	log.SetOutput(writer)

	type Object struct {
		Field      string
		OtherField int
	}

	i := 0
	for {
		log.Print("very long stringvery long stringvery long stringvery long stringvery long stringvery long stringvery long stringvery long stringvery long string")
		log.Print("WARN: ", "hello world")
		log.Print("ERROR: ", "hello world", Object{Field: "hello", OtherField: i})

		i++

		time.Sleep(1 * time.Second)
	}
}
