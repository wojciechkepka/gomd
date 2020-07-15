package mdserver

import (
	"bytes"
	"fmt"
	html "gomd/mdserver/html"
	"gomd/mdserver/ws"
	u "gomd/util"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	sleepDuration = 1000
	httpPrefix    = "http://"
	filesTitle    = "gomd - Files"

	// Endpoints
	filelistviewEp = "/"
	fileviewEp     = "/file/"
	themeEp        = "/theme/"
	themeLightEp   = "/theme/light"
	themeDarkEp    = "/theme/dark"
	staticEp       = "/static/"
	reloadEp       = "/reload"
	pingEp         = "/ping"
)

//################################################################################
// Server

var (
	hub *ws.Hub
)

//MdServer - http server used for displaying rendered markdown files
type MdServer struct {
	bindHost   string
	bindPort   int
	path       string
	Files      []MdFile
	theme      string
	darkMode   bool
	showHidden bool
}

//NewMdServer - Initializes MdServer
func NewMdServer(bindHost string, bindPort int, path, theme string, showHidden, quiet bool) MdServer {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		u.Logln(u.Warn, "Specified path doesn't exist. Using default.")
		path = "./"
	}

	if quiet {
		u.IsVerbose = false
	}

	files := LoadFiles(path)

	return MdServer{
		bindHost:   bindHost,
		bindPort:   bindPort,
		path:       path,
		Files:      files,
		theme:      theme,
		darkMode:   true,
		showHidden: showHidden,
	}
}

//BindAddr - Returns binding address of this server.
func (md *MdServer) BindAddr() string {
	return fmt.Sprintf("%v:%v", md.bindHost, md.bindPort)
}

//URL - Returns a url of mdserver
func (md *MdServer) URL() string {
	return httpPrefix + md.BindAddr()
}

//IsDarkMode - Returns true if dark mode is on
func (md *MdServer) IsDarkMode() bool {
	return md.darkMode
}

//SetDarkMode - Set value of md.darkMode field
func (md *MdServer) SetDarkMode(on bool) {
	md.darkMode = on
}

//SetTheme - Set theme of markdown code snippets
func (md *MdServer) SetTheme(theme string) {
	md.theme = theme
}

//WatchFiles - Loops endlessly checking all md.Files whether they changed
//also runs FindNewFiles on each loop
func (md *MdServer) WatchFiles() {
	for {
		for i := 0; i < len(md.Files); i++ {
			f := &md.Files[i]
			if hasChanged, _ := f.HasModTimeChanged(); hasChanged {
				u.Logf(u.Info, "File %v changed. Reloading.", f.Filename)
				err := f.ReloadMdFile()
				if err != nil {
					u.Logln(u.Warn, "Failed to reload file - ", err)
				}
				sendReload()
			}
		}
		md.FindNewFiles()
		time.Sleep(sleepDuration * time.Millisecond)
	}
}

//FindNewFiles - Checks for new files in md.Path
func (md *MdServer) FindNewFiles() {
	err := filepath.Walk(md.path, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() && !u.IsSubDirPath(md.path, p) {
			if !md.isFileInFiles(p) {
				u.Logf(u.Info, "New file found - '%v'", p)
				file, err := LoadMdFile(p)
				if err != nil {
					u.Logln(u.Error, "Failed to load file - ", err)
					return nil
				}
				md.Files = append(md.Files, file)
				sendReload()
			}

		}
		return nil
	})
	if err != nil {
		u.Logf(u.Error, "Error: failed to read directory %v - %v", md.path, err)
	}

}

//########################################n
// Other

func (md *MdServer) isFileInFiles(path string) bool {
	for _, f := range md.Files {
		if f.Path == path {
			return true
		}
	}
	return false
}

//OpenURL - opens server's url in default web browser
func (md *MdServer) OpenURL() {
	u.URLOpen(md.URL())
}

//########################################
// Html methods

// Serves markdown file as html
func (md *MdServer) serveFile(path string) string {
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

//########################################
// Server methods

// Handler for FileView
func (md *MdServer) fileViewHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.RequestURI[len(fileviewEp)-1:]
	u.Logf(u.Info, "Serving file %v", filePath)
	fmt.Fprintln(w, string(md.serveFile(filePath)))
}

// Handler for FileListView
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
		if html.IsInThemes(theme) {
			u.Logf(u.Info, "Changing theme to %v", theme)
			md.SetTheme(theme)
		}
	}
}

func (md *MdServer) watchHandler(w http.ResponseWriter, r *http.Request) {
	ws.ServeWs(hub, w, r)
}

func (md *MdServer) pingHandler(w http.ResponseWriter, r *http.Request) {
	u.Logln(u.Info, "Ping")
	fmt.Fprintln(w, "pong")
}

// Serve - Mount all endpoints and serve...
func (md *MdServer) Serve() {
	u.Logf(u.Info, "Listening at %v", md.URL())
	u.Logf(u.Info, "Directory: %v", md.path)
	u.Logf(u.Info, "Theme: %v", md.theme)
	fs := http.FileServer(http.Dir("./static"))
	hub = ws.NewHub()
	go hub.Run()
	http.HandleFunc(filelistviewEp, md.fileListViewHandler)
	http.HandleFunc(fileviewEp, md.fileViewHandler)
	http.HandleFunc(themeEp, md.themeHandler)
	http.HandleFunc(reloadEp, md.watchHandler)
	http.HandleFunc(pingEp, md.pingHandler)
	http.Handle(staticEp, http.StripPrefix(staticEp, fs))
	go md.WatchFiles()
	go md.OpenURL()
	u.LogFatal(http.ListenAndServe(md.BindAddr(), nil))
}

func sendReload() {
	message := bytes.TrimSpace([]byte("reload"))
	hub.Broadcast <- message
}
