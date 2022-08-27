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
	alignFields(lines)

	// convert lines back to string
	str = ""
	for i := 0; i < len(lines); i++ {
		l := lines[i]
		str += l + "\n"
	}

	return str
}
