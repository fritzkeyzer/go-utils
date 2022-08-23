package stacks

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/fritzkeyzer/go-utils/stringutil"
)

func (r *GoTrace) string() string {
	str := fmt.Sprintf("goroutine %d: %s\n", r.Number, r.Status)

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 2, 30, 1, ' ', tabwriter.TabIndent)

	fmt.Fprintf(w, "%s \t%s \t%s \t%s\n", "package", "function", "file:line", "snippit")
	fmt.Fprintf(w, "%s \t%s \t%s \t%s\n", "---", "---", "---", "---")
	for _, p := range r.Stack {
		fmt.Fprintf(w, "%s \t%s \t%s:%d \t%s\n", p.Pkg, p.Fn, filepath.Base(p.File), p.Line, p.Snippit)
	}
	w.Flush()
	str += stringutil.Indent(buf.String(), "\t")

	return str
}

// Trace returns a pretty formatted stacktrace string
func Trace(args ...any) string {
	str := string(debug.Stack())
	split := strings.Split(str, "goroutine")

	str = fmt.Sprintf("stack: %s\n", fmt.Sprint(args...))
	for _, s := range split {
		if len(s) == 0 {
			continue
		}
		r := parseRoutine(s)
		r.Stack = r.Stack[2:]

		str += stringutil.Indent(r.string(), "\t")
	}

	return str
}

// PrettyPanic returns a pretty formatted stack trace along with the panic message.
//
// If used within a defer func use:
//
//	deferred = true
func PrettyPanic(panic any, deferred bool) string {
	str := string(debug.Stack())
	split := strings.Split(str, "goroutine")

	str = fmt.Sprintf("panic: %s\n", panic)
	skip := 1
	if deferred {
		skip = 4
	}
	for _, s := range split {
		if len(s) == 0 {
			continue
		}
		r := parseRoutine(s)
		r.Stack = r.Stack[skip:]
		str += stringutil.Indent(r.string(), "\t")
	}

	return str
}

func parseRoutine(str string) GoTrace {
	lines := strings.Split(str, "\n")

	line0 := strings.TrimSpace(lines[0])
	spl := strings.Split(line0, " ")
	status := spl[1][1 : len(spl[1])-2]

	r := GoTrace{
		//Number: spl[0],
		Status: status,
	}

	for i := 1; i+1 < len(lines); i += 2 {
		// pkg
		l1 := lines[i]
		upTo := strings.LastIndex(l1, "(")
		pre := l1
		suffix := ""
		if upTo > 0 {
			pre = l1[:upTo]
			suffix = l1[upTo:]
		}

		preSplit := strings.Split(pre, ".")
		pkgPath := preSplit
		if len(preSplit) > 1 {
			pkgPath = strings.Split(preSplit[len(preSplit)-2], "/")
		}

		pkg := strings.Split(pkgPath[len(pkgPath)-1], " ")

		// file
		l2 := lines[i+1]
		l2Split := strings.Split(strings.TrimSpace(l2), " ")
		//fp := strings.Split(l2Split[0], "/")
		//file := fp[len(fp)-1]
		//file := fp[len(fp)-2] + "/" + fp[len(fp)-1]
		fileWithLineNum := strings.TrimSpace((l2Split[0]))
		fileS := strings.Split(fileWithLineNum, ":")
		file := fileS[0]
		ln, _ := strconv.ParseInt(fileS[1], 10, 64)

		p := TraceLevel{
			Pkg:     pkg[len(pkg)-1],
			File:    file,
			Line:    int(ln),
			Fn:      preSplit[len(preSplit)-1] + suffix,
			Snippit: readSnippit(file, int(ln)),
		}

		r.Stack = append(r.Stack, p)
	}

	return r
}

func readSnippit(filename string, line int) string {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return "<error reading snippit>"
	}

	lines := strings.Split(string(buf), "\n")

	return strings.TrimSpace(lines[line-1])
}
