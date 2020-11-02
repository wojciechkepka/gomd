package html

import (
	"bytes"
	"github.com/gobuffalo/packr"
	"github.com/gomarkdown/markdown"
	"gomd/mdserver/gen"
	h "gomd/mdserver/highlight"
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

type FilesList struct {
	Files *[]MdFile
}

func (fv *FilesList) Template() (*template.Template, error) {
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

type RenderedFileView struct {
	IsDarkMode      bool
	BindAddr, Theme string
	Links           *map[string]string
	File            *MdFile
}

func (tb *RenderedFileView) Template() (*template.Template, error) {
	return TemplateFromBox("./assets", "file.html", "rendered_file")
}

func (f *RenderedFileView) SidebarHTML() template.HTML {
	sb := Sidebar{Links: f.Links}
	return RenderHTML(&sb)
}
func (f *RenderedFileView) TopbarHTML() template.HTML {
	themes := h.Themes()
	tb := Topbar{DisplayButtons: true, IsDarkMode: f.IsDarkMode, Themes: &themes}
	return RenderHTML(&tb)
}
func (f *RenderedFileView) RenderedContent() template.HTML {
	return template.HTML(h.HighlightHTML(string(markdown.ToHTML(f.File.Content, nil, nil)), h.ChromaName(f.Theme, f.IsDarkMode)))
}
func (f *RenderedFileView) FileDisplayStyle() template.HTML {
	return template.HTML("<style>" + gen.MdFileStyle(f.IsDarkMode, f.Theme) + "</style>")
}
func (f *RenderedFileView) FileDisplayScripts() template.HTML {
	return template.HTML("<script>" + gen.ReloadJs(f.BindAddr) + "</script>" +
		"<script>" + gen.JS + "</script>")
}

func RenderMdFile(f *MdFile, isDarkMode bool, bindAddr, theme string, links *map[string]string) string {
	rendered := RenderedFileView{
		IsDarkMode: isDarkMode,
		BindAddr:   bindAddr,
		Theme:      theme,
		Links:      links,
		File:       f,
	}
	return RenderString(&rendered)
}
