package assets

import (
	h "gomd/html"
)

//Custom html elements
const (
	themeSlider        = `<label class="switch"><input type="checkbox" onclick="themeChange(this);"><span class="slider"></span></label>`
	themeSliderChecked = `<label class="switch"><input type="checkbox" checked="checked" onclick="themeChange(this);"><span class="slider"></span></label>`
)

//BackButton returns a back button html
func BackButton(href, text string) string {
	tag := h.NewTag("a")
	tag.AddAttr("href", href)
	tag.AddAttr("class", "bbut")
	tag.SetContent(text)
	return tag.Render()
}

//ThemeDropdown returns a theme dropdown html
func ThemeDropdown(themes []string) string {
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
		return h.Render(h.Div("top-bar", themeSlider+BackButton("/", "<<")+ThemeDropdown(Themes())))
	}

	return h.Render(h.Div("top-bar", themeSliderChecked+BackButton("/", "<<")+ThemeDropdown(Themes())))
}

//TopBarSliderDropdown returns a div with theme slider and theme dropdown menu
func TopBarSliderDropdown(isDarkMode bool) string {
	if isDarkMode {
		return h.Render(h.Div("top-bar", themeSlider+ThemeDropdown(Themes())))
	}

	return h.Render(h.Div("top-bar", themeSliderChecked+ThemeDropdown(Themes())))
}
