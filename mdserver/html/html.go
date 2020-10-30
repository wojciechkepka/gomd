package html

import (
	"bytes"
	"github.com/gobuffalo/packr"
	. "gomd/mdserver/mdfile"
	"gomd/util"
	"html/template"
)

func TemplateFromBox(path, file, name string) (*template.Template, error) {
	box := packr.NewBox(path)
	f, err := box.FindString(file)
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New(name).Parse(f)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func RenderTemplate(tmpl *template.Template, obj interface{}) string {
	buf := []byte{}
	w := bytes.NewBuffer(buf)
	err := tmpl.Execute(w, obj)
	if err != nil {
		util.Logln(util.Error, err)
	}
	return w.String()
}

type ITemplate interface {
	Template() (*template.Template, error)
}

func RenderString(templ ITemplate) string {
	tmpl, err := templ.Template()
	if err != nil {
		return ""
	}
	return RenderTemplate(tmpl, templ)

}

func RenderHTML(tmpl ITemplate) template.HTML {
	return template.HTML(RenderString(tmpl))
}

type Sidebar struct {
	Links *map[string]string
}

func (sb *Sidebar) Template() (*template.Template, error) {
	return TemplateFromBox("./assets", "sidebar.html", "sidebar")
}

type FilesView struct {
	Files *[]MdFile
}

func (fv *FilesView) Template() (*template.Template, error) {
	return TemplateFromBox("./assets", "filesdiv.html", "fileview")
}

type ThemeDropdown struct {
	Themes *[]string
}

func (td *ThemeDropdown) Template() (*template.Template, error) {
	return TemplateFromBox("./assets", "theme_dropdown.html", "theme_dropdown")
}

type Topbar struct {
	DisplayButtons, IsDarkMode bool
	Themes                     *[]string
}

func (tb *Topbar) ThemeDropdown() template.HTML {
	td := ThemeDropdown{Themes: tb.Themes}
	return RenderHTML(&td)
}

func (tb *Topbar) Template() (*template.Template, error) {
	return TemplateFromBox("./assets", "top_bar.html", "topbar")
}
