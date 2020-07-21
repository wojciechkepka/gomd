package assets

//MdFileStyle returns style for markdown file
func MdFileStyle(isDarkMode bool, theme string) string {
	if isDarkMode {
		return Fonts + Ghmd + GhmdDark + CSS + TopBarDark
	}

	return Fonts + Ghmd + GhmdLight + CSS + TopBarLight
}

//FileListViewStyle returns style for main file list view
func FileListViewStyle(isDarkMode bool) string {
	if isDarkMode {
		return Fonts + FvCommon + FvDark + CSS + TopBarDark
	}

	return Fonts + FvCommon + FvLight + CSS + TopBarLight
}
