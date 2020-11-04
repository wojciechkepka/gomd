package js

import (
	tmpl "gomd/mdserver/template"
	"text/template"
)

//ReloadJS is a script responsible for starting a websocket
//listening for reload messages and triggering tab refresh
//in the browser.
type ReloadJS struct {
	BindAddr string
}

func (rjs *ReloadJS) Template() (*template.Template, error) {
	return tmpl.TTemplateFromPath("/assets/js/ReloadJs.js", "ReloadJS")
}

//JS is all common javascript functionality groupped together
type JS struct{}

func (js *JS) Template() (*template.Template, error) {
	return tmpl.TTemplateFromPath("/assets/js/JS.js", "JS")
}
