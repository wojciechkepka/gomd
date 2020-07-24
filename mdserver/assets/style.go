package assets

//MdFileStyle returns style for markdown file
func MdFileStyle(isDarkMode bool, theme string) string {
	if isDarkMode {
		return Ghmd + GhmdDark + CSS
	}

	return Ghmd + GhmdLight + CSS
}

//FileListViewStyle returns style for main file list view
func FileListViewStyle(isDarkMode bool) string {
	if isDarkMode {
		return FvCommon + FvDark + CSS
	}

	return FvCommon + FvLight + CSS
}
