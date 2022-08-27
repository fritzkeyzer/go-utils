package pretty

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

type Options struct {
	// 	default = false
	// Write everything to a single line.
	// Overwrites Indent and AlignFields options.
	SingleLine bool

	// 	default = "   "
	Indent string

	// 	default = true
	AlignFields bool

	// 	default = 5
	// MaxDepth sets the limit on nested struct traversal.
	// Clipped fields are printed as ...
	MaxDepth int

	// 	default = false
	//
	//  // eg: PrintPointers = true
	//  *pretty_test.Object: 0x14000092120 -> {
	//  	...
	//  }
	PrintPointers bool
}

// DefaultOptions is the set of options used for pretty.Print() and pretty.String()
// if no Options are passed.
var DefaultOptions = Options{
	SingleLine:    false,
	Indent:        "   ",
	AlignFields:   true,
	MaxDepth:      5,
	PrintPointers: false,
}

// Print pretty prints any object to StdErr.
// Options are optional, if none are provided, DefaultOptions is used.
// Only the first Options argument is used.
//
//	// Print(o) is equivalent to:
//	fmt.Println(String(o))
//
//	// Print(o, options) is equivalent to:
//	fmt.Println(String(o, options))
func Print(o any, options ...Options) {
	fmt.Println(String(o, options...))
}

// String returns a pretty formatted string of any input.
// Options are optional, if none are provided, DefaultOptions is used.
// Only the first Options argument is used.
func String(o any, options ...Options) string {
	opt := DefaultOptions
	if len(options) > 0 {
		opt = options[0]
	}

	nlc := "\n"
	if opt.SingleLine {
		nlc = " "
	}

	w := &strings.Builder{}
	writeAsString(w, o, &opt, nlc)
	lines := strings.Split(w.String(), nlc)

	if opt.AlignFields && !opt.SingleLine {
		alignFields(lines)
	}

	w.Reset()
	for i := 0; i < len(lines); i++ {
		w.WriteString(strings.TrimSuffix(lines[i], " "))
		w.WriteString(nlc)
	}

	return w.String()
}

// writeAsString a pretty formatted string of any input to a provided io.Writer.
func writeAsString(w io.Writer, value any, options *Options, nlc string) {
	_, _ = fmt.Fprintf(w, "%s: ", reflect.ValueOf(value).Type())

	if options.SingleLine {
		options.Indent = ""
	}

	node(w, "", reflect.ValueOf(value), 0, options, nlc)
}

func node(w io.Writer, prefix string, v reflect.Value, depth int, options *Options, nlc string) {
	if v.CanInterface() {
		value := v.Interface()
		if stringer, ok := value.(fmt.Stringer); ok {
			if !(v.Kind() == reflect.Ptr && v.IsNil()) {
				str := strings.ReplaceAll(stringer.String(), "\n", nlc)
				_, _ = fmt.Fprint(w, str)
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
		if options.MaxDepth > 0 && depth > options.MaxDepth {
			_, _ = fmt.Fprintf(w, "...")
			return
		}
		if v.IsNil() {
			_, _ = fmt.Fprint(w, "nil")
			return
		}
		if options.PrintPointers {
			_, _ = fmt.Fprintf(w, "%p ", v.Interface())
		}

		_, _ = fmt.Fprint(w, "-> ")
		node(w, prefix, v.Elem(), depth+1, options, nlc)
	case reflect.Interface:
		if options.MaxDepth > 0 && depth > options.MaxDepth {
			_, _ = fmt.Fprintf(w, "...")
			return
		}
		if v.IsNil() {
			_, _ = fmt.Fprint(w, "nil")
			return
		}
		node(w, prefix, v.Elem(), depth+1, options, nlc)
	case reflect.Slice:
		if v.IsNil() {
			_, _ = fmt.Fprint(w, "nil")
			return
		}
		fallthrough
	case reflect.Array:
		if options.MaxDepth > 0 && depth > options.MaxDepth {
			_, _ = fmt.Fprintf(w, "[...]")
			return
		}
		if v.Type().Elem().Kind() == reflect.Uint8 {
			_, _ = fmt.Fprintf(w, "%d bytes", v.Len())
			return
		}
		_, _ = fmt.Fprintf(w, "[")
		for i := 0; i < v.Len(); i++ {
			_, _ = fmt.Fprintf(w, "%s%s", nlc, prefix+options.Indent)
			node(w, prefix+options.Indent, v.Index(i), depth+1, options, nlc)
			_, _ = fmt.Fprintf(w, ", ")
		}
		_, _ = fmt.Fprintf(w, "%s%s]", nlc, prefix)
	case reflect.Map:
		if options.MaxDepth > 0 && depth > options.MaxDepth {
			_, _ = fmt.Fprintf(w, "{...}")
			return
		}
		_, _ = fmt.Fprintf(w, "{")
		for _, key := range v.MapKeys() {
			_, _ = fmt.Fprintf(w, "%s%s", nlc, prefix+options.Indent)
			node(w, prefix+options.Indent, key, depth+1, options, nlc)
			_, _ = fmt.Fprintf(w, ": ")
			node(w, prefix+options.Indent, v.MapIndex(key), depth+1, options, nlc)
			_, _ = fmt.Fprintf(w, ", ")
		}
		_, _ = fmt.Fprintf(w, "%s%s}", nlc, prefix)
	case reflect.Struct:
		if options.MaxDepth > 0 && depth > options.MaxDepth {
			_, _ = fmt.Fprintf(w, "{...}")
			return
		}
		_, _ = fmt.Fprintf(w, "{")
		for i := 0; i < v.NumField(); i++ {
			_, _ = fmt.Fprintf(w, "%s%s", nlc, prefix+options.Indent)
			_, _ = fmt.Fprintf(w, "%s: ", v.Type().Field(i).Name)
			node(w, prefix+options.Indent, v.Field(i), depth+1, options, nlc)
			_, _ = fmt.Fprintf(w, " ")
		}
		_, _ = fmt.Fprintf(w, "%s%s}", nlc, prefix)
	default:
		_, _ = fmt.Fprintf(w, "%s (no pretty printer)", v.Type().String())
	}
}
