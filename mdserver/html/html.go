package html

import (
	"bytes"
	"github.com/gobuffalo/packr"
	. "gomd/mdserver/mdfile"
	"gomd/util"
	"html/template"
)

func templateFromBox(path, file, name string) (*template.Template, error) {
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

func RenderTemplate(path, file, name string, obj interface{}) string {
	tmpl, err := templateFromBox(path, file, name)
	if err != nil {
		util.Logln(util.Error, err)
	}
	buf := []byte{}
	w := bytes.NewBuffer(buf)
	err = tmpl.Execute(w, obj)
	if err != nil {
		util.Logln(util.Error, err)
	}
	return w.String()
}

type Sidebar struct {
	Links *map[string]string
}

func (sb *Sidebar) Render() string {
	return RenderTemplate("./assets", "sidebar.html", "sidebar", sb)
}

type FilesView struct {
	Files *[]MdFile
}

func (fv *FilesView) Render() string {
	return RenderTemplate("./assets", "filesdiv.html", "fileview", fv)
}
