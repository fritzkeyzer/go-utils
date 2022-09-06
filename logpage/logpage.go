package logpage

import (
	_ "embed"
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

// maxLines is the length of the lines buffer.
// This is the maximum number of lines that will render on a single html page.
const maxLines = 1000

//go:embed logpage.html
var pageRaw []byte

// funcs contains a map of functions that can be used within templates
var funcs = map[string]any{
	"contains":  strings.Contains,
	"hasPrefix": strings.HasPrefix,
	"hasSuffix": strings.HasSuffix,
}

type line struct {
	Message string
}

type page struct {
	Lines []line
}

// LogPage satisfies the io.Writer and http.Handler interfaces.
type LogPage struct {
	page page
}

// New creates a LogPage and initialises the write buffer.
func New() *LogPage {
	l := LogPage{
		page: page{
			Lines: make([]line, 0),
		},
	}

	return &l
}

// Host logpage on a specified port.
func (p *LogPage) Host(port int, path string) error {
	http.HandleFunc("/"+path, p.ServeHTTP)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

// Write to LogPage logs. Each call is a separate log line.
func (p *LogPage) Write(data []byte) (n int, err error) {
	if len(p.page.Lines) > maxLines {
		p.page.Lines = append([]line{{Message: string(data)}}, p.page.Lines[:maxLines-1]...)

		return len(data), nil
	}

	p.page.Lines = append([]line{{Message: string(data)}}, p.page.Lines...)

	return len(data), nil
}

// ServeHTTP satisfies the http.Handler interface.
// Populates an html template with log lines.
func (p *LogPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.New("logpage.html").Funcs(funcs).Parse(string(pageRaw)))
	err := temp.Execute(w, p.page)
	if err != nil {
		panic(err)
	}
}
