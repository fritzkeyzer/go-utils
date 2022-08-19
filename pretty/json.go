package pretty

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fritzkeyzer/go-utils/pretty/flect"
)

// Json converts an object to a json string.
//
//	NOTE! types and fields must be exported.
func Json(input any) string {
	inpucEnc, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		return fmt.Sprintf("%+v", input)
	}

	return fmt.Sprintf("\"%s\":%s",
		flect.GetType(input),
		string(inpucEnc),
	)
}

// JsonString converts a single line json string to a
// formatted and indented multiline json string.
func JsonString(jsonS string) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(jsonS), "", "\t"); err != nil {
		return jsonS
	}
	return prettyJSON.String()
}
