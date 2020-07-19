package html

//TagName is an enum with all allowed HTML tags.
type TagName string

//################################################################################
// HTML Elements
const (
	HTMLTag   TagName = "html"
	HeadTag   TagName = "head"
	MetaTag   TagName = "meta"
	LinkTag   TagName = "link"
	TitleTag  TagName = "title"
	ScriptTag TagName = "script"
	StyleTag  TagName = "style"
	BodyTag   TagName = "body"
	ATag      TagName = "a"
	DivTag    TagName = "div"
	UlTag     TagName = "ul"
	LiTag     TagName = "li"
	LabelTag  TagName = "label"
	SpanTag   TagName = "span"

	Doctype   = "<!DOCTYPE html>"
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
func Title(title string) Tag {
	tag := NewTag(TitleTag)
	tag.SetContent(title)
	return tag
}

//A returns a hyperlink with link set to href and text to content
func A(href, content string) Tag {
	tag := NewTag(ATag)
	tag.AddAttr("href", href)
	tag.SetContent(content)
	return tag
}

//Body returns a body enclosed by opening and closing body tag
func Body(body string) Tag {
	tag := NewTag(BodyTag)
	tag.SetContent(body)
	return tag
}

//Div returns a div with class and content specified enclosed in div tags
func Div(class, content string) Tag {
	tag := NewTag(DivTag)
	tag.AddAttr("class", class)
	tag.SetContent(content)
	return tag
}

//Script returns a script tag with content set.
func Script(content string) Tag {
	tag := NewTag(ScriptTag)
	tag.SetContent(content)
	return tag
}

//ScriptSrc returns a script tag with src attribute set to src.
func ScriptSrc(src string) Tag {
	tag := NewTag(ScriptTag)
	tag.AddAttr("src", src)
	return tag
}

//Style returns a style tag wit content set.
func Style(content string) Tag {
	tag := NewTag(StyleTag)
	tag.SetContent(content)
	return tag
}

//MetaCharset returns a meta charset tag.
func MetaCharset(charset string) Tag {
	tag := NewTag(MetaTag)
	tag.AddAttr("charset", charset)
	tag.HasClosingTag(false)
	return tag
}

//Meta returns a meta tag with name and conent attributes.
func Meta(name, content string) Tag {
	tag := NewTag(MetaTag)
	tag.AddAttr("name", name)
	tag.AddAttr("content", content)
	tag.HasClosingTag(false)
	return tag
}

//Link returns a link tag with rel and href attributes.
func Link(rel, href string) Tag {
	tag := NewTag(LinkTag)
	tag.AddAttr("rel", rel)
	tag.AddAttr("href", href)
	tag.HasClosingTag(false)
	return tag
}
