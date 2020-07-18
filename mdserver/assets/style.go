package assets

//MdFileStyle returns style for markdown file
func MdFileStyle(isDarkMode bool, theme string) string {
	if isDarkMode {
		return Fonts + Ghmd + GhmdDark + Css + TopBarDark + ThemeCSS(isDarkMode, theme)
	}

	return Fonts + Ghmd + GhmdLight + Css + TopBarLight + ThemeCSS(isDarkMode, theme)
}

//FileListViewStyle returns style for main file list view
func FileListViewStyle(isDarkMode bool) string {
	if isDarkMode {
		return Fonts + FvCommon + FvDark + Css + TopBarDark
	}

	return Fonts + FvCommon + FvLight + Css + TopBarLight
}
