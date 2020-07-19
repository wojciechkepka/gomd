package html

import "strings"

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

	Doctype = "<!DOCTYPE html>"
)

//Tag is a struct for creating custom html tags
type Tag struct {
	Type       TagName
	content    string //Content is only appliable when it has a closing tag.
	attributes map[string]string
	hasClosing bool //whether this tag has a closing tag required.
}

//NewTag initializes a new instance of CustomTag
func NewTag(name TagName) Tag {
	return Tag{
		Type:       name,
		content:    "",
		attributes: make(map[string]string),
		hasClosing: true,
	}
}

//AddAttr adds an attribute to this tag's map
func (t *Tag) AddAttr(k, v string) {
	t.attributes[k] = v
}

//HasClosingTag sets whethet this tag has a closing tag
func (t *Tag) HasClosingTag(val bool) {
	t.hasClosing = val
}

//SetContent sets content of this tag. Content will only be
//visible if this tag has a closing tag so the content can
//be placed in between the tags.
func (t *Tag) SetContent(content string) {
	t.content = content
}

//Render renders this tag as html adding all attributes and content
//and enclosing it in tags.
func (t *Tag) Render() string {
	s := strings.Builder{}
	s.WriteRune('<')
	s.WriteString(string(t.Type))
	s.WriteRune(' ')
	for attr, val := range t.attributes {
		s.WriteString(attr)
		s.WriteString("=\"")
		s.WriteString(val)
		s.WriteString("\" ")
	}
	s.WriteRune('>')

	if t.hasClosing {
		s.WriteString(t.content)
		s.WriteString("</")
		s.WriteString(string(t.Type))
		s.WriteRune('>')
	}

	return s.String()
}

//Render renders this tag as html adding all attributes and content
//and enclosing it in tags.
func Render(t Tag) string {
	return t.Render()
}

//################################################################################
// Funcs for tag creation

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
