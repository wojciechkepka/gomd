package mdserver

/********************************************************************************/
/*                                MdServer                                      */
/*                                                                              */
/********************************************************************************/

import (
	"bytes"
	"fmt"
	. "gomd/mdserver/mdfile"
	"gomd/mdserver/ws"
	u "gomd/util"
	"net/http"
	"os"
	"strconv"
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
	reloadEp       = "/reload"
	pingEp         = "/ping"
	sidebarEp      = "/sidebar/"
	sidebarOpenEp  = "/sidebar/open"
	sidebarCloseEp = "/sidebar/close"
	sidebarCheckEp = "/sidebar/check"
)

//MdServer - http server used for displaying rendered markdown files
type MdServer struct {
	bindHost      string
	bindPort      int
	path          string
	Files         []MdFile
	theme         string
	darkMode      bool
	showHidden    bool
	isSidebarOpen bool
	hub           *ws.Hub
}

//New - Initializes MdServer
func New(bindHost string, bindPort int, path, theme string, showHidden, quiet bool) MdServer {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		u.Logln(u.Warn, "Specified path doesn't exist. Using default.")
		path = "./"
	}

	if quiet {
		u.LogEnabled(false)
	}

	files := LoadMdFiles(path)

	return MdServer{
		bindHost:      bindHost,
		bindPort:      bindPort,
		path:          path,
		Files:         files,
		theme:         theme,
		darkMode:      true,
		isSidebarOpen: false,
		showHidden:    showHidden,
		hub:           ws.NewHub(),
	}
}

//FromOpts creates MdServer instance from MdOpts
func FromOpts(opts MdOpts) MdServer {
	var port int
	var err error
	port, err = strconv.Atoi(*opts.BindPort)
	if err != nil {
		u.Logln(u.Warn, "Invalid port '", port, "' using default", DefPort)
		port, _ = strconv.Atoi(DefPort)
	}
	md := New(*opts.BindAddr, port, *opts.Dir, *opts.Theme, *opts.ShowHidden, *opts.Quiet)

	return md
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

//OpenURL - opens server's url in default web browser
func (md *MdServer) OpenURL() {
	u.URLOpen(md.URL())
}

//sendReload sends a "reload" message that is then broadcasted to a websocket which
//reloads a webpage
func (md *MdServer) sendReload() {
	message := bytes.TrimSpace([]byte("reload"))
	md.hub.Broadcast <- message
}

// Serve - Mount all endpoints and serve...
func (md *MdServer) Serve() {
	u.Logf(u.Info, "Listening at %v", md.URL())
	u.Logf(u.Info, "Directory: %v", md.path)
	u.Logf(u.Info, "Theme: %v", md.theme)

	go md.hub.Run()
	http.HandleFunc(filelistviewEp, md.fileListViewHandler)
	http.HandleFunc(fileviewEp, md.fileViewHandler)
	http.HandleFunc(themeEp, md.themeHandler)
	http.HandleFunc(reloadEp, md.watchHandler)
	http.HandleFunc(pingEp, md.pingHandler)
	http.HandleFunc(sidebarEp, md.sidebarHandler)
	go md.watchFiles()
	go md.OpenURL()
	u.LogFatal(http.ListenAndServe(md.BindAddr(), nil))
}

//Run parses commandline opts and prints help if necessary otherwise starts mdserver with
//provided options
func Run() {
	opts := ParseMdOpts()
	opts.CheckHelp()
	md := FromOpts(opts)
	md.Serve()
}
