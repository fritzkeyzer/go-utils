package main

import (
	"fmt"
	"runtime/debug"

	"github.com/fritzkeyzer/go-utils/stacks"
)

func main() {
	a()

	// you can prettify panics too
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
			fmt.Println("- - - - - - - - - - - - - - - - - - - - ")
			fmt.Println()
			fmt.Println(stacks.PrettyPanic(p, true))
			fmt.Println("- - - - - - - - - - - - - - - - - - - - ")
			fmt.Println()

		}

	}()

	// c() results in a panic
	c()
}

func a() {
	b()
}

func b() {
	fmt.Println(string(debug.Stack()))
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")
	fmt.Println()

	fmt.Println(stacks.Trace())
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")
	fmt.Println()
}

func c() {
	d()
}

func d() {
	panic("some panic occured!")
}
