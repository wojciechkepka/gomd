package html

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
