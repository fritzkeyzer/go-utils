package stacks_test

import (
	"fmt"
	"runtime/debug"
	"testing"
	"time"

	"github.com/fritzkeyzer/go-utils/stacks"
)

func TestStack(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(stacks.PrettyPanic(p, true))
		}
	}()

	// note

	// fmt.Println(string(debug.Stack()))

	go func() {
		for {
			//a()
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(time.Second)
	a()

	want := `stack:
	1: running
			stack_test.go:27 stack()
		pretty   stack_test.go:23  b(...)
		pretty   stack_test.go:9  a(...)
		pretty   stack_test.go:8   TestStack(0x0?)
		testing  testing.go:1446   tRunner(0x140000036c0, 0x102666020)
		testing  testing.go:1493   Run()`

	_ = want
}

func a() {
	b()
}

func b() {
	//panic("hmm")
	fmt.Println(stacks.Trace("dumping stacktrace here, at timestamp: ", time.Now()))
	fmt.Println(string(debug.Stack()))
}
