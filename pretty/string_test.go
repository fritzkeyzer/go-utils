package pretty_test

import (
	"fmt"
	"github.com/fritzkeyzer/go-utils/pretty"
)

func ExamplePrint() {
	type Object struct {
		Field             string
		privateField      string
		SomeLongFieldName string
		O                 struct {
			Field             string
			AnotherField      string
			SomeLongFieldName string
		}
	}

	o := Object{
		Field:             "hello world",
		privateField:      "world",
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

	pretty.Print(o)
	// Output:
	// pretty_test.Object: {
	//    Field:             "hello world"
	//    privateField:      "world"
	//    SomeLongFieldName: "more text"
	//    O: {
	//       Field:             "more stuff"
	//       AnotherField:      "asdasd"
	//       SomeLongFieldName: "asdasd"
	//    }
	// }
}

func ExampleString() {
	type Object struct {
		Field             string
		privateField      string
		SomeLongFieldName string
		O                 struct {
			Field             string
			AnotherField      string
			SomeLongFieldName string
		}
	}

	o := Object{
		Field:             "hello world",
		privateField:      "world",
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

	fmt.Println(pretty.String(&o))
	// Output:
	// *pretty_test.Object: -> {
	//    Field:             "hello world"
	//    privateField:      "world"
	//    SomeLongFieldName: "more text"
	//    O: {
	//       Field:             "more stuff"
	//       AnotherField:      "asdasd"
	//       SomeLongFieldName: "asdasd"
	//    }
	// }
}

func ExampleStringDepth() {
	type Object struct {
		Field string

		SomeLongFieldName string
		Complex           struct {
			Field             string
			AnotherField      string
			SomeLongFieldName string
		}
		privateField string
	}

	o := Object{
		Field:             "hello world",
		SomeLongFieldName: "more text",
		Complex: struct {
			Field             string
			AnotherField      string
			SomeLongFieldName string
		}{
			Field:             "more stuff",
			AnotherField:      "random",
			SomeLongFieldName: "text",
		},
		privateField: "secrets",
	}

	fmt.Println(pretty.StringDepth(&o, 1))
	// Output:
	// *pretty_test.Object: -> {
	//    Field:             "hello world"
	//    SomeLongFieldName: "more text"
	//    Complex:           {...}
	//    privateField:      "secrets"
	// }
}
