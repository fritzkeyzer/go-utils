package logs_test

import (
	"github.com/fritzkeyzer/go-utils/logs"
	"github.com/fritzkeyzer/go-utils/logs/page"
	"github.com/fritzkeyzer/go-utils/logs/std"
	"testing"
	"time"
)

func TestNewAppLogger(t *testing.T) {
	log := logs.NewAppLogger(
		std.Out(),
		page.NewServer(80, true),
	)

	type Object struct {
		Field        string
		AnotherField string
	}

	o := Object{
		Field:        "structs can be added as arguments",
		AnotherField: "extra info",
	}

	log.Event("hello world")
	log.Event("this is nice and easy huh? ", o)

	for {
		log.Event("loop here: ", o)

		time.Sleep(2 * time.Second)
	}
}
