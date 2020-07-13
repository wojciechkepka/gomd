package mdserver

import (
	"bytes"
	"fmt"
	. "gomd/mdserver/html"
	"gomd/mdserver/ws"
	util "gomd/util"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const SLEEP_DURATION = 1000
const HTTP = "http://"

//################################################################################
// Endpoints

const FILELISTVIEW_EP = "/"
const FILEVIEW_EP = "/file/"
const THEME_EP = "/theme/"
const THEME_LIGHT_EP = "/theme/light"
const THEME_DARK_EP = "/theme/dark"
const STATIC_EP = "/static/"
const RELOAD_EP = "/reload"

//################################################################################
// Server

var hub *ws.Hub

type MdServer struct {
	bind_host string
	bind_port int
	path      string
	Files     []MdFile
	theme     string
	darkMode  bool
}

// Initializes MdServer
func NewMdServer(bind_host string, bind_port int, path, theme string) MdServer {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println("Specified path doesn't exist. Using default.")
		path = "./"
	}

	files := LoadFiles(path)

	return MdServer{
		bind_host: bind_host,
		bind_port: bind_port,
		path:      path,
		Files:     files,
		theme:     theme,
		darkMode:  true,
	}
}

func (md *MdServer) BindAddr() string {
	return fmt.Sprintf("%v:%v", md.bind_host, md.bind_port)
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

func (md *MdServer) OpenUrl() {
	util.UrlOpen(md.Url())
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

func LoadFiles(path string) []MdFile {
	var files []MdFile
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			file, err := LoadMdFile(p)
			if err != nil {
				log.Fatalf("Failed to load file - %v", err)
				return nil
			}
			files = append(files, file)
		}
		return nil
	})
	if err != nil {
		log.Printf("Error: failed to read file - %v", err)
	}
	return files
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
	html := UL_BEG + NL
	for _, file := range md.Files {
		html += LI_BEG
		end_point := FILEVIEW_EP + file.Path
		html += fmt.Sprintf(A_BEG, end_point)
		html += file.Path
		html += A_END
		html += LI_END + NL
	}
	html += UL_END
	return Div("files", html)
}

// Prepares full FileListView html
func (md *MdServer) filesHtml() string {
	body, style := TopBarSliderDropdown(md.IsDarkMode()), FileListViewStyle(md.IsDarkMode())
	body += md.filesBody()
	style += ReloadJs(md.BindAddr())
	return Html(FILES_TITLE, style, body)
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
		if IsInThemes(theme) {
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
