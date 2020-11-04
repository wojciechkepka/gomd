package mdserver

import (
	"gomd/mdserver/css"
	"gomd/mdserver/highlight"
	. "gomd/mdserver/html"
	"gomd/mdserver/js"
	tmpl "gomd/mdserver/template"
	"gomd/util"
	"html/template"
)

// Serves markdown file as html
func (md *MdServer) serveFileAsHTML(path string) string {
	if md.path == "." {
		path = path[1:]
	}
	links := &md.Files.Links
	for _, file := range md.Files.Files {
		if file.Path == path {
			return RenderMdFile(&file, md.IsDarkMode(), md.BindAddr(), md.theme, links)
		}
	}
	return ""
}

func (md *MdServer) FilesViewHTML() template.HTML {
	fv := FilesList{Files: &md.Files.Files, ShowHidden: md.showHidden}
	return tmpl.RenderHTML(&fv)
}

func (md *MdServer) TopbarHTML() template.HTML {
	themes := highlight.Themes()
	tb := Topbar{
		IsDarkMode:     md.IsDarkMode(),
		Themes:         &themes,
		DisplayButtons: false,
	}
	return tmpl.RenderHTML(&tb)
}

func (md *MdServer) MainStyle() template.HTML {
	return template.HTML("<style>" + css.FileListViewStyle(md.IsDarkMode()) + "</style>")
}

func (md *MdServer) MainScripts() template.HTML {
	reload := js.ReloadJS{BindAddr: md.BindAddr()}
	js := js.JS{}
	return template.HTML("<script>" + tmpl.RenderTString(&reload) + "</script>" +
		"<script>" + tmpl.RenderTString(&js) + "</script>")
}

// Prepares full FileListView html
func (md *MdServer) MainViewHTML() string {
	templ, err := tmpl.HTemplateFromPath("/assets/html/main.html", "main")
	if err != nil {
		util.Logln(util.Error, err)
	}
	return tmpl.RenderHTemplate(templ, md)
}

func (md *MdServer) SidebarHTML() string {
	links := &md.Files.Links
	sb := Sidebar{Links: links}
	return tmpl.RenderHString(&sb)
}
