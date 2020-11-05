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
	"os/exec"
	rt "runtime"
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
	noOpen        bool
	isSidebarOpen bool
	isShowingDiff bool
	hub           *ws.Hub
}

//NewMdServer initializes MdServer
func NewMdServer(bindHost string, bindPort int, path, theme string, showHidden, quiet, debug, noOpen bool) MdServer {
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
		isShowingDiff: false,
		showHidden:    showHidden,
		noOpen:        noOpen,
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
		*opts.NoOpen,
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
	u.Logf(u.Debug, "Setting theme to '%v'", theme)
	md.theme = theme
}

//OpenURL opens server's url in default web browser
func (md *MdServer) OpenURL() {
	u.Logf(u.Debug, "Opening '%v' in the browser", md.URL())
	u.URLOpen(md.URL())
}

//sendReload sends a "reload" message that is then broadcasted to a websocket which
//reloads a webpage
func (md *MdServer) sendReload() {
	u.Logln(u.Debug, "Sending reload to a hub")
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
	if !md.noOpen {
		go md.OpenURL()
	}
	u.LogFatal(md.server.ListenAndServe())
}

//listenForReload on receiveing true message sends a reload to websocket
//responsible for page reload.
func (md *MdServer) listenForReload(c chan bool) {
	for {
		v := <-c
		if v {
			md.sendReload()
		}
	}
}

// undaemonArgs removes `--daemon` and `-daemon` flags from passed args
// returning a new array.
func undaemonArgs(args *[]string) []string {
	newArgs := []string{}
	for _, arg := range *args {
		if arg == "--daemon" || arg == "-daemon" {
			continue
		}

		newArgs = append(newArgs, arg)
	}
	return newArgs
}

// RunDaemon runs mdserver in background. How it is achieved varies on each platform.
// On macOS and Linux `nohup` is used to start a child process.
// On windows TODO....
func RunDaemon() {
	args := undaemonArgs(&os.Args)
	switch sys := rt.GOOS; sys {
	case "darwin":
	case "linux":
		cmd := exec.Command("nohup", args...)
		cmd.Env = os.Environ()
		cmd.Start()
	case "windows":
		args = append([]string{"/b", `""`}, args...)
		cmd := exec.Command("START", args...)
		cmd.Env = os.Environ()
		cmd.Start()
	default:
		fmt.Printf("`--daemon` not supported on '%v'", sys)
		return
	}
}

//Run parses commandline opts and prints help if necessary otherwise starts mdserver with
//provided options
func Run() {
	opts := ParseMdOpts()
	opts.CheckHelp()
	u.InitLog(!*opts.Quiet, *opts.Debug)
	if *opts.Daemon {
		RunDaemon()
	} else {
		md := FromOpts(opts)
		md.Serve()
	}
}
