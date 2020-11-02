package js

import (
	tmpl "gomd/mdserver/template"
	"text/template"
)

type ReloadJS struct {
	BindAddr string
}

func (rjs *ReloadJS) Template() (*template.Template, error) {
	return tmpl.TTemplateFromBox("../../assets/js", "ReloadJs.js", "ReloadJS")
}

type JS struct{}

func (js *JS) Template() (*template.Template, error) {
	return tmpl.TTemplateFromBox("../../assets/js", "JS.js", "JS")
}
