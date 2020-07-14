package html

//Themes - available themes in mdserver
func Themes() []string {
	return []string{
		"gruvbox",
		"solarized",
	}
}

//IsInThemes - checks whether theme is in available themes
func IsInThemes(theme string) bool {
	for _, t := range Themes() {
		if t == theme {
			return true
		}
	}

	return false
}

//ThemeCSS - returns style of specified theme
func ThemeCSS(isDarkMode bool, theme string) string {
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
