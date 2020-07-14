package html

//MdFileStyle returns style for markdown file
func MdFileStyle(isDarkMode bool, theme string) string {
	if isDarkMode {
		return StyleBeg + FONTS + GHMD + GHMD_DARK + STYLE + ThemeCSS(isDarkMode, theme) + NL + StyleEnd
	}

	return StyleBeg + FONTS + GHMD + GHMD_LIGHT + STYLE + ThemeCSS(isDarkMode, theme) + NL + StyleEnd
}

//FileListViewStyle returns style for main file list view
func FileListViewStyle(isDarkMode bool) string {
	if isDarkMode {
		return StyleBeg + FONTS + FV_COMMON + FV_DARK + STYLE + NL + StyleEnd
	}

	return StyleBeg + FONTS + FV_COMMON + FV_LIGHT + STYLE + NL + StyleEnd
}
