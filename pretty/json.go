package pretty

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// JSONString Convert any object in a pretty, indented and aligned json string.
func JSONString(o any) string {
	s, _ := json.MarshalIndent(o, "", "   ")
	str := fmt.Sprintln(string(s))

	return FormatJSONString(str)
}

// FormatJSONString formats a json string into a prettier, indented and aligned string.
func FormatJSONString(input string) string {
	// covert single line to multiline
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(input), "", "\t"); err != nil {
		return input
	}
	str := prettyJSON.String()

	// convert to spaces
	str = strings.ReplaceAll(str, "\t", "   ")

	// align values
	lines := strings.Split(str, "\n")
	alignLines(lines)

	// convert lines back to string
	str = ""
	for i := 0; i < len(lines); i++ {
		l := lines[i]
		str += l + "\n"
	}

	return str
}

func alignLines(lines []string) {
	align := make([]int, len(lines))
	currAlign := 0
	iScope := 0
	for i := 0; i < len(lines); i++ {
		l := lines[i]
		if lineOpening(l) || lineClosing(l) {
			align[i] = 0
			currAlign = 0
			iScope = i + 1
			continue
		}

		iSpace := strings.Index(l, ":")
		align[i] = currAlign
		if iSpace > currAlign {
			for j := iScope; j < i; j++ {
				align[j] = iSpace
				currAlign = iSpace
			}
		}
	}

	for i := 0; i < len(lines); i++ {
		l := lines[i]

		if align[i] == 0 {
			continue
		}

		// modify l to offset and align values
		iSpace := strings.Index(l, ":")
		if iSpace > align[i] {
			align[i] = iSpace
		}

		pad := ""
		for j := 0; j < align[i]-iSpace; j++ {
			pad += " "
		}

		lines[i] = l[:iSpace+1] + pad + l[iSpace+1:]
	}
}

func lineOpening(l string) bool {
	return strings.Contains(l, "{") && !strings.Contains(l, "}")
}

func lineClosing(l string) bool {
	return strings.Contains(l, "}") && !strings.Contains(l, "{")
}
