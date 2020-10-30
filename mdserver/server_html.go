package mdserver

import (
	h "gomd/html"
	"gomd/mdserver/assets"
	. "gomd/mdserver/html"
)

// Serves markdown file as html
func (md *MdServer) serveFileAsHTML(path string) string {
	if md.path == "." {
		path = path[1:]
	}
	for _, file := range md.Files {
		if file.Path == path {
			return file.AsHTML(md.IsDarkMode(), md.theme, md.BindAddr(), md.SidebarHTML())
		}
	}
	return ""
}

func (md *MdServer) FilesViewHTML() string {
	fv := FilesView{Files: &md.Files}
	return RenderString(&fv)
}

// Prepares full FileListView html
func (md *MdServer) MainViewHTML() string {
	h := h.New()
	h.AddMeta("viewport", "width=device-width, initial-scale=1.0")
	h.AddStyle(assets.FileListViewStyle(md.IsDarkMode()))
	h.AddScript(assets.ReloadJs(md.BindAddr()))
	h.AddScript(assets.JS)
	themes := assets.Themes()
	tb := Topbar{
		IsDarkMode:     md.IsDarkMode(),
		Themes:         &themes,
		DisplayButtons: false,
	}
	h.AddBodyItem(RenderString(&tb))
	h.AddBodyItem(md.FilesViewHTML())
	return h.Render()
}

func (md *MdServer) SidebarHTML() string {
	links := make(map[string]string)
	for _, f := range md.Files {
		if f.IsHidden() && !md.showHidden {
			continue
		}
		links[f.Filename] = fileviewEp + f.Path
	}
	sb := Sidebar{Links: &links}
	return RenderString(&sb)
}
