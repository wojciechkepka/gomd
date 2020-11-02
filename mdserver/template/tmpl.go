package template

import (
	"bytes"
	"github.com/gobuffalo/packr"
	"gomd/util"
	tmpl "html/template"
)

type ITemplate interface {
	Template() (*tmpl.Template, error)
}

func TemplateFromBox(path, file, name string) (*tmpl.Template, error) {
	box := packr.NewBox(path)
	f, err := box.FindString(file)
	if err != nil {
		return nil, err
	}
	tmpl, err := tmpl.New(name).Parse(f)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func RenderTemplate(tmpl *tmpl.Template, obj interface{}) string {
	buf := []byte{}
	w := bytes.NewBuffer(buf)
	err := tmpl.Execute(w, obj)
	if err != nil {
		util.Logln(util.Error, err)
	}
	return w.String()
}

func RenderString(templ ITemplate) string {
	tmpl, err := templ.Template()
	if err != nil {
		return ""
	}
	return RenderTemplate(tmpl, templ)

}

func RenderHTML(templ ITemplate) tmpl.HTML {
	return tmpl.HTML(RenderString(templ))
}

func RenderJS(templ ITemplate) tmpl.JS {
	return tmpl.JS(RenderString(templ))
}
