package stacks_test

import (
	"fmt"
	"runtime/debug"
	"testing"

	"github.com/fritzkeyzer/go-utils/pretty"

	"github.com/fritzkeyzer/go-utils/stacks"
)

func TestParse(t *testing.T) {
	a1()

}

func a1() {
	b1()
}

func b1() {
	c1()
}

func c1() {
	st := stacks.ParseStackTrace(string(debug.Stack()))

	fmt.Println(pretty.JsonString(st))
}
