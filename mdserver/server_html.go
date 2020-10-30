package mdserver

import (
	"gomd/mdserver/assets"
	. "gomd/mdserver/html"
	"gomd/util"
	"html/template"
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

func (md *MdServer) FilesViewHTML() template.HTML {
	fv := FilesView{Files: &md.Files}
	return RenderHTML(&fv)
}

func (md *MdServer) TopbarHTML() template.HTML {
	themes := assets.Themes()
	tb := Topbar{
		IsDarkMode:     md.IsDarkMode(),
		Themes:         &themes,
		DisplayButtons: false,
	}
	return RenderHTML(&tb)
}

func (md *MdServer) MainStyle() template.HTML {
	return template.HTML("<style>" + assets.FileListViewStyle(md.IsDarkMode()) + "</style>")
}

func (md *MdServer) MainScripts() template.HTML {
	return template.HTML(
		"<script>" + assets.ReloadJs(md.BindAddr()) + "</script>" +
			"<script>" + assets.JS + "</script>")
}

// Prepares full FileListView html
func (md *MdServer) MainViewHTML() string {
	tmpl, err := TemplateFromBox("./assets", "main.html", "main")
	if err != nil {
		util.Logln(util.Error, err)
	}
	return RenderTemplate(tmpl, md)
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
