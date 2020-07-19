package mdserver

import (
	h "gomd/html"
	"gomd/mdserver/assets"
)

// Prepares FileListView body html
func (md *MdServer) filesBody() string {
	ul := h.NewTag(h.UlTag)
	ulContent := ""
	for _, file := range md.Files {
		if file.IsHidden() && !md.showHidden {
			continue
		}
		li := h.NewTag(h.LiTag)
		endPoint := fileviewEp + file.Path
		link := h.A(endPoint, file.Path)
		li.SetContent(link.Render())
		ulContent += li.Render()
	}
	ul.SetContent(ulContent)
	return h.Render(h.Div("files", ul.Render()))
}

// Prepares full FileListView html
func (md *MdServer) filesHTML() string {
	h := h.New()
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
