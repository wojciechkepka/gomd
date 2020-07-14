package html

import "log"

//################################################################################
// HTML Elements
const (
	Doctype      = "<!DOCTYPE html>"
	MetaCharset  = "<meta charset=\"utf-8\">"
	MetaViewport = "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">"

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

//################################################################################
// Funcs

//Title returns a title string enclosed in title tags
func Title(title string) string {
	return TitleBeg + title + TitleEnd
}

//A returns a hyperlink with link set to href and text to content
func A(href, content string) string {
	a := ABeg + href + AMid + content + AEnd
	log.Println(a)
	return a
}

//Head returns a full head with style, metadata, title and scripts included
func Head(title, extra string) string {
	return HeadBeg + NL +
		MetaCharset + MetaViewport + NL +
		Title(title) + NL +
		HighlightJs +
		ScriptBeg + JS + ScriptEnd + NL +
		extra + NL +
		HeadEnd
}

//Body returns a body enclosed by opening and closing body tag
func Body(body string) string {
	return BodyBeg + NL + body + NL + BodyEnd
}

//HTML creates a full webpage
func HTML(title, style, body string) string {
	return Doctype + NL +
		HTMLBeg + NL +
		Head(title, style) + NL +
		Body(body) + NL +
		HTMLEnd
}

//Div returns a div with class and content specified enclosed in div tags
func Div(class, content string) string {
	return "<div class=\"" + class + "\">" + content + DivEnd
}
