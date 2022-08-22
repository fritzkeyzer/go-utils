# Pretty

## Example:

https://go.dev/play/p/xPlvzlt99AN

## Turn this:

```go
fmt.Println(string(debug.Stack()))
```

```
goroutine 4 [running]:
runtime/debug.Stack()
	/opt/homebrew/Cellar/go/1.19/libexec/src/runtime/debug/stack.go:24 +0x64
github.com/fritzkeyzer/go-utils/pretty_test.b()
	/Users/fritzkeyzer/GolandProjects/go-utils/pretty/stack_test.go:44 +0x60
github.com/fritzkeyzer/go-utils/pretty_test.a(...)
	/Users/fritzkeyzer/GolandProjects/go-utils/pretty/stack_test.go:39
github.com/fritzkeyzer/go-utils/pretty_test.TestStack(0x0?)
	/Users/fritzkeyzer/GolandProjects/go-utils/pretty/stack_test.go:22 +0x20
testing.tRunner(0x140000036c0, 0x10089e160)
	/opt/homebrew/Cellar/go/1.19/libexec/src/testing/testing.go:1446 +0x10c
created by testing.(*T).Run
	/opt/homebrew/Cellar/go/1.19/libexec/src/testing/testing.go:1493 +0x300
```

## Into this:

```go
fmt.Println(pretty.Stack())
```

```
stack:
	goroutine 4: running
		pretty       Stack()                              stack.go:64
		pretty_test  b()                                  stack_test.go:43
		pretty_test  a(...)                               stack_test.go:39
		pretty_test  TestStack(0x0?)                      stack_test.go:22
		testing      tRunner(0x140000036c0, 0x10089e160)  testing.go:1446
		testing      (*T).Run                             testing.go:1493
```
