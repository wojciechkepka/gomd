package html

import (
	"github.com/gomarkdown/markdown"
	"gomd/mdserver/css"
	h "gomd/mdserver/highlight"
	"gomd/mdserver/js"
	. "gomd/mdserver/mdfile"
	tmpl "gomd/mdserver/template"
	"html/template"
)

type Sidebar struct {
	Links *map[string]string
}

func (sb *Sidebar) Template() (*template.Template, error) {
	return tmpl.HTemplateFromPath("/assets/html/sidebar.html", "sidebar")
}

type FilesList struct {
	Files      *[]MdFile
	ShowHidden bool
}

func (fv *FilesList) Template() (*template.Template, error) {
	return tmpl.HTemplateFromPath("/assets/html/filesdiv.html", "fileview")
}

type ThemeDropdown struct {
	Themes *[]string
}

func (td *ThemeDropdown) Template() (*template.Template, error) {
	return tmpl.HTemplateFromPath("/assets/html/theme_dropdown.html", "theme_dropdown")
}

type Topbar struct {
	DisplayButtons, IsDarkMode, IsDiff bool
	Themes                             *[]string
}

func (tb *Topbar) ThemeDropdown() template.HTML {
	td := ThemeDropdown{Themes: tb.Themes}
	return tmpl.RenderHTML(&td)
}

func (tb *Topbar) Template() (*template.Template, error) {
	return tmpl.HTemplateFromPath("/assets/html/top_bar.html", "topbar")
}

type RenderedFileView struct {
	Diff, IsDarkMode bool
	BindAddr, Theme  string
	Links            *map[string]string
	File             *MdFile
}

func (tb *RenderedFileView) Template() (*template.Template, error) {
	return tmpl.HTemplateFromPath("/assets/html/file.html", "rendered_file")
}

func (f *RenderedFileView) SidebarHTML() template.HTML {
	sb := Sidebar{Links: f.Links}
	return tmpl.RenderHTML(&sb)
}
func (f *RenderedFileView) TopbarHTML() template.HTML {
	themes := h.Themes()
	tb := Topbar{DisplayButtons: true, IsDarkMode: f.IsDarkMode, Themes: &themes, IsDiff: f.Diff}
	return tmpl.RenderHTML(&tb)
}
func (f *RenderedFileView) HighlightedContent() string {
	return h.HighlightHTML(string(markdown.ToHTML(f.File.Content, nil, nil)), h.ChromaName(f.Theme, f.IsDarkMode))
}
func (f *RenderedFileView) RenderedContent() template.HTML {
	if !f.Diff {
		return template.HTML(f.HighlightedContent())
	} else {
		return template.HTML(`<pre>` + f.File.Diff() + `</pre>`)
	}
}
func (f *RenderedFileView) FileDisplayStyle() template.HTML {
	return template.HTML("<style>" + css.MdFileStyle(f.IsDarkMode, f.Theme) + "</style>")
}
func (f *RenderedFileView) FileDisplayScripts() template.HTML {
	reload := js.ReloadJS{BindAddr: f.BindAddr}
	js := js.JS{}
	return template.HTML("<script>" + tmpl.RenderTString(&reload) + "</script>" +
		"<script>" + tmpl.RenderTString(&js) + "</script>")
}

func RenderMdFile(f *MdFile, isDarkMode, diff, raw bool, bindAddr, theme string, links *map[string]string) string {
	rendered := RenderedFileView{
		IsDarkMode: isDarkMode,
		BindAddr:   bindAddr,
		Theme:      theme,
		Links:      links,
		File:       f,
		Diff:       diff,
	}
	if raw {
		if diff {
			return `<pre>` + f.Diff() + `</pre>`
		} else {
			return rendered.HighlightedContent()
		}
	} else {
		return tmpl.RenderHString(&rendered)
	}
}
