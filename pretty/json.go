package pretty

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Print(o any) {
	fmt.Println(JsonString(o))
}

func JsonString(o any) string {
	s, _ := json.MarshalIndent(o, "", "\t")
	str := fmt.Sprintln(string(s))

	return FormatJsonString(str)
}

func FormatJsonString(input string) string {
	str := input

	// covert single line to multiline

	// convert to spaces
	str = strings.ReplaceAll(str, "\t", "    ")

	// align values

	//

	lines := strings.Split(str, "\n")

	// depth := 0
	align := make([]int, len(lines))
	currAlign := 0
	iScope := 0
	for i := 0; i < len(lines); i++ {
		l := lines[i]

		if strings.Contains(l, "{") {
			// depth ++
			align[i] = 0
			currAlign = 0
			iScope = i + 1
			continue
		}

		if strings.Contains(l, "}") {
			//depth --
			align[i] = 0
			currAlign = 0
			iScope = i + 1
			continue
		}

		iSpace := strings.Index(l, ":")
		if iSpace > currAlign {
			for j := iScope; j <= i; j++ {
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

	str = ""
	for i := 0; i < len(lines); i++ {
		l := lines[i]
		str += l + "\n"
	}

	return str
}
