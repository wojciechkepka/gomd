package html

//################################################################################
// HTML
const FILES_TITLE = "gomd - Files"

const DOCTYPE = "<!DOCTYPE html>"
const META_CHARSET = "<meta charset=\"utf-8\">"
const META_VIEWPORT = "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">"

const LINK_STYLE_LIGHT = "<link rel=\"stylesheet\" href=\"/static/css/style_light.css\">\n"
const LINK_STYLE_DARK = "<link rel=\"stylesheet\" href=\"/static/css/style_dark.css\">\n"
const LINK_STYLE_OTHER = "<link rel=\"stylesheet\" href=\"/static/css/other.css\">\n"
const LINK_STYLE_GH_LIGHT = "<link rel=\"stylesheet\" href=\"/static/css/ghmd_light.css\">\n"
const LINK_STYLE_GH_DARK = "<link rel=\"stylesheet\" href=\"/static/css/ghmd_dark.css\">\n"

const HTML_BEG = "<html>"
const HTML_END = "</html>"
const BODY_BEG = "<body>"
const BODY_END = "</body>"
const HEAD_BEG = "<head>"
const HEAD_END = "</head>"
const TITLE_BEG = "<title>"
const TITLE_END = "</title>"
const SCRIPT_BEG = "<script>"
const SCRIPT_END = "</script>"
const A_BEG = "<a href=\"%v\">"
const A_END = "</a>"
const DIV_END = "</div>\n"
const UL_BEG = "<ul>"
const UL_END = "</ul>"
const LI_BEG = "<li>"
const LI_END = "</li>"

const DARK_BG = "#1c1c1c"
const LIGHT_BG = "#ffffff"
const LIGHT_TEXT = "#f8efe1"
const DARK_TEXT = "#000000"
const LIGHT_BORDER = "#eaeaea"
const DARK_BORDER = "#666666"

const NL = "\n"

//################################################################################
// JS

const JS = `
function themeChange(cb) {
	var rq = new XMLHttpRequest();

	if (cb.checked) {
		rq.open("GET", "/theme/light", true);
	} else {
		rq.open("GET", "/theme/dark", true);
	}
	rq.onreadystatechange = reload;
	rq.send();
}
function reload() {
	location.reload();
}
`

//################################################################################
// Funcs

func HtmlTitle(title string) string {
	return TITLE_BEG + title + TITLE_END
}

func HtmlHead(title, extra string) string {
	return HEAD_BEG + NL +
		META_CHARSET + META_VIEWPORT + NL +
		HtmlTitle(title) + NL +
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
		return Div("top-bar", THEME_SLIDER+BackButtonHtml("/", "<<"))
	} else {
		return Div("top-bar", THEME_SLIDER_CHECKED+BackButtonHtml("/", "<<"))
	}
}

func TopBarSlider(isDarkMode bool) string {
	if isDarkMode {
		return Div("top-bar", THEME_SLIDER)
	} else {
		return Div("top-bar", THEME_SLIDER_CHECKED)
	}
}
