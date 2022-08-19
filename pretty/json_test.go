package pretty_test

import (
	"fmt"
	"github.com/fritzkeyzer/go-utils/pretty"
)

func ExampleJson() {
	type Object struct {
		Field string
	}

	o := Object{
		Field: "hello world",
	}

	fmt.Println(pretty.Json(o))
	//Output:"Object":{
	//	"Field": "hello world"
	//}
}

func ExampleJsonString() {
	jString := `{"Field": "hello world", "NestedField":{"Val":0}}`

	fmt.Println(pretty.JsonString(jString))
	//Output:{
	//	"Field": "hello world",
	//	"NestedField": {
	//		"Val": 0
	//	}
	//}
}
