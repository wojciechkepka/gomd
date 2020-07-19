package html

import "strings"

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

//AsHTML renders this tag as html adding all attributes and content
//and enclosing it in tags.
func (t *Tag) AsHTML() string {
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
