package stacks

import (
	"fmt"
	"strconv"
	"strings"
)

type GoTrace struct {
	Number int
	Status string
	Stack  []TraceLevel
}

type TraceLevel struct {
	Fn      string
	Pkg     string
	File    string
	Line    int
	Snippit string
}

// ParseStackTrace takes a go stack trace string and returns a
// list of traces, for each goroutine.
func ParseStackTrace(trace string) []GoTrace {
	split := strings.Split(trace, "\ngoroutine ")
	ret := make([]GoTrace, len(split))
	for _, s := range split {
		if len(s) == 0 {
			continue
		}
		r := parseRoutine2(s)
		ret = append(ret, r)
	}

	return ret
}

func parseRoutine2(str string) GoTrace {
	lines := strings.Split(str, "\n")

	line0 := strings.TrimSpace(lines[0])
	iSqO := strings.LastIndex(line0, "[")
	iSqC := strings.LastIndex(line0, "]")
	status := line0[iSqO+1 : iSqC]
	numStr := strings.TrimPrefix(line0[:iSqO], "goroutine ")
	numStr = strings.TrimSpace(numStr)
	num, _ := strconv.ParseInt(numStr, 10, 64)

	fmt.Println("line0: ", line0)

	r := GoTrace{
		Number: int(num),
		Status: status,
	}

	for i := 1; i+1 < len(lines); i += 2 {
		// pkg
		r.Stack = append(r.Stack, parseTracePoint(lines[i], lines[i+1]))
	}

	return r
}

func parseTracePoint(l1, l2 string) TraceLevel {
	l1 = strings.TrimSpace(l1)
	l2 = strings.TrimSpace(l2)

	iBrackOpen := strings.LastIndex(l1, "(")
	iFnDot := strings.LastIndex(l1[:iBrackOpen], ".")
	fn := l1[iFnDot+1:]
	pkg := l1[:iFnDot]
	// TODO find other cases (created by, ..., ...)
	pkg = strings.TrimPrefix(pkg, "created by ")

	// file
	iColon := strings.LastIndex(l2, ":")
	lineAndPtr := l2[iColon+1:]
	lineStr := strings.Split(lineAndPtr, " ")[0]
	line, _ := strconv.ParseInt(lineStr, 10, 64)
	file := l2[:iColon]

	p := TraceLevel{
		Pkg:     pkg,
		File:    file,
		Line:    int(line),
		Fn:      fn,
		Snippit: readSnippit(file, int(line)),
	}

	return p
}
