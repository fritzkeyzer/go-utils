package pretty

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/fritzkeyzer/go-utils/pretty/flect"
)

// Print formats any inputs in a readable format.
//
//	NOTE! unexported fields and types will not be printed
func Print(input ...any) string {
	str := ""
	for i, a := range input {
		if i > 0 {
			str += " "
		}
		str += print(a)
	}

	return str
}

func print(input any) string {
	var asd []byte
	buf := bytes.NewBuffer(asd)
	enc := toml.NewEncoder(buf)

	err := enc.Encode(input)
	if err != nil {
		return fmt.Sprint(input)
	}

	return fmt.Sprintf("%s:\n%s",
		flect.GetType(input),
		Indent(buf.String(), "  "),
	)
}
