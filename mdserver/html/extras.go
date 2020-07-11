package html

import "fmt"

const BACK_BUTTON = `<a href="%v" class="bbut">%v</a>`

const THEME_SLIDER = `
	<label class="switch">
		<input type="checkbox" onclick="themeChange(this);">
		<span class="slider"></span>
	</label>
`
const THEME_SLIDER_CHECKED = `
	<label class="switch">
		<input type="checkbox" checked="checked" onclick="themeChange(this);">
		<span class="slider"></span>
	</label>
`

func BackButtonHtml(href, text string) string {
	return fmt.Sprintf(BACK_BUTTON, href, text)
}
