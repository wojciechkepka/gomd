package html

import "strings"

//Html is a struct responsible for creating an HTML page.
//On rendering it automatically adds all the html tags
//except for the elements inside of body.
type Html struct {
	charset                                      string
	meta                                         map[string]string
	styles, scripts, bodyItems, links, scriptSrc []string
}

//New creates a new empty Html struct with default 'utf-8' charset
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

//SetCharset sets meta charset
func (h *Html) SetCharset(charset string) {
	h.charset = charset
}

//AddMeta adds a meta line to this page
func (h *Html) AddMeta(name, content string) {
	h.meta[name] = content
}

//AddStyle adds a style to this page.
func (h *Html) AddStyle(style string) {
	h.styles = append(h.styles, style)
}

//AddScript adds a cript to this page.
func (h *Html) AddScript(script string) {
	h.scripts = append(h.scripts, script)
}

//AddBodyItem adds item to this page's body.
func (h *Html) AddBodyItem(item string) {
	h.bodyItems = append(h.bodyItems, item)
}

//AddLink adds a link tag to this page's head.
func (h *Html) AddLink(rel, href string) {
	h.links = append(h.links, Link(rel, href))
}

//AddScriptSrc adds a script with src attribute set to
//argument of this func.
func (h *Html) AddScriptSrc(src string) {
	h.scriptSrc = append(h.scriptSrc, src)
}

//Render renders the whole page adding all elements together.
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
