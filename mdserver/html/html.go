package html

import "fmt"

func Themes() []string {
	return []string{
		"gruvbox",
		"solarized",
	}
}

func IsInThemes(theme string) bool {
	for _, t := range Themes() {
		if t == theme {
			return true
		}
	}

	return false
}

func ThemeCss(isDarkMode bool, theme string) string {
	switch theme {
	case "gruvbox":
		if isDarkMode {
			return GRUVBOX_DARK_HLJS
		} else {
			return GRUVBOX_LIGHT_HLJS
		}
	case "solarized":
		if isDarkMode {
			return SOLARIZED_DARK_HLJS
		} else {
			return SOLARIZED_LIGHT_HLJS
		}
	default:
		return ""
	}
}

//################################################################################
// HTML
const FILES_TITLE = "gomd - Files"

const DOCTYPE = "<!DOCTYPE html>"
const META_CHARSET = "<meta charset=\"utf-8\">"
const META_VIEWPORT = "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">"

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
const STYLE_BEG = "<style>"
const STYLE_END = "</style>"

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
function codeHlChange(a_theme) {
	var rq = new XMLHttpRequest();
	rq.open("GET", "/theme/" + a_theme.textContent, true);
	rq.onreadystatechange = reload;
	rq.send();
}
function reload() {
	location.reload();
}

`

const HIGHLIGHT_JS = `
<link rel="stylesheet"
      href="//cdnjs.cloudflare.com/ajax/libs/highlight.js/10.1.1/styles/default.min.css">
<script src="//cdnjs.cloudflare.com/ajax/libs/highlight.js/10.1.1/highlight.min.js"></script>
<script>hljs.initHighlightingOnLoad();</script>
`

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

func MdFileStyle(isDarkMode bool, theme string) string {
	if isDarkMode {
		return STYLE_BEG + FONTS + GHMD + GHMD_DARK + STYLE + ThemeCss(isDarkMode, theme) + NL + STYLE_END
	} else {
		return STYLE_BEG + FONTS + GHMD + GHMD_LIGHT + STYLE + ThemeCss(isDarkMode, theme) + NL + STYLE_END
	}
}

func FileListViewStyle(isDarkMode bool) string {
	if isDarkMode {
		return STYLE_BEG + FONTS + FV_COMMON + FV_DARK + STYLE + NL + STYLE_END
	} else {
		return STYLE_BEG + FONTS + FV_COMMON + FV_LIGHT + STYLE + NL + STYLE_END
	}

}

func ReloadJs(bindAddr string) string {
    return fmt.Sprintf(`
<script>
function tryConnectToReload(address) {
  var conn = new WebSocket(address);

  conn.onclose = function() {
    setTimeout(function() {
      tryConnectToReload(address);
    }, 2000);
  };

  conn.onmessage = function(evt) {
    location.reload()
  };
}

try {
  if (window["WebSocket"]) {
    // The reload endpoint is hosted on a statically defined port.
    try {
      tryConnectToReload("ws://%v/reload");
    }
    catch (ex) {
      // If an exception is thrown, that means that we couldn't connect to to WebSockets because of mixed content
      // security restrictions, so we try to connect using wss.
      tryConnectToReload("wss://%v/reload");
    }
  } else {
    console.log("Your browser does not support WebSockets, cannot connect to the Reload service.");
  }
} catch (ex) {
  console.error('Exception during connecting to Reload:', ex);
}
</script>
`, bindAddr, bindAddr)
}
