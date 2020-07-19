package html

import "strings"

//HTML is a struct responsible for creating an HTML page.
//On rendering it automatically adds all the html tags
//except for the elements inside of body.
type HTML struct {
	charset                                      string
	meta                                         map[string]string
	styles, scripts, bodyItems, links, scriptSrc []string
}

//New creates a new empty Html struct with default 'utf-8' charset
func New() HTML {
	return HTML{
		charset:   "utf-8",
		meta:      make(map[string]string),
		styles:    []string{},
		scripts:   []string{},
		scriptSrc: []string{},
		bodyItems: []string{},
		links:     []string{},
	}
}

//SetCharset sets meta charset
func (h *HTML) SetCharset(charset string) {
	h.charset = charset
}

//AddMeta adds a meta line to this page
func (h *HTML) AddMeta(name, content string) {
	h.meta[name] = content
}

//AddStyle adds a style to this page.
func (h *HTML) AddStyle(style string) {
	h.styles = append(h.styles, style)
}

//AddScript adds a cript to this page.
func (h *HTML) AddScript(script string) {
	h.scripts = append(h.scripts, script)
}

//AddBodyItem adds item to this page's body.
func (h *HTML) AddBodyItem(item string) {
	h.bodyItems = append(h.bodyItems, item)
}

//AddLink adds a link tag to this page's head.
func (h *HTML) AddLink(rel, href string) {
	h.links = append(h.links, Render(Link(rel, href)))
}

//AddScriptSrc adds a script with src attribute set to
//argument of this func.
func (h *HTML) AddScriptSrc(src string) {
	h.scriptSrc = append(h.scriptSrc, src)
}

//Render renders the whole page adding all elements together.
func (h *HTML) Render() string {
	s := strings.Builder{}
	s.WriteString(Doctype)
	html := NewTag(HTMLTag)

	//Head
	head, headContent := NewTag(HeadTag), strings.Builder{}
	headContent.WriteString(Render(MetaCharset(h.charset)))
	for name, content := range h.meta {
		headContent.WriteString(Render(Meta(name, content)))
	}
	for _, link := range h.links {
		headContent.WriteString(link)
	}
	for _, style := range h.styles {
		headContent.WriteString(Render(Style(style)))
	}
	for _, script := range h.scripts {
		headContent.WriteString(Render(Script(script)))
	}
	for _, src := range h.scriptSrc {
		headContent.WriteString(Render(ScriptSrc(src)))
	}
	head.SetContent(headContent.String())

	//Body
	body, bodyContent := NewTag(BodyTag), strings.Builder{}
	for _, item := range h.bodyItems {
		bodyContent.WriteString(item)
	}
	body.SetContent(bodyContent.String())

	html.SetContent(head.Render() + body.Render())
	s.WriteString(html.Render())
	return s.String()
}
