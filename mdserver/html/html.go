package html

import (
	"gomd/mdserver/css"
	. "gomd/mdserver/file"
	h "gomd/mdserver/highlight"
	"gomd/mdserver/js"
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
	Files      *[]File
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
	File             *File
}

func (fv *RenderedFileView) Template() (*template.Template, error) {
	return tmpl.HTemplateFromPath("/assets/html/file.html", "rendered_file")
}

func (fv *RenderedFileView) Title() string {
	return fv.File.Filename
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
	style := h.ChromaName(f.Theme, f.IsDarkMode)
	return f.File.RenderHTML(style)
}
func (f *RenderedFileView) RenderedContent() template.HTML {
	if !f.Diff {
		return template.HTML(f.HighlightedContent())
	} else {
		return template.HTML(f.File.Diff(true))
	}
}
func (f *RenderedFileView) FileDisplayStyle() template.HTML {
	return template.HTML("<style>" + css.MdFileStyle(f.IsDarkMode, f.Theme) + "</style>")
}
func (f *RenderedFileView) FileDisplayScripts() template.HTML {
	js := js.JS{BindAddr: f.BindAddr}
	return template.HTML("<script>" + tmpl.RenderTString(&js) + "</script>")
}

func RenderMdFile(f *File, isDarkMode, diff, raw bool, bindAddr, theme string, links *map[string]string) string {
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
			return f.Diff(true)
		} else {
			return rendered.HighlightedContent()
		}
	} else {
		return tmpl.RenderHString(&rendered)
	}
}
