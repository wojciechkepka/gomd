package html

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
	DivEnd    = "</div>"
	UlBeg     = "<ul>"
	UlEnd     = "</ul>"
	LiBeg     = "<li>"
	LiEnd     = "</li>"
	StyleBeg  = "<style>"
	StyleEnd  = "</style>"

	NL = "\n"
)

//Title returns a title string enclosed in title tags
func Title(title string) string {
	return TitleBeg + title + TitleEnd
}

//A returns a hyperlink with link set to href and text to content
func A(href, content, class string) string {
	return "<a href=\"" + href + "\" class=\"" + class + "\">" + content + "</a>"
}

//Body returns a body enclosed by opening and closing body tag
func Body(body string) string {
	return BodyBeg + body + BodyEnd
}

//Div returns a div with class and content specified enclosed in div tags
func Div(class, content string) string {
	return "<div class=\"" + class + "\">" + content + DivEnd
}

//Script returns content enclosed in <script> tags
func Script(content string) string {
	return ScriptBeg + content + ScriptEnd
}

//ScriptSrc returns a script tag with src attribute set to src.
func ScriptSrc(src string) string {
	return "<script src=\"" + src + "\"></script>"
}

//Style returns content enclosed in <style> tags
func Style(content string) string {
	return StyleBeg + content + StyleEnd
}

//MetaCharset returns a meta charset tag.
func MetaCharset(charset string) string {
	return "<meta charset=\"" + charset + "\">"
}

//Meta returns a meta tag with name and conent attributes.
func Meta(name, content string) string {
	return "<meta name=\"" + name + "\" content=\"" + content + "\">"
}

//Link returns a link tag with rel and href attributes.
func Link(rel, href string) string {
	return "<link rel=\"" + rel + "\" href=\"" + href + "\">"
}
