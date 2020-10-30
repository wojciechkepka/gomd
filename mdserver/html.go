package mdserver

import (
	"bytes"
	"github.com/gobuffalo/packr"
	h "gomd/html"
	"gomd/mdserver/assets"
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

func (md *MdServer) filesBody() string {
	return RenderTemplate("./assets", "filesdiv.html", "filesList", md)
}

// Prepares full FileListView html
func (md *MdServer) filesHTML() string {
	h := h.New()
	h.AddMeta("viewport", "width=device-width, initial-scale=1.0")
	h.AddStyle(assets.FileListViewStyle(md.IsDarkMode()))
	h.AddScript(assets.ReloadJs(md.BindAddr()))
	h.AddScript(assets.JS)
	h.AddBodyItem(assets.TopBarSliderDropdown(md.IsDarkMode()))
	h.AddBodyItem(md.filesBody())
	return h.Render()
}

// Serves markdown file as html
func (md *MdServer) serveFileAsHTML(path string) string {
	if md.path == "." {
		path = path[1:]
	}
	for _, file := range md.Files {
		if file.Path == path {
			return file.AsHTML(md.IsDarkMode(), md.theme, md.BindAddr(), md.sidebarHTML())
		}
	}
	return ""
}

func (md *MdServer) sidebarHTML() string {
	links := make(map[string]string)
	for _, f := range md.Files {
		if f.IsHidden() && !md.showHidden {
			continue
		}
		links[f.Filename] = fileviewEp + f.Path
	}

	return assets.Sidebar(links)
}
