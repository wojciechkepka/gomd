package html

import "fmt"

//################################################################################
// HTML

const (
	DOCTYPE       = "<!DOCTYPE html>"
	META_CHARSET  = "<meta charset=\"utf-8\">"
	META_VIEWPORT = "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">"

	HTML_BEG   = "<html>"
	HTML_END   = "</html>"
	BODY_BEG   = "<body>"
	BODY_END   = "</body>"
	HEAD_BEG   = "<head>"
	HEAD_END   = "</head>"
	TITLE_BEG  = "<title>"
	TITLE_END  = "</title>"
	SCRIPT_BEG = "<script>"
	SCRIPT_END = "</script>"
	A_BEG      = "<a href=\"%v\">"
	A_END      = "</a>"
	DIV_END    = "</div>\n"
	UL_BEG     = "<ul>"
	UL_END     = "</ul>"
	LI_BEG     = "<li>"
	LI_END     = "</li>"
	STYLE_BEG  = "<style>"
	STYLE_END  = "</style>"

	DARK_BG      = "#1c1c1c"
	LIGHT_BG     = "#ffffff"
	LIGHT_TEXT   = "#f8efe1"
	DARK_TEXT    = "#000000"
	LIGHT_BORDER = "#eaeaea"
	DARK_BORDER  = "#666666"

	NL = "\n"
)

//################################################################################
// Funcs

func HtmlTitle(title string) string {
	return TITLE_BEG + title + TITLE_END
}

func HtmlA(href, content string) string {
	return fmt.Sprintf(A_BEG, href) + content + A_END
}

func HtmlHead(title, extra string) string {
	return HEAD_BEG + NL +
		META_CHARSET + META_VIEWPORT + NL +
		HtmlTitle(title) + NL +
		HIGHLIGHT_JS +
		SCRIPT_BEG + JS + SCRIPT_END + NL +
		extra + NL +
		HEAD_END
}

func HtmlBody(body string) string {
	return BODY_BEG + NL + body + NL + BODY_END
}

func Html(title, style, body string) string {
	return DOCTYPE + NL +
		HTML_BEG + NL +
		HtmlHead(title, style) + NL +
		HtmlBody(body) + NL +
		HTML_END
}

func Div(class, content string) string {
	return "<div class=\"" + class + "\">" + content + DIV_END
}

func TopBar(isDarkMode bool) string {
	if isDarkMode {
		return Div("top-bar", THEME_SLIDER+BackButtonHtml("/", "<<")+ThemeDropdownHtml(Themes()))
	} else {
		return Div("top-bar", THEME_SLIDER_CHECKED+BackButtonHtml("/", "<<")+ThemeDropdownHtml(Themes()))
	}
}

func TopBarSliderDropdown(isDarkMode bool) string {
	if isDarkMode {
		return Div("top-bar", THEME_SLIDER+ThemeDropdownHtml(Themes()))
	} else {
		return Div("top-bar", THEME_SLIDER_CHECKED+ThemeDropdownHtml(Themes()))
	}
}
