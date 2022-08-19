package page

import (
	"fmt"
	"github.com/fritzkeyzer/go-utils/pretty"
	"html/template"
	"net/http"
	"time"
)

type Server struct {
	eventsBuf chan string
	page      page
}

func NewServer(port int, healthPage bool) *Server {
	s := Server{
		eventsBuf: make(chan string, 100),
	}

	if healthPage {
		s.startPage(port)
	}

	return &s
}

func (s *Server) Print(args ...any) {
	s.page.Events = append(s.page.Events, event{
		Time:    time.Now().Format("2006/01/02 15:04:05"),
		Message: pretty.Print(args...),
	})
}

func (s *Server) startPage(port int) {
	go func() {
		fmt.Println("Server Starting")
		http.HandleFunc("/logs", s.logsHandler)
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()
}

func (s *Server) logsHandler(w http.ResponseWriter, request *http.Request) {
	temp := template.Must(template.New("logpage.html").Parse(string(pageRaw)))
	err := temp.Execute(w, s.page)
	if err != nil {
		panic(err)
	}
}
