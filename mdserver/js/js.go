package js

import (
	tmpl "gomd/mdserver/template"
	"text/template"
)

//JS is all common javascript functionality groupped together
type JS struct {
	//BindAddr is binding address of websocket responsible for reloading browser tab
	BindAddr string
}

func (js *JS) Template() (*template.Template, error) {
	return tmpl.TTemplateFromPath("/assets/js/JS.js", "JS")
}
