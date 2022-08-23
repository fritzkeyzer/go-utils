package pretty

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

const indent = "   "
const defaultMaxDepth = 10

// Print pretty prints any object to StdErr
func Print(o any) {
	fmt.Println(String(o))
}

// String returns a pretty formatted string of any input
func String(o any) string {
	w := &strings.Builder{}
	Write(w, o, defaultMaxDepth)
	lines := strings.Split(w.String(), "\n")

	alignLines(lines)

	w.Reset()
	for i := 0; i < len(lines); i++ {
		w.WriteString(strings.TrimSuffix(lines[i], " "))
		w.WriteString("\n")
	}

	return w.String()
}

// StringDepth returns a pretty formatted string of any input.
// Recursion into fields is limited by maxDepth.
func StringDepth(o any, maxDepth int) string {
	w := &strings.Builder{}
	Write(w, o, maxDepth)
	lines := strings.Split(w.String(), "\n")

	alignLines(lines)

	w.Reset()
	for i := 0; i < len(lines); i++ {
		w.WriteString(strings.TrimSuffix(lines[i], " "))
		w.WriteString("\n")
	}

	return w.String()
}

// Write a pretty formatted string of any input to a provided io.Writer.
// Recursion into fields is limited by maxDepth.
func Write(w io.Writer, value any, maxDepth int) {
	_, _ = fmt.Fprintf(w, "%s: ", reflect.ValueOf(value).Type())

	write(w, "", reflect.ValueOf(value), 0, maxDepth)
}

func write(w io.Writer, prefix string, v reflect.Value, depth, maxDepth int) {
	if v.CanInterface() {
		value := v.Interface()
		if stringer, ok := value.(fmt.Stringer); ok {
			if !(v.Kind() == reflect.Ptr && v.IsNil()) {
				_, _ = fmt.Fprint(w, stringer.String())
				return
			}
		}
	}

	switch v.Kind() {
	case reflect.Bool:
		_, _ = fmt.Fprintf(w, "%t", v.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		_, _ = fmt.Fprintf(w, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		_, _ = fmt.Fprintf(w, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		_, _ = fmt.Fprintf(w, "%f", v.Float())

	case reflect.String:
		_, _ = fmt.Fprintf(w, "\"%s\"", v.String())

	case reflect.Ptr:
		if maxDepth > 0 && depth > maxDepth {
			_, _ = fmt.Fprintf(w, "...")
			return
		}
		if v.IsNil() {
			_, _ = fmt.Fprint(w, "nil")
			return
		}
		// _, _ = fmt.Fprintf(w, "%p -> ", v.Interface())
		_, _ = fmt.Fprint(w, "-> ")
		write(w, prefix, v.Elem(), depth+1, maxDepth)

	case reflect.Interface:
		if maxDepth > 0 && depth > maxDepth {
			_, _ = fmt.Fprintf(w, "...")
			return
		}
		if v.IsNil() {
			_, _ = fmt.Fprint(w, "nil")
			return
		}
		write(w, prefix, v.Elem(), depth+1, maxDepth)

	case reflect.Slice:
		if v.IsNil() {
			_, _ = fmt.Fprint(w, "nil")
			return
		}
		fallthrough
	case reflect.Array:
		if maxDepth > 0 && depth > maxDepth {
			_, _ = fmt.Fprintf(w, "[...]")
			return
		}
		if v.Type().Elem().Kind() == reflect.Uint8 {
			_, _ = fmt.Fprintf(w, "%d bytes", v.Len())
			return
		}
		_, _ = fmt.Fprintf(w, "[")
		for i := 0; i < v.Len(); i++ {
			_, _ = fmt.Fprintf(w, "\n%s", prefix+indent)
			write(w, prefix+indent, v.Index(i), depth+1, maxDepth)
			_, _ = fmt.Fprintf(w, ", ")
		}
		_, _ = fmt.Fprintf(w, "\n%s]", prefix)

	case reflect.Map:
		if maxDepth > 0 && depth > maxDepth {
			_, _ = fmt.Fprintf(w, "{...}")
			return
		}
		_, _ = fmt.Fprintf(w, "{")
		for _, key := range v.MapKeys() {
			_, _ = fmt.Fprintf(w, "\n%s", prefix+indent)
			write(w, prefix+indent, key, depth+1, maxDepth)
			_, _ = fmt.Fprintf(w, ": ")
			write(w, prefix+indent, v.MapIndex(key), depth+1, maxDepth)
			_, _ = fmt.Fprintf(w, ", ")
		}
		_, _ = fmt.Fprintf(w, "\n%s}", prefix)

	case reflect.Struct:
		if maxDepth > 0 && depth > maxDepth {
			_, _ = fmt.Fprintf(w, "{...}")
			return
		}
		_, _ = fmt.Fprintf(w, "{")
		for i := 0; i < v.NumField(); i++ {
			_, _ = fmt.Fprintf(w, "\n%s", prefix+indent)
			_, _ = fmt.Fprintf(w, "%s: ", v.Type().Field(i).Name)
			write(w, prefix+indent, v.Field(i), depth+1, maxDepth)
			_, _ = fmt.Fprintf(w, " ")
		}
		_, _ = fmt.Fprintf(w, "\n%s}", prefix)

	default:
		_, _ = fmt.Fprintf(w, "%s (no pretty printer)", v.Type().String())
	}
}
