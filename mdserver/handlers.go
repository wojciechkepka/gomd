package mdserver

/********************************************************************************/
/*                                  Handlers                                    */
/*                                                                              */
/********************************************************************************/

import (
	"fmt"
	"gomd/mdserver/assets"
	"gomd/mdserver/ws"
	u "gomd/util"
	"net/http"
	"path/filepath"
)

func (md *MdServer) fileViewHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.RequestURI[len(fileviewEp)-1:]
	u.Logf(u.Info, "Serving file %v", filePath)
	fmt.Fprintln(w, string(md.serveFileAsHTML(filePath)))
}

func (md *MdServer) fileListViewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, md.filesHTML())
}

func (md *MdServer) themeHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.RequestURI()
	if url == themeDarkEp {
		u.Logln(u.Info, "Switching theme to dark")
		md.SetDarkMode(true)
	} else if url == themeLightEp {
		u.Logln(u.Info, "Switching theme to light")
		md.SetDarkMode(false)
	} else {
		_, theme := filepath.Split(url)
		if assets.IsInThemes(theme) {
			u.Logf(u.Info, "Changing theme to %v", theme)
			md.SetTheme(theme)
		}
	}
}

func (md *MdServer) watchHandler(w http.ResponseWriter, r *http.Request) {
	ws.ServeWs(md.hub, w, r)
}

func (md *MdServer) pingHandler(w http.ResponseWriter, r *http.Request) {
	u.Logln(u.Info, "Ping")
	fmt.Fprintln(w, "pong")
}

func (md *MdServer) sidebarHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.RequestURI()
	if url == sidebarCloseEp {
		md.isSidebarOpen = false
	} else if url == sidebarOpenEp {
		md.isSidebarOpen = true
	} else if url == sidebarCheckEp {
		if md.isSidebarOpen {
			fmt.Fprintln(w, "open")
		} else {
			fmt.Fprintln(w, "close")
		}
	}
}
