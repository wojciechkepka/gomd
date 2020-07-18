package html

//MdFileStyle returns style for markdown file
func MdFileStyle(isDarkMode bool, theme string) string {
	if isDarkMode {
		return StyleBeg + Fonts + Ghmd + GhmdDark + Css + TopBarDark + ThemeCSS(isDarkMode, theme) + NL + StyleEnd
	}

	return StyleBeg + Fonts + Ghmd + GhmdLight + Css + TopBarLight + ThemeCSS(isDarkMode, theme) + NL + StyleEnd
}

//FileListViewStyle returns style for main file list view
func FileListViewStyle(isDarkMode bool) string {
	if isDarkMode {
		return StyleBeg + Fonts + FvCommon + FvDark + Css + TopBarDark + NL + StyleEnd
	}

	return StyleBeg + Fonts + FvCommon + FvLight + Css + TopBarLight + NL + StyleEnd
}
