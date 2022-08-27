package pretty_test

import (
	"fmt"
	"github.com/fritzkeyzer/go-utils/pretty"
)

// See example:
func ExamplePrint() {
	type Object struct {
		Field             string
		privateField      string
		SomeLongFieldName string
		NestedObject      struct {
			Field             string
			AnotherField      string
			SomeLongFieldName string
		}
	}

	o := Object{
		Field:             "hello world",
		privateField:      "world",
		SomeLongFieldName: "more text",
		NestedObject: struct {
			Field             string
			AnotherField      string
			SomeLongFieldName string
		}{
			Field:             "more stuff",
			AnotherField:      "asdasd",
			SomeLongFieldName: "asdasd",
		},
	}

	pretty.Print(&o)
	// Output:
	// *pretty_test.Object: -> {
	//    Field:             "hello world"
	//    privateField:      "world"
	//    SomeLongFieldName: "more text"
	//    NestedObject: {
	//       Field:             "more stuff"
	//       AnotherField:      "asdasd"
	//       SomeLongFieldName: "asdasd"
	//    }
	// }
}

func ExampleString_withOptions() {
	type Object struct {
		Field             string
		privateField      string
		SomeLongFieldName string
		NestedObject      struct {
			Field             string
			AnotherField      string
			SomeLongFieldName string
		}
	}

	o := Object{
		Field:             "hello world",
		privateField:      "world",
		SomeLongFieldName: "more text",
		NestedObject: struct {
			Field             string
			AnotherField      string
			SomeLongFieldName string
		}{
			Field:             "more stuff",
			AnotherField:      "asdasd",
			SomeLongFieldName: "asdasd",
		},
	}

	// options:
	// type Options struct {
	//    SingleLine    bool
	//    Indent        string
	//    AlignFields   bool
	//    MaxDepth      int
	//    PrintPointers bool
	//}

	options := pretty.DefaultOptions
	options.MaxDepth = 1 // max depth clips nested structs and slices like: {...}

	fmt.Println(pretty.String(&o, options))
	// Output:
	// *pretty_test.Object: -> {
	//    Field:             "hello world"
	//    privateField:      "world"
	//    SomeLongFieldName: "more text"
	//    NestedObject:      {...}
	// }
}

func ExampleString_singleLine() {
	type Object struct {
		Field             string
		privateField      string
		SomeLongFieldName string
		NestedObject      struct {
			Field             string
			AnotherField      string
			SomeLongFieldName string
		}
	}

	o := Object{
		Field:             "hello world",
		privateField:      "world",
		SomeLongFieldName: "more text",
		NestedObject: struct {
			Field             string
			AnotherField      string
			SomeLongFieldName string
		}{
			Field:             "more stuff",
			AnotherField:      "asdasd",
			SomeLongFieldName: "asdasd",
		},
	}

	// with options:
	options := pretty.DefaultOptions
	options.MaxDepth = 1 // max depth clips nested structs and slices like: {...}
	options.SingleLine = true
	fmt.Println(pretty.String(&o, options))
	// Output:
	// *pretty_test.Object: -> { Field: "hello world"  privateField: "world"  SomeLongFieldName: "more text"  NestedObject: {...}  }
}
