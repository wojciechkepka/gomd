package mdserver

import (
	"gomd/mdserver/gen"
	"gomd/mdserver/highlight"
	. "gomd/mdserver/html"
	"gomd/util"
	"html/template"
)

// Serves markdown file as html
func (md *MdServer) serveFileAsHTML(path string) string {
	if md.path == "." {
		path = path[1:]
	}
	links := md.Links()
	for _, file := range md.Files {
		if file.Path == path {
			return RenderMdFile(&file, md.IsDarkMode(), md.theme, md.BindAddr(), &links)
		}
	}
	return ""
}

func (md *MdServer) FilesViewHTML() template.HTML {
	fv := FilesList{Files: &md.Files}
	return RenderHTML(&fv)
}

func (md *MdServer) TopbarHTML() template.HTML {
	themes := highlight.Themes()
	tb := Topbar{
		IsDarkMode:     md.IsDarkMode(),
		Themes:         &themes,
		DisplayButtons: false,
	}
	return RenderHTML(&tb)
}

func (md *MdServer) MainStyle() template.HTML {
	return template.HTML("<style>" + gen.FileListViewStyle(md.IsDarkMode()) + "</style>")
}

func (md *MdServer) MainScripts() template.HTML {
	return template.HTML(
		"<script>" + gen.ReloadJs(md.BindAddr()) + "</script>" +
			"<script>" + gen.JS + "</script>")
}

// Prepares full FileListView html
func (md *MdServer) MainViewHTML() string {
	tmpl, err := TemplateFromBox("../../assets/html", "main.html", "main")
	if err != nil {
		util.Logln(util.Error, err)
	}
	return RenderTemplate(tmpl, md)
}

func (md *MdServer) SidebarHTML() string {
	links := md.Links()
	sb := Sidebar{Links: &links}
	return RenderString(&sb)
}
