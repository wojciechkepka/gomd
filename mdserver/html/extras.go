package html

import "fmt"

const (
	BACK_BUTTON  = `<a href="%v" class="bbut">%v</a>`
	THEME_SLIDER = `
	<label class="switch">
		<input type="checkbox" onclick="themeChange(this);">
		<span class="slider"></span>
	</label>
`
	THEME_SLIDER_CHECKED = `
	<label class="switch">
		<input type="checkbox" checked="checked" onclick="themeChange(this);">
		<span class="slider"></span>
	</label>
`
	THEME_DROPDOWN = `
	<div class="dropdown">
	  <a class="bbut">Themes</a>
	  <div class="dropdown-content">
		%v
	  </div>
	</div> 
`
	THEME_A_ONCLICK = `<a onclick="codeHlChange(this);" >%v</a>`
)

func BackButtonHtml(href, text string) string {
	return fmt.Sprintf(BACK_BUTTON, href, text)
}

func ThemeDropdownHtml(themes []string) string {
	links := ""
	for _, theme := range themes {
		links += fmt.Sprintf(THEME_A_ONCLICK, theme)
	}
	return fmt.Sprintf(THEME_DROPDOWN, links)
}
