package mdserver

import "gomd/mdserver/html"

// Prepares FileListView body html
func (md *MdServer) filesBody() string {
	body := html.UlBeg + html.NL
	for _, file := range md.Files {
		if file.IsHidden() && !md.showHidden {
			continue
		}
		body += html.LiBeg
		endPoint := fileviewEp + file.Path
		body += html.A(endPoint, file.Path)
		body += html.LiEnd + html.NL
	}
	body += html.UlEnd
	return html.Div("files", body)
}

// Prepares full FileListView html
func (md *MdServer) filesHTML() string {
	body, style := html.TopBarSliderDropdown(md.IsDarkMode()), html.FileListViewStyle(md.IsDarkMode())
	body += md.filesBody()
	style += html.ReloadJs(md.BindAddr())
	return html.HTML(filesTitle, style, body)
}

// Serves markdown file as html
func (md *MdServer) serveFileAsHTML(path string) string {
	if md.path == "." {
		path = path[1:]
	}
	for _, file := range md.Files {
		if file.Path == path {
			return file.AsHTML(md.IsDarkMode(), md.theme, md.BindAddr())
		}
	}
	return ""
}
