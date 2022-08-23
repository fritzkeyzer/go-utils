# Go Utils

Is a collection of utility packages, written in Golang.

## envutil

- Loading environment variables into tagged config structs.
- Replacing substrings using tagged config structs.
- Pretty printing config structs, while obfuscating secrets.

## logpage

- Create a localhost server that with a log page.
- Implements io.Writer, so is very easily added to std log output (as a typical example)
- Max buffer size.
- Auto page refresh.
- Newest logs displayed at top.

## pretty

- Print
  - Print any object, using JsonString()
- JsonString
  - Convert any object in a pretty, indented and aligned json string.
- FormatJsonString
  - Formats any json string into a prettier, indented and aligned string.

## sshutil

- Client package to interact with remote machines via SSH
- Supports commands and file copying via stp

## stacks

- Parse golang stacktraces into "StackTrace" structs.
- Pretty print "StackTrace" structs in easily readable way.
- Can be used when parsing log files etc.
- Or when generating stacktraces, for debugging purposes or catching panics
  and pretty printing the panic stacktrace.

## stringutil

- Indent
  - Adds indentation to a string - specifically supporting multiline strings.
- IndentAndWrap
  - Wraps long lines and adds indentation to all resulting lines.
  - Optional wrapping character can be specified, to indicate that the line has been wrapped.