package html

import "fmt"

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
		%v
	  </div>
	</div> 
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
	return fmt.Sprintf(themeDropdown, links)
}

//TopBar returns a TopBar with theme slider, back button and dropdown theme chooser
func TopBar(isDarkMode bool) string {
	if isDarkMode {
		return Div("top-bar", themeSlider+BackButton("/", "<<")+ThemeDropdown(Themes()))
	}

	return Div("top-bar", themeSliderChecked+BackButton("/", "<<")+ThemeDropdown(Themes()))
}

//TopBarSliderDropdown returns a div with theme slider and theme dropdown menu
func TopBarSliderDropdown(isDarkMode bool) string {
	if isDarkMode {
		return Div("top-bar", themeSlider+ThemeDropdown(Themes()))
	}

	return Div("top-bar", themeSliderChecked+ThemeDropdown(Themes()))
}
