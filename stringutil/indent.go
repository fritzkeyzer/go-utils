package stringutil

import (
	"fmt"
	"strings"
)

// Indent will add a specified indent to a string.
// Supports multiline strings.
func Indent(input, indent string) string {
	indReplace := "\n" + indent
	output := strings.ReplaceAll(input, "\n", indReplace)
	return fmt.Sprintf("%s%s", indent, output)
}

// IndentAndWrap will add a specified indent to a string and wrap long lines,
// to fit within a specified number of characters.
// Supports multiline strings.
func IndentAndWrap(input, indent string, wrap int, wrapChar string) string {
	res := indent
	lineLen := len(indent)
	for i := 0; i < len(input); i++ {
		c := string(input[i])

		if c == "\n" {
			res += c
			res += indent
			lineLen = len(indent)
			continue
		}

		if lineLen >= wrap {
			res += "\n"
			res += indent
			res += wrapChar
			res += c
			lineLen = len(indent + wrapChar + c)
			continue
		}

		res += c
		lineLen++
	}

	return res
}
