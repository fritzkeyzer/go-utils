package logpage

import (
	_ "embed"
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

const maxLines = 1000

//go:embed logpage.html
var pageRaw []byte

type line struct {
	Message string
}

type page struct {
	Lines []line
}

type LogPage struct {
	page page
}

func NewLogPage(port int) *LogPage {
	l := LogPage{
		page: page{
			Lines: make([]line, 0),
		},
	}
	l.start(port)

	return &l
}

func (p *LogPage) Write(data []byte) (n int, err error) {

	if len(p.page.Lines) > maxLines {
		p.page.Lines = append([]line{{Message: string(data)}}, p.page.Lines[:maxLines-1]...)

		return len(data), nil
	}

	p.page.Lines = append([]line{{Message: string(data)}}, p.page.Lines...)

	return len(data), nil
}

func (p *LogPage) start(port int) {
	go func() {
		http.HandleFunc("/logs", p.logsHandler)
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()
}

func (p *LogPage) logsHandler(w http.ResponseWriter, request *http.Request) {
	funcs := map[string]any{
		"contains":  strings.Contains,
		"hasPrefix": strings.HasPrefix,
		"hasSuffix": strings.HasSuffix}

	temp := template.Must(template.New("logpage.html").Funcs(funcs).Parse(string(pageRaw)))
	err := temp.Execute(w, p.page)
	if err != nil {
		panic(err)
	}
}
