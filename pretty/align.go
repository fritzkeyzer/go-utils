package pretty

import "strings"

func alignFields(lines []string) {
	align := make([]int, len(lines))
	currAlign := 0
	iScope := 0
	for i := 0; i < len(lines); i++ {
		l := lines[i]
		if lineOpening(l) {
			align[i] = 0
			currAlign = 0
			iScope = i + 1
			continue
		}

		if lineClosing(l) {
			align[i] = 0
			currAlign = 0
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
	trimmed := strings.TrimSpace(l)
	return strings.HasSuffix(trimmed, "{") || strings.HasSuffix(trimmed, "[")
}

func lineClosing(l string) bool {
	trimmed := strings.TrimSpace(l)

	if trimmed == "}," {
		return true
	}

	if len(trimmed) != 1 {
		return false
	}

	return trimmed == "}" || trimmed == "]"
}
