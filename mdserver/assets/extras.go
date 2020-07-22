package assets

import (
	h "gomd/html"
	"sort"
	"strings"
)

//Custom html elements
const (
	themeSlider        = `<label class="switch"><input type="checkbox" onclick="themeChange(this);"><span class="slider"></span></label>`
	themeSliderChecked = `<label class="switch"><input type="checkbox" checked="checked" onclick="themeChange(this);"><span class="slider"></span></label>`
)

//backBtn returns a back button html
func backBtn(href, text string) string {
	tag := h.NewTag("a")
	tag.AddAttr("href", href)
	tag.AddAttr("class", "bbut")
	tag.SetContent(text)
	return tag.Render()
}

//themeDropdown returns a theme dropdown html
func themeDropdown(themes []string) string {
	links := ""
	for _, theme := range themes {
		tag := h.NewTag("a")
		tag.AddAttr("onclick", "codeHlChange(this);")
		tag.SetContent(theme)
		links += tag.Render()
	}
	bbut := h.A("", "Themes")
	bbut.AddAttr("class", "bbut")
	return h.Render(h.Div("dropdown", bbut.Render()+h.Render(h.Div("dropdown-content", links))))
}

//TopBar returns a TopBar with theme slider, back button and dropdown theme chooser
func TopBar(isDarkMode bool) string {
	if isDarkMode {
		return h.Render(h.Div("top-bar", openSidebarBtn()+backBtn("/", "<<")+themeDropdown(Themes())+themeSlider))
	}

	return h.Render(h.Div("top-bar", openSidebarBtn()+backBtn("/", "<<")+themeDropdown(Themes())+themeSliderChecked))
}

//TopBarSliderDropdown returns a div with theme slider and theme dropdown menu
func TopBarSliderDropdown(isDarkMode bool) string {
	if isDarkMode {
		return h.Render(h.Div("top-bar", themeDropdown(Themes())+themeSlider))
	}

	return h.Render(h.Div("top-bar", themeDropdown(Themes())+themeSliderChecked))
}

func openSidebarBtn() string {
	btn := h.NewTag("button")
	btn.AddAttr("class", "openbtn")
	btn.AddAttr("onclick", "openNav();")
	btn.SetContent("&#9776;")
	return h.Render(btn)
}

// Sidebar returns a div with class 'sidebar'. Values are hrefs
// and keys are displayed strings. The map is sorted by keys alphabeticaly.
func Sidebar(links map[string]string) string {
	as := strings.Builder{}
	closeBtn := h.A("javascript:void(0)", "x")
	closeBtn.AddAttr("onclick", "closeNav()")
	closeBtn.AddAttr("class", "closebtn")
	as.WriteString(h.Render(closeBtn))

	keys := make([]string, 0, len(links))
	for k := range links {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, name := range keys {
		as.WriteString(h.Render(h.A(links[name], name)))
	}
	return h.Render(h.Div("sidebar", as.String()))
}
