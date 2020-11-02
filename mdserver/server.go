package mdserver

/********************************************************************************/
/*                                MdServer                                      */
/*                                                                              */
/********************************************************************************/

import (
	"bytes"
	"fmt"
	"gomd/mdserver/ws"
	u "gomd/util"
	"net/http"
	"os"
	"strconv"
	"time"
)

const sleepDuration = 1000

//MdServer is an http server used for displaying rendered markdown files
type MdServer struct {
	server        *http.Server
	bindHost      string
	bindPort      int
	path          string
	Files         *MdFiles
	theme         string
	darkMode      bool
	showHidden    bool
	isSidebarOpen bool
	hub           *ws.Hub
}

//NewMdServer initializes MdServer
func NewMdServer(bindHost string, bindPort int, path, theme string, showHidden, quiet, debug bool) MdServer {
	u.InitLog(!quiet, debug)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		u.Logln(u.Warn, "Specified path doesn't exist. Using default.")
		path = "./"
	}

	files := NewMdFiles(path, showHidden)

	md := MdServer{
		server:        nil,
		bindHost:      bindHost,
		bindPort:      bindPort,
		path:          path,
		Files:         &files,
		theme:         theme,
		darkMode:      true,
		isSidebarOpen: false,
		showHidden:    showHidden,
		hub:           ws.NewHub(),
	}

	s := &http.Server{
		Handler:      md.ServeMuxHandler(),
		Addr:         md.BindAddr(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	md.server = s

	return md
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
	md := NewMdServer(
		*opts.BindAddr,
		port,
		*opts.Dir,
		*opts.Theme,
		*opts.ShowHidden,
		*opts.Quiet,
		*opts.Debug,
	)

	return md
}

//BindAddr returns binding address of this server.
func (md *MdServer) BindAddr() string {
	return fmt.Sprintf("%v:%v", md.bindHost, md.bindPort)
}

//URL returns a url of mdserver
func (md *MdServer) URL() string {
	return "http://" + md.BindAddr()
}

//IsDarkMode returns true if dark mode is on
func (md *MdServer) IsDarkMode() bool {
	return md.darkMode
}

func (md *MdServer) SetDarkMode(on bool) {
	md.darkMode = on
}

//SetTheme - Set theme of markdown code snippets
func (md *MdServer) SetTheme(theme string) {
	md.theme = theme
}

//OpenURL opens server's url in default web browser
func (md *MdServer) OpenURL() {
	u.URLOpen(md.URL())
}

//sendReload sends a "reload" message that is then broadcasted to a websocket which
//reloads a webpage
func (md *MdServer) sendReload() {
	message := bytes.TrimSpace([]byte("reload"))
	md.hub.Broadcast <- message
}

// Serve starts up MdServer
func (md *MdServer) Serve() {
	u.Logf(u.Info, "Listening at %v", md.URL())
	u.Logf(u.Info, "Directory: %v", md.path)
	u.Logf(u.Info, "Theme: %v", md.theme)

	changed := make(chan bool)
	newFound := make(chan bool)

	go md.hub.Run()
	go md.listenForReload(changed)
	go md.listenForReload(newFound)
	go md.Files.Watch(changed, newFound)
	go md.OpenURL()
	u.LogFatal(md.server.ListenAndServe())
}

//listenForReload on receiveing true message sends a reload to websocket
//responsible for page reload.
func (md *MdServer) listenForReload(c chan bool) {
	for {
		v := <-c
		if v {
			u.Logln(u.Debug, "Sending reload")
			md.sendReload()
			md.Files.RegenerateLinks()
		}
	}
}

//Run parses commandline opts and prints help if necessary otherwise starts mdserver with
//provided options
func Run() {
	opts := ParseMdOpts()
	opts.CheckHelp()
	md := FromOpts(opts)
	md.Serve()
}
