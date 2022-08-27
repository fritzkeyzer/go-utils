package pretty_test

import (
	"fmt"
	"github.com/fritzkeyzer/go-utils/pretty"
)

func ExampleJSONString() {
	type Object struct {
		Field             string
		AnotherField      string
		SomeLongFieldName string
		O                 struct {
			Field             string
			AnotherField      string
			SomeLongFieldName string
		}
	}

	o := Object{
		Field:             "hello world",
		AnotherField:      "world",
		SomeLongFieldName: "more text",
		O: struct {
			Field             string
			AnotherField      string
			SomeLongFieldName string
		}{
			Field:             "more stuff",
			AnotherField:      "asdasd",
			SomeLongFieldName: "asdasd",
		},
	}

	fmt.Println(pretty.JSONString(o))
	// Output:
	// {
	//    "Field":             "hello world",
	//    "AnotherField":      "world",
	//    "SomeLongFieldName": "more text",
	//    "O": {
	//       "Field":             "more stuff",
	//       "AnotherField":      "asdasd",
	//       "SomeLongFieldName": "asdasd"
	//    }
	// }
}

func ExampleFormatJSONString() {
	js := `{"Field":"hello world","AnotherField":"world","SomeLongFieldName":"more text","O": {"Field":"more stuff","AnotherField":"asdasd","SomeLongFieldName":"asdasd"}}`

	fmt.Println(pretty.FormatJSONString(js))
	// Output:
	// {
	//    "Field":             "hello world",
	//    "AnotherField":      "world",
	//    "SomeLongFieldName": "more text",
	//    "O": {
	//       "Field":             "more stuff",
	//       "AnotherField":      "asdasd",
	//       "SomeLongFieldName": "asdasd"
	//    }
	// }
}
