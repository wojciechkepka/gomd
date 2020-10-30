package assets

import (
	h "gomd/html"
)

//Custom html elements
const (
	themeSlider        = `<label class="switch"><input type="checkbox" onclick="themeChange(this);"><span class="slider"></span></label>`
	themeSliderChecked = `<label class="switch"><input type="checkbox" checked="checked" onclick="themeChange(this);"><span class="slider"></span></label>`
	arrowLeft          = "&#8592;"
)

//backBtn returns a back button html
func backBtn(href, text string) string {
	tag := h.NewTag("a")
	tag.AddAttr("href", href)
	tag.AddAttr("class", "bbtn")
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
	bbtn := h.A("", "Themes")
	bbtn.AddAttr("class", "tbtn")
	return h.Render(h.Div("dropdown", bbtn.Render()+h.Render(h.Div("dropdown-content", links))))
}

//TopBar returns a TopBar with theme slider, back button and dropdown theme chooser
func TopBar(isDarkMode bool) string {
	if isDarkMode {
		return h.Render(h.Div("top-bar", openSidebarBtn()+backBtn("/", arrowLeft)+themeDropdown(Themes())+themeSlider))
	}

	return h.Render(h.Div("top-bar", openSidebarBtn()+backBtn("/", arrowLeft)+themeDropdown(Themes())+themeSliderChecked))
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
	btn.AddAttr("class", "bbtn")
	btn.AddAttr("onclick", "openNav();")
	btn.SetContent("&#9776;")
	return h.Render(btn)
}
