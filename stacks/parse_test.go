package stacks_test

import (
	"github.com/fritzkeyzer/go-utils/pretty"
	"github.com/fritzkeyzer/go-utils/stacks"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	stack := `goroutine 4 [running]:
runtime/debug.Stack()
	/opt/homebrew/Cellar/go/1.19/libexec/src/runtime/debug/stack.go:24 +0x64
github.com/fritzkeyzer/go-utils/stacks_test.TestParse(0x0?)
	/Users/fritzkeyzer/GolandProjects/go-utils/stacks/parse_test.go:12 +0x1c
testing.tRunner(0x140000036c0, 0x1028b6168)
	/opt/homebrew/Cellar/go/1.19/libexec/src/testing/testing.go:1446 +0x10c
created by testing.(*T).Run
	/opt/homebrew/Cellar/go/1.19/libexec/src/testing/testing.go:1493 +0x300`

	got := stacks.Parse(stack)

	want := []stacks.GoTrace{
		{
			Number: 4,
			Status: "running",
			Stack: []stacks.TracePoint{
				{
					Fn:   "Stack()",
					Pkg:  "runtime/debug",
					File: "/opt/homebrew/Cellar/go/1.19/libexec/src/runtime/debug/stack.go",
					Line: 24,
				},
				{
					Fn:   "TestParse(0x0?)",
					Pkg:  "github.com/fritzkeyzer/go-utils/stacks_test",
					File: "/Users/fritzkeyzer/GolandProjects/go-utils/stacks/parse_test.go",
					Line: 12,
				},
				{
					Fn:   "tRunner(0x140000036c0, 0x1028b6168)",
					Pkg:  "testing",
					File: "/opt/homebrew/Cellar/go/1.19/libexec/src/testing/testing.go",
					Line: 1446,
				},
				{
					Fn:   "(*T).Run",
					Pkg:  "testing",
					File: "/opt/homebrew/Cellar/go/1.19/libexec/src/testing/testing.go",
					Line: 1493,
				},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got != want:\ngot: %s\nwant: %s", pretty.String(want), pretty.String(got))
	}

	pretty.Print(got)
}

func ExampleParse() {
	// stack string could come from anywhere, panics in a log file or from debug.Stack()
	stack := `goroutine 4 [running]:
runtime/debug.Stack()
	/opt/homebrew/Cellar/go/1.19/libexec/src/runtime/debug/stack.go:24 +0x64
github.com/fritzkeyzer/go-utils/stacks_test.TestParse(0x0?)
	/Users/fritzkeyzer/GolandProjects/go-utils/stacks/parse_test.go:12 +0x1c
testing.tRunner(0x140000036c0, 0x1028b6168)
	/opt/homebrew/Cellar/go/1.19/libexec/src/testing/testing.go:1446 +0x10c
created by testing.(*T).Run
	/opt/homebrew/Cellar/go/1.19/libexec/src/testing/testing.go:1493 +0x300`

	parsed := stacks.Parse(stack)

	pretty.Print(parsed)
	// Output:
	// []stacks.GoTrace: [
	//    {
	//       Number: 4
	//       Status: "running"
	//       Stack:  [
	//          {
	//             Fn:   "Stack()"
	//             Pkg:  "runtime/debug"
	//             File: "/opt/homebrew/Cellar/go/1.19/libexec/src/runtime/debug/stack.go"
	//             Line: 24
	//          },
	//          {
	//             Fn:   "TestParse(0x0?)"
	//             Pkg:  "github.com/fritzkeyzer/go-utils/stacks_test"
	//             File: "/Users/fritzkeyzer/GolandProjects/go-utils/stacks/parse_test.go"
	//             Line: 12
	//          },
	//          {
	//             Fn:   "tRunner(0x140000036c0, 0x1028b6168)"
	//             Pkg:  "testing"
	//             File: "/opt/homebrew/Cellar/go/1.19/libexec/src/testing/testing.go"
	//             Line: 1446
	//          },
	//          {
	//             Fn:   "(*T).Run"
	//             Pkg:  "testing"
	//             File: "/opt/homebrew/Cellar/go/1.19/libexec/src/testing/testing.go"
	//             Line: 1493
	//          },
	//       ]
	//    },
	// ]
}
