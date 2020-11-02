package template

import (
	"bytes"
	"github.com/gobuffalo/packr"
	"gomd/util"
	htmpl "html/template"
	ttmpl "text/template"
)

// HTemplate is an interface grouping all HTML templates together
// providing functionality like RenderHString
type HTemplate interface {
	Template() (*htmpl.Template, error)
}

// TTemplate is an interface grouping all text templates together
// providing functionality like RenderHString
type TTemplate interface {
	Template() (*ttmpl.Template, error)
}

// HTemplateFromBox returns a HTML template with name set to name,
// found in a file within a box specified by path
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

// RenderHTemplate renders html template to string logging errors and
// returning empty string by default
func RenderHTemplate(tmpl *htmpl.Template, obj interface{}) string {
	buf := []byte{}
	w := bytes.NewBuffer(buf)
	err := tmpl.Execute(w, obj)
	if err != nil {
		util.Logln(util.Error, err)
	}
	return w.String()
}

// RenderHString grabs an HTML template specified by templ interface
// and renders it using that object.
func RenderHString(templ HTemplate) string {
	tmpl, err := templ.Template()
	if err != nil {
		return ""
	}
	return RenderHTemplate(tmpl, templ)

}

// RenderHTML returns a rendered template as template.HTML
func RenderHTML(templ HTemplate) htmpl.HTML {
	return htmpl.HTML(RenderHString(templ))
}

// TTemplateFromBox returns a text template with name set to name,
// found in a file within a box specified by path
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

// RenderTTemplate renders text template to string logging errors and
// returning empty string by default
func RenderTTemplate(tmpl *ttmpl.Template, obj interface{}) string {
	buf := []byte{}
	w := bytes.NewBuffer(buf)
	err := tmpl.Execute(w, obj)
	if err != nil {
		util.Logln(util.Error, err)
	}
	return w.String()
}

// RenderTString grabs a text template specified by templ interface
// and renders it using that object.
func RenderTString(templ TTemplate) string {
	tmpl, err := templ.Template()
	if err != nil {
		return ""
	}
	return RenderTTemplate(tmpl, templ)

}
