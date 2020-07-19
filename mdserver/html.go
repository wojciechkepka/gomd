package mdserver

import (
	"gomd/html"
	"gomd/mdserver/assets"
)

// Prepares FileListView body html
func (md *MdServer) filesBody() string {
	body := html.UlBeg
	for _, file := range md.Files {
		if file.IsHidden() && !md.showHidden {
			continue
		}
		body += html.LiBeg
		endPoint := fileviewEp + file.Path
		link := html.A(endPoint, file.Path)
		body += link.Render()
		body += html.LiEnd
	}
	body += html.UlEnd
	return html.Render(html.Div("files", body))
}

// Prepares full FileListView html
func (md *MdServer) filesHTML() string {
	h := html.New()
	h.AddMeta("viewport", "width=device-width, initial-scale=1.0")
	h.AddStyle(assets.FileListViewStyle(md.IsDarkMode()))
	h.AddScript(assets.ReloadJs(md.BindAddr()))
	h.AddScript(assets.JS)
	h.AddBodyItem(assets.TopBarSliderDropdown(md.IsDarkMode()))
	h.AddBodyItem(md.filesBody())
	return h.Render()
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
