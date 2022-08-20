package envutil

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Setter is called for any complex struct field with an
// implementation, allowing developers to override Set
// behaviour.
type Setter interface {
	Set(string) error
}

func LoadCfgFromEnv(ptr any) error {
	v := reflect.ValueOf(ptr)

	// Don't try to process a non-pointer value.
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("%s is not a pointer", v.Kind())
	}

	v = v.Elem()
	t := reflect.TypeOf(ptr).Elem()

	for i := 0; i < t.NumField(); i++ {
		if err := processField(t.Field(i), v.Field(i)); err != nil {
			return err
		}
	}

	return nil
}

func processField(t reflect.StructField, v reflect.Value) error {
	envTag, ok := t.Tag.Lookup("env")
	if !ok {
		return fmt.Errorf("field '%s' is missing 'env' tag", t.Name)
	}

	if !v.CanSet() {
		return fmt.Errorf("field '%s' cannot be set", t.Name)
	}

	// get from environment
	if env, ok := os.LookupEnv(envTag); ok {
		err := setField(t, v, env)
		if err != nil {
			return fmt.Errorf("setting field '%s' from env: %w", t.Name, err)
		}
		return nil
	}

	// get from default
	if defaultValue, ok := t.Tag.Lookup("default"); ok {
		err := setField(t, v, defaultValue)
		if err != nil {
			return fmt.Errorf("setting field '%s' from default: %w", t.Name, err)
		}
		return nil
	}

	return fmt.Errorf("environment var: %s could not be found and no default specified", envTag)
}

func setField(t reflect.StructField, v reflect.Value, value string) error {
	// If field implements the Setter interface, invoke it now and
	// don't continue attempting to set the primitive values.
	if _, ok := v.Interface().(Setter); ok {
		instance := reflect.New(t.Type.Elem())
		v.Set(instance)

		// Re-assert the type with the newed-up instance and call.
		setter := v.Interface().(Setter)
		if err := setter.Set(value); err != nil {
			return fmt.Errorf("custom setter: %w", err)
		}
		return nil
	}

	switch v.Kind() {
	case reflect.Slice:
		return setSlice(t, v, value)
	case reflect.Bool:
		return setBool(v, value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return setInt(v, value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return setUint(v, value)
	case reflect.Float32, reflect.Float64:
		return setFloat(v, value)
	case reflect.String:
		return setString(v, value)
	default:
		return fmt.Errorf("%s is not supported", v.Kind())
	}
}

func setBool(fieldValue reflect.Value, value string) error {
	b, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}

	fieldValue.SetBool(b)
	return nil
}

func setInt(fieldValue reflect.Value, value string) error {
	if fieldValue.Type() == reflect.TypeOf((*time.Duration)(nil)).Elem() {
		return setDuration(fieldValue, value)
	}

	i, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return err
	}

	fieldValue.SetInt(i)
	return nil
}

func setUint(fieldValue reflect.Value, value string) error {
	i, err := strconv.ParseUint(value, 0, 64)
	if err != nil {
		return err
	}

	fieldValue.SetUint(i)
	return nil
}

func setFloat(fieldValue reflect.Value, value string) error {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}

	fieldValue.SetFloat(f)
	return nil
}

func setDuration(fieldValue reflect.Value, value string) error {
	d, err := time.ParseDuration(value)
	if err != nil {
		return err
	}

	fieldValue.SetInt(d.Nanoseconds())
	return nil
}

func setString(fieldValue reflect.Value, value string) error {
	fieldValue.SetString(value)
	return nil
}

func setSlice(t reflect.StructField, v reflect.Value, value string) error {
	// []uint8 and []byte are special cases, as they can be used to store
	// binary data, which we'll favour over storing comma-separated uint8's.
	binaryType := reflect.TypeOf([]uint8{})
	if t.Type == binaryType {
		v.SetBytes([]byte(value))
		return nil
	}

	delimiter := getDelimiter(t)
	rawValues := split(value, delimiter)
	if len(rawValues) == 0 {
		return nil
	}

	var slice reflect.Value
	n := len(rawValues)
	switch v.Type() {
	case reflect.TypeOf([]string{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]string{}), n, n)
	case reflect.TypeOf([]bool{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]bool{}), n, n)
	case reflect.TypeOf([]int{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]int{}), n, n)
	case reflect.TypeOf([]int8{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]int8{}), n, n)
	case reflect.TypeOf([]int16{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]int16{}), n, n)
	case reflect.TypeOf([]int32{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]int32{}), n, n)
	case reflect.TypeOf([]int64{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]int64{}), n, n)
	case reflect.TypeOf([]uint{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]uint{}), n, n)
	case reflect.TypeOf([]uint16{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]uint16{}), n, n)
	case reflect.TypeOf([]uint32{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]uint32{}), n, n)
	case reflect.TypeOf([]uint64{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]uint64{}), n, n)
	case reflect.TypeOf([]float32{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]float32{}), n, n)
	case reflect.TypeOf([]float64{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]float64{}), n, n)
	case reflect.TypeOf([]time.Duration{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]time.Duration{}), n, n)
	default:
		return fmt.Errorf("%v is not supported", v.Type())
	}

	for i, value := range rawValues {
		var err error
		switch slice.Index(i).Kind() {
		case reflect.Bool:
			err = setBool(slice.Index(i), value)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			err = setInt(slice.Index(i), value)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			err = setUint(slice.Index(i), value)
		case reflect.Float32, reflect.Float64:
			err = setFloat(slice.Index(i), value)
		case reflect.String:
			err = setString(slice.Index(i), value)
		default:
			err = fmt.Errorf("%s is not supported", slice.Kind())
		}
		if err != nil {
			return fmt.Errorf("populating slice[%d]: %w", i, err)
		}
	}

	v.Set(slice)

	return nil
}

func split(value string, delimiter string) []string {
	var out []string

	raw := strings.Split(value, delimiter)
	for _, r := range raw {
		sanitised := strings.Trim(r, " ")
		if len(sanitised) > 0 {
			out = append(out, sanitised)
		}
	}

	return out
}

func getDelimiter(t reflect.StructField) string {
	if d, ok := t.Tag.Lookup("delimiter"); ok {
		return d
	}
	return ","
}
