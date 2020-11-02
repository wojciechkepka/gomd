package template

import (
	"bytes"
	"github.com/gobuffalo/packr"
	"gomd/util"
	htmpl "html/template"
	ttmpl "text/template"
)

type HTemplate interface {
	Template() (*htmpl.Template, error)
}

type TTemplate interface {
	Template() (*ttmpl.Template, error)
}

func HTemplateFromBox(path, file, name string) (*htmpl.Template, error) {
	box := packr.NewBox(path)
	f, err := box.FindString(file)
	if err != nil {
		return nil, err
	}
	tmpl, err := htmpl.New(name).Parse(f)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func RenderHTemplate(tmpl *htmpl.Template, obj interface{}) string {
	buf := []byte{}
	w := bytes.NewBuffer(buf)
	err := tmpl.Execute(w, obj)
	if err != nil {
		util.Logln(util.Error, err)
	}
	return w.String()
}

func RenderHString(templ HTemplate) string {
	tmpl, err := templ.Template()
	if err != nil {
		return ""
	}
	return RenderHTemplate(tmpl, templ)

}

func RenderHTML(templ HTemplate) htmpl.HTML {
	return htmpl.HTML(RenderHString(templ))
}

func TTemplateFromBox(path, file, name string) (*ttmpl.Template, error) {
	box := packr.NewBox(path)
	f, err := box.FindString(file)
	if err != nil {
		return nil, err
	}
	tmpl, err := ttmpl.New(name).Parse(f)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func RenderTTemplate(tmpl *ttmpl.Template, obj interface{}) string {
	buf := []byte{}
	w := bytes.NewBuffer(buf)
	err := tmpl.Execute(w, obj)
	if err != nil {
		util.Logln(util.Error, err)
	}
	return w.String()
}

func RenderTString(templ TTemplate) string {
	tmpl, err := templ.Template()
	if err != nil {
		return ""
	}
	return RenderTTemplate(tmpl, templ)

}
