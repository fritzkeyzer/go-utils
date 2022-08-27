package env

import (
	"fmt"
	"reflect"
	"strings"
)

/*
ReplaceVars will replace variables in an input string.

eg:

	type Config struct {
		Var string `env:"VAR_NAME"`
	}
	cfg := Config{ Var: "xxx" }
	ReplaceVars("var = {VAR_NAME}", &cfg) // result: "var = xxx"
*/
func ReplaceVars(input string, ptr any) string {
	v := reflect.ValueOf(ptr)

	// Don't try to process a non-pointer value.
	if v.Kind() != reflect.Ptr || v.IsNil() {
		fmt.Printf("%s is not a pointer", v.Kind())
	}

	v = v.Elem()
	t := reflect.TypeOf(ptr).Elem()

	output := input

	for i := 0; i < t.NumField(); i++ {
		envName, ok := t.Field(i).Tag.Lookup("env")
		if !ok {
			continue
		}

		envName = "{" + envName + "}"

		output = strings.ReplaceAll(output, envName, v.Field(i).String())
	}

	return output
}
