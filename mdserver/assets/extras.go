package assets

import (
	"fmt"
	h "gomd/html"
)

//Custom html elements
const (
	backButton  = `<a href="%v" class="bbut">%v</a>`
	themeSlider = `
	<label class="switch">
		<input type="checkbox" onclick="themeChange(this);">
		<span class="slider"></span>
	</label>
`
	themeSliderChecked = `
	<label class="switch">
		<input type="checkbox" checked="checked" onclick="themeChange(this);">
		<span class="slider"></span>
	</label>
`
	themeDropdown = `
	<div class="dropdown">
	  <a class="bbut">Themes</a>
	  <div class="dropdown-content">
`
	themeAOnClick = `<a onclick="codeHlChange(this);" >%v</a>`
)

//BackButton returns a back button html
func BackButton(href, text string) string {
	return fmt.Sprintf(backButton, href, text)
}

//ThemeDropdown returns a theme dropdown html
func ThemeDropdown(themes []string) string {
	links := ""
	for _, theme := range themes {
		links += fmt.Sprintf(themeAOnClick, theme)
	}
	return h.Div("dropdown", h.A("", "Themes", "bbut")+h.Div("dropdown-content", links))
}

//TopBar returns a TopBar with theme slider, back button and dropdown theme chooser
func TopBar(isDarkMode bool) string {
	if isDarkMode {
		return h.Div("top-bar", themeSlider+BackButton("/", "<<")+ThemeDropdown(Themes()))
	}

	return h.Div("top-bar", themeSliderChecked+BackButton("/", "<<")+ThemeDropdown(Themes()))
}

//TopBarSliderDropdown returns a div with theme slider and theme dropdown menu
func TopBarSliderDropdown(isDarkMode bool) string {
	if isDarkMode {
		return h.Div("top-bar", themeSlider+ThemeDropdown(Themes()))
	}

	return h.Div("top-bar", themeSliderChecked+ThemeDropdown(Themes()))
}
