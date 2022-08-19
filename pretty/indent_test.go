package pretty_test

import (
	"fmt"
	"github.com/fritzkeyzer/go-utils/pretty"
	"testing"
)

func TestIndent(t *testing.T) {
	type args struct {
		input  string
		indent string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "single line text",
			args: args{
				input:  "single line text",
				indent: "\t",
			},
			want: "\tsingle line text",
		},
		{
			name: "2 line text",
			args: args{
				input:  "first line\n2nd line",
				indent: "\t",
			},
			want: "\tfirst line\n\t2nd line",
		},
		{
			name: "2 line already indented text",
			args: args{
				input:  "first line\n\t2nd line",
				indent: "\t",
			},
			want: "\tfirst line\n\t\t2nd line",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pretty.Indent(tt.args.input, tt.args.indent); got != tt.want {
				t.Errorf("Indent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndentAndWrap(t *testing.T) {
	type args struct {
		input    string
		indent   string
		wrap     int
		wrapChar string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "single line text",
			args: args{
				input:    "single line text",
				indent:   "  ",
				wrap:     20,
				wrapChar: "_︎",
			},
			want: "  single line text",
		},
		{
			name: "wrapping",
			args: args{
				input:    "0123456789",
				indent:   "  ",
				wrap:     10,
				wrapChar: "_",
			},
			want: "  01234567\n  _89",
		},
		{
			name: "wrapping multiline",
			args: args{
				input:    "0123456789\n0123456789",
				indent:   "  ",
				wrap:     10,
				wrapChar: "_",
			},
			want: "  01234567\n  _89\n  01234567\n  _89",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(pretty.IndentAndWrap(tt.args.input, tt.args.indent, tt.args.wrap, tt.args.wrapChar))
			if got := pretty.IndentAndWrap(tt.args.input, tt.args.indent, tt.args.wrap, tt.args.wrapChar); got != tt.want {
				t.Errorf("got != want:\ngot: %v\nwant: %v", got, tt.want)
			}
		})
	}
}
