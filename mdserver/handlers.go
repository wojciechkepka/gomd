package mdserver

/********************************************************************************/
/*                                  Handlers                                    */
/*                                                                              */
/********************************************************************************/

import (
	"fmt"
	"gomd/mdserver/highlight"
	"gomd/mdserver/ws"
	u "gomd/util"
	"net/http"
	"net/url"
	"path/filepath"
)

const (
	// Endpoints
	filelistviewEp = "/"
	fileviewEp     = "/file/"
	themeEp        = "/theme/"
	themeLightEp   = "/theme/light"
	themeDarkEp    = "/theme/dark"
	reloadEp       = "/reload"
	pingEp         = "/ping"
	sidebarEp      = "/sidebar/"
	sidebarOpenEp  = "/sidebar/open"
	sidebarCloseEp = "/sidebar/close"
	sidebarCheckEp = "/sidebar/check"
)

func (md *MdServer) fileViewHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.RequestURI[len(fileviewEp)-1:]
	path, err := url.QueryUnescape(filePath)
	if err != nil {
		u.Logln(u.Error, err)
		fmt.Fprintln(w, "Invalid path -", filePath)
		return
	}
	u.Logf(u.Info, "Serving file %v", path)
	fmt.Fprintln(w, string(md.serveFileAsHTML(path)))
}

func (md *MdServer) fileListViewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, md.MainViewHTML())
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
		if highlight.IsInThemes(theme) {
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

// MdServerMuxHandler returns a ServeMux with all endpoint handlers attached
func (md *MdServer) ServeMuxHandler() *http.ServeMux {
	sm := http.NewServeMux()
	sm.HandleFunc(filelistviewEp, md.fileListViewHandler)
	sm.HandleFunc(fileviewEp, md.fileViewHandler)
	sm.HandleFunc(themeEp, md.themeHandler)
	sm.HandleFunc(reloadEp, md.watchHandler)
	sm.HandleFunc(pingEp, md.pingHandler)
	sm.HandleFunc(sidebarEp, md.sidebarHandler)
	return sm
}
