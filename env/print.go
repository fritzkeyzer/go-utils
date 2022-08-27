package env

import (
	"fmt"
	"reflect"
	"strings"
)

// Print dumps the contents of a struct to a formatted string.
// All fields are include except ones tagged with `secret:"true"`
// Secret fields are printed with asterisks.
func Print(ptr any) string {
	v := reflect.ValueOf(ptr)

	// Don't try to process a non-pointer value.
	if v.Kind() != reflect.Ptr || v.IsNil() {
		fmt.Printf("%s is not a pointer", v.Kind())
	}

	v = v.Elem()
	t := reflect.TypeOf(ptr).Elem()

	out := "Config:\n"

	for i := 0; i < t.NumField(); i++ {
		val := v.Field(i)

		secret, ok := t.Field(i).Tag.Lookup("secret")
		if ok && strings.ToLower(secret) == "true" {
			str := ""
			for i := range v.Field(i).String() {
				str += "*"
				if i > 10 {
					break
				}
			}
			val = reflect.ValueOf(str)

			out += fmt.Sprintf("\t%v: '%v' (SECRET)\n", t.Field(i).Name, val)
			continue
		}

		out += fmt.Sprintf("\t%v: %v\n", t.Field(i).Name, val)
	}

	return out
}
