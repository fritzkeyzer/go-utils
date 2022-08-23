package pretty

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

const Indent = "   "
const DefaultMaxDepth = 10

func Print(o any) {
	fmt.Println(String(o))
}

func String(o any) string {
	w := &strings.Builder{}
	Write(w, o, DefaultMaxDepth)
	lines := strings.Split(w.String(), "\n")

	alignLines(lines)

	w.Reset()
	for i := 0; i < len(lines); i++ {
		w.WriteString(strings.TrimSuffix(lines[i], " "))
		w.WriteString("\n")
	}

	return w.String()
}

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
			_, _ = fmt.Fprintf(w, "\n%s", prefix+Indent)
			write(w, prefix+Indent, v.Index(i), depth+1, maxDepth)
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
			_, _ = fmt.Fprintf(w, "\n%s", prefix+Indent)
			write(w, prefix+Indent, key, depth+1, maxDepth)
			_, _ = fmt.Fprintf(w, ": ")
			write(w, prefix+Indent, v.MapIndex(key), depth+1, maxDepth)
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
			_, _ = fmt.Fprintf(w, "\n%s", prefix+Indent)
			_, _ = fmt.Fprintf(w, "%s: ", v.Type().Field(i).Name)
			write(w, prefix+Indent, v.Field(i), depth+1, maxDepth)
			_, _ = fmt.Fprintf(w, " ")
		}
		_, _ = fmt.Fprintf(w, "\n%s}", prefix)

	default:
		_, _ = fmt.Fprintf(w, "%s (no pretty printer)", v.Type().String())
	}
}
