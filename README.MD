[![Test Status][workflow.tests.badge]][workflow.tests]

# Go Utils

Is a collection of utility packages, written in Golang.

## envutil

Is a utility package for loading environment variables into tagged structs.

Example struct:
```go
type Config struct {
	EnvName    string   `env:"ENV_NAME" default:"dev"`
	SomeSecret string   `env:"SOME_SECRET" secret:"true"`
	SomeSlice  []string `env:"SOME_SLICE" default:"'hello', 'world'"`
}
```

- `Load(ptr any)`
  - takes a pointer to a struct.
  - For each field tagged with `env:"ENV_NAME"` LoadCfg will attempt to load the environment variable, parse it to the correct type and set the field.
  - If the named env variable isn't found and a 'default' tag is not specified, LoadCfg returns an error.
- `ReplaceVars(input string, ptr any)`
  - Replacing substrings using tagged config structs.
  - Useful for injecting variables into strings.
- `Print(ptr any)`
  - Pretty printing config structs, while obfuscating secret fields.
  - Fields tagged with `secret:"true"` are obfuscated.


## logpage

- Create a localhost server that with a log page.
- Implements io.Writer, so is very easily added to std log output (as a typical example)
- Max buffer size.
- Auto page refresh.
- Newest logs displayed at top.

## pretty

- `Print`
  - Print any object with `Options`
  - Is a wrapper for `String`
- `String`
  - Convert anything to a pretty string with `Options`
- `JsonString`
  - Convert any object in a pretty, indented and aligned json string.
- `FormatJsonString`
  - Formats any json string into a prettier, indented and aligned string.

Example `pretty.Print` output:
```
*pretty_test.Object: -> {
   Field:             "hello world"
   privateField:      "world"
   SomeLongFieldName: "more text"
   NestedObject: {
      Field:             "more stuff"
      AnotherField:      "asdasd"
      SomeLongFieldName: "asdasd"
   }
}
```

## stacks

- Parse golang stacktraces into structs

Example parsed stacktrace:
```
[]stacks.GoTrace: [
   {
      Number: 4
      Status: "running"
      Stack: [
         {
            Fn:   "Stack()"
            Pkg:  "runtime/debug"
            File: "/opt/homebrew/Cellar/go/1.19/libexec/src/runtime/debug/stack.go"
            Line: 24
         },
         {
            Fn:   "TestParse(0x0?)"
            Pkg:  "github.com/fritzkeyzer/go-utils/stacks_test"
            File: "/Users/fritzkeyzer/GolandProjects/go-utils/stacks/parse_test.go"
            Line: 12
         },
         {
            Fn:   "tRunner(0x140000036c0, 0x1028b6168)"
            Pkg:  "testing"
            File: "/opt/homebrew/Cellar/go/1.19/libexec/src/testing/testing.go"
            Line: 1446
         },
         {
            Fn:   "(*T).Run"
            Pkg:  "testing"
            File: "/opt/homebrew/Cellar/go/1.19/libexec/src/testing/testing.go"
            Line: 1493
         },
      ]
   },
]
```

## stringutil

- Indent
  - Adds indentation to a string - specifically supporting multiline strings.
- IndentAndWrap
  - Wraps long lines and adds indentation to all resulting lines.
  - Optional wrapping character can be specified, to indicate that the line has been wrapped.


---
[godoc.badge]: https://godoc.org/github.com/go-fritz/gappfs?status.svg
[godoc]: https://godoc.org/github.com/fritzkeyzer/go-utils
[workflow.tests.badge]: https://github.com/fritzkeyzer/go-utils/workflows/tests/badge.svg
[workflow.tests]: https://github.com/fritzkeyzer/go-utils/actions?query=workflow%3Atests