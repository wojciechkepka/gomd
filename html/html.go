package html

import "strings"

//################################################################################
// HTML Elements
const (
	Doctype = "<!DOCTYPE html>"

	HTMLBeg   = "<html>"
	HTMLEnd   = "</html>"
	BodyBeg   = "<body>"
	BodyEnd   = "</body>"
	HeadBeg   = "<head>"
	HeadEnd   = "</head>"
	TitleBeg  = "<title>"
	TitleEnd  = "</title>"
	ScriptBeg = "<script>"
	ScriptEnd = "</script>"
	ABeg      = "<a href=\""
	AMid      = "\">"
	AEnd      = "</a>"
	DivEnd    = "</div>\n"
	UlBeg     = "<ul>"
	UlEnd     = "</ul>"
	LiBeg     = "<li>"
	LiEnd     = "</li>"
	StyleBeg  = "<style>"
	StyleEnd  = "</style>"

	NL = "\n"
)

type Html struct {
	charset                                      string
	meta                                         map[string]string
	styles, scripts, bodyItems, links, scriptSrc []string
}

func New() Html {
	return Html{
		charset:   "utf-8",
		meta:      make(map[string]string),
		styles:    []string{},
		scripts:   []string{},
		scriptSrc: []string{},
		bodyItems: []string{},
		links:     []string{},
	}
}

func (h *Html) SetCharset(charset string) {
	h.charset = charset
}

func (h *Html) AddMeta(name, content string) {
	h.meta[name] = content
}

func (h *Html) AddStyle(style string) {
	h.styles = append(h.styles, style)
}

func (h *Html) AddScript(script string) {
	h.scripts = append(h.scripts, script)
}

func (h *Html) AddBodyItem(item string) {
	h.bodyItems = append(h.bodyItems, item)
}

func (h *Html) Render() string {
	s := strings.Builder{}
	s.WriteString(Doctype + HTMLBeg + HeadBeg)

	//Head
	s.WriteString(MetaCharset(h.charset))
	for name, content := range h.meta {
		s.WriteString(Meta(name, content))
	}
	for _, link := range h.links {
		s.WriteString(link)
	}
	for _, style := range h.styles {
		s.WriteString(Style(style))
	}
	for _, script := range h.scripts {
		s.WriteString(Script(script))
	}
	for _, src := range h.scriptSrc {
		s.WriteString(ScriptSrc(src))
	}
	s.WriteString(HeadEnd)

	//Body
	s.WriteString(BodyBeg)
	for _, item := range h.bodyItems {
		s.WriteString(item)
	}
	s.WriteString(BodyEnd + HTMLEnd)

	return s.String()
}

func (h *Html) AddLink(rel, href string) {
	h.links = append(h.links, Link(rel, href))
}

func (h *Html) AddScriptSrc(src string) {
	h.scriptSrc = append(h.scriptSrc, src)
}

//################################################################################
// Funcs

//Title returns a title string enclosed in title tags
func Title(title string) string {
	return TitleBeg + title + TitleEnd
}

//A returns a hyperlink with link set to href and text to content
func A(href, content string) string {
	return ABeg + href + AMid + content + AEnd
}

//Body returns a body enclosed by opening and closing body tag
func Body(body string) string {
	return BodyBeg + NL + body + NL + BodyEnd
}

//Div returns a div with class and content specified enclosed in div tags
func Div(class, content string) string {
	return "<div class=\"" + class + "\">" + content + DivEnd
}

//Script returns content enclosed in <script> tags
func Script(content string) string {
	return ScriptBeg + content + ScriptEnd
}

func ScriptSrc(src string) string {
	return "<script src=\"" + src + "\"></script>"
}

//Style returns content enclosed in <style> tags
func Style(content string) string {
	return StyleBeg + content + StyleEnd
}

func MetaCharset(charset string) string {
	return "<meta charset=\"" + charset + "\">"
}

func Meta(name, content string) string {
	return "<meta name=\"" + name + "\" content=\"" + content + "\">"
}

func Link(rel, href string) string {
	return "<link rel=\"" + rel + "\" href=\"" + href + "\">"
}
