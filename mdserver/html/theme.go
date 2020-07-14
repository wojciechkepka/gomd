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
			return GruvboxDarkHljs
		}

		return GruvboxLightHljs
	case "solarized":
		if isDarkMode {
			return SolarizedDarkHljs
		}

		return SolarizedLightHljs
	default:
		return ""
	}
}
