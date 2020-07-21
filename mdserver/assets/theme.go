package assets

//Themes - available themes in mdserver
func Themes() []string {
	return []string{
		"dracula",
		"paraiso",
		"monokai",
		"solarized",
		"github",
		"vs",
		"xcode",
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

//ChromaName returns a theme name compatible with chroma based on isDarkMode
func ChromaName(theme string, isDarkMode bool) string {
	switch theme {
	case "paraiso":
		if isDarkMode {
			return "paraiso-dark"
		}
		return "paraiso-light"
	case "monokai":
		if isDarkMode {
			return "monokai"
		}
		return "monokailight"
	case "solarized":
		if isDarkMode {
			return "solarized-dark"
		}
		return "solarized-light"
	default:
		return theme
	}
}
