package mdserver

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	html "gomd/mdserver/html"
	"gomd/mdserver/ws"
	util "gomd/util"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	SLEEP_DURATION = 1000
	HTTP           = "http://"
	FILES_TITLE    = "gomd - Files"

	// Endpoints
	FILELISTVIEW_EP = "/"
	FILEVIEW_EP     = "/file/"
	THEME_EP        = "/theme/"
	THEME_LIGHT_EP  = "/theme/light"
	THEME_DARK_EP   = "/theme/dark"
	STATIC_EP       = "/static/"
	RELOAD_EP       = "/reload"
)

//################################################################################
// Server

var hub *ws.Hub

type MdServer struct {
	bindHost string
	bindPort int
	path     string
	Files    []MdFile
	theme    string
	darkMode bool
}

// Initializes MdServer
func NewMdServer(bindHost string, bindPort int, path, theme string) MdServer {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println("Specified path doesn't exist. Using default.")
		path = "./"
	}

	files := LoadFiles(path)

	return MdServer{
		bindHost: bindHost,
		bindPort: bindPort,
		path:     path,
		Files:    files,
		theme:    theme,
		darkMode: true,
	}
}

func (md *MdServer) BindAddr() string {
	return fmt.Sprintf("%v:%v", md.bindHost, md.bindPort)
}

func (md *MdServer) Url() string {
	return HTTP + md.BindAddr()
}

func (md *MdServer) IsDarkMode() bool {
	return md.darkMode
}

func (md *MdServer) SetDarkMode(on bool) {
	md.darkMode = on
}

func (md *MdServer) SetTheme(theme string) {
	md.theme = theme
}

func (md *MdServer) WatchFiles() {
	for {
		for i := 0; i < len(md.Files); i++ {
			f := &md.Files[i]
			if hasChanged, _ := f.HasModTimeChanged(); hasChanged {
				log.Printf("File %v changed. Reloading.", f.Filename)
				err := f.ReloadMdFile()
				if err != nil {
					log.Fatalf("Failed to reload file - %v", err)
				}
				sendReload()
			}
		}
		md.FindNewFiles()
		time.Sleep(SLEEP_DURATION * time.Millisecond)
	}
}

func (md *MdServer) FindNewFiles() {
	err := filepath.Walk(md.path, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if !md.isFileInFiles(p) {
				log.Printf("New file found - '%v'", p)
				file, err := LoadMdFile(p)
				if err != nil {
					log.Fatalf("Failed to load file - %v", err)
					return nil
				}
				md.Files = append(md.Files, file)
				sendReload()
			}

		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error: failed to read directory %v - %v", md.path, err)
	}

}

//########################################
// Other

func (md *MdServer) isFileInFiles(path string) bool {
	for _, f := range md.Files {
		if f.Path == path {
			return true
		}
	}
	return false
}

func (md *MdServer) OpenUrl() {
	util.UrlOpen(md.Url())
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
			return file.AsHtml(md.IsDarkMode(), md.theme, md.BindAddr())
		}
	}
	return ""
}

// Prepares FileListView body html
func (md *MdServer) filesBody() string {
	body := html.UL_BEG + html.NL
	for _, file := range md.Files {
		body += html.LI_BEG
		end_point := FILEVIEW_EP + file.Path
		body += fmt.Sprintf(html.A_BEG, end_point)
		body += file.Path
		body += html.A_END
		body += html.LI_END + html.NL
	}
	body += html.UL_END
	return html.Div("files", body)
}

// Prepares full FileListView html
func (md *MdServer) filesHtml() string {
	body, style := html.TopBarSliderDropdown(md.IsDarkMode()), html.FileListViewStyle(md.IsDarkMode())
	body += md.filesBody()
	style += html.ReloadJs(md.BindAddr())
	return html.Html(FILES_TITLE, style, body)
}

//########################################
// Server methods

// Handler for FileView
func (md *MdServer) fileViewHandler(w http.ResponseWriter, r *http.Request) {
	file_path := r.RequestURI[len(FILEVIEW_EP)-1:]
	log.Printf("Serving file %v", file_path)
	fmt.Fprintf(w, string(md.serveFile(file_path)))
}

// Handler for FileListView
func (md *MdServer) fileListViewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, md.filesHtml())
}

func (md *MdServer) themeHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.RequestURI()
	if url == THEME_DARK_EP {
		log.Println("Switching theme to dark")
		md.SetDarkMode(true)
	} else if url == THEME_LIGHT_EP {
		log.Println("Switching theme to light")
		md.SetDarkMode(false)
	} else {
		_, theme := filepath.Split(url)
		if html.IsInThemes(theme) {
			log.Printf("Changing theme to %v", theme)
			md.SetTheme(theme)
		}
	}
}

func (md *MdServer) watchHandler(w http.ResponseWriter, r *http.Request) {
	ws.ServeWs(hub, w, r)
}

// Serve - Mount all endpoints and serve...
func (md *MdServer) Serve() {
	log.Printf("Listening at %v", md.Url())
	log.Printf("Directory: %v", md.path)
	log.Printf("Theme: %v", md.theme)
	fs := http.FileServer(http.Dir("./static"))
	hub = ws.NewHub()
	go hub.Run()
	http.HandleFunc(FILELISTVIEW_EP, md.fileListViewHandler)
	http.HandleFunc(FILEVIEW_EP, md.fileViewHandler)
	http.HandleFunc(THEME_EP, md.themeHandler)
	http.HandleFunc(RELOAD_EP, md.watchHandler)
	http.Handle(STATIC_EP, http.StripPrefix(STATIC_EP, fs))
	go md.WatchFiles()
	go md.OpenUrl()
	log.Fatal(http.ListenAndServe(md.BindAddr(), nil))
}

func sendReload() {
	message := bytes.TrimSpace([]byte("reload"))
	hub.Broadcast <- message
}
