package flect

import "reflect"

// GetType returns the type of the input, as a string
func GetType(input any) string {
	t := reflect.TypeOf(input)
	if t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	}
	return t.Name()
}
