package page

import _ "embed"

type event struct {
	Time    string
	Message string
}

type page struct {
	Events []event
}

//go:embed page.html
var pageRaw []byte
