package mdserver

import (
	"fmt"
	. "gomd/mdserver/html"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const RESCAN_SLEEP_DURATION = 1000

//################################################################################
// Endpoints

const FILELISTVIEW_EP = "/"
const FILEVIEW_EP = "/file/"
const THEME_DARK_EP = "/theme/dark"
const THEME_LIGHT_EP = "/theme/light"
const RESCAN_FILES_EP = "/rescan"
const STATIC_EP = "/static/"

//################################################################################
// Server

type MdServer struct {
	bind_host string
	bind_port int
	path      string
	Files     []MdFile
	theme     string
}

// Initializes MdServer
func NewMdServer(bind_host string, bind_port int, path, theme string) MdServer {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println("Specified path doesn't exist. Using default.")
		path = "."
	}

	files := LoadFiles(path)

	return MdServer{
		bind_host: bind_host,
		bind_port: bind_port,
		path:      path,
		Files:     files,
		theme:     theme,
	}
}

func (md *MdServer) bindAddr() string {
	return fmt.Sprintf("%v:%v", md.bind_host, md.bind_port)
}

func (md *MdServer) IsDarkMode() bool {
	if md.theme == "dark" {
		return true
	} else {
		return false
	}
}

func (md *MdServer) SetTheme(theme string) {
	md.theme = theme
}

func (md *MdServer) ReloadFiles() {
	for {
		md.Files = LoadFiles(md.path)
		time.Sleep(RESCAN_SLEEP_DURATION * time.Millisecond)
	}
}

//########################################
// Other

func LoadFiles(path string) []MdFile {
	var files []MdFile
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			file, err := LoadMdFile(p)
			if err != nil {
				return err
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
	for _, file := range md.Files {
		if file.Path == path {
			return file.AsHtml(md.IsDarkMode())
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
	body, style := TopBarSlider(md.IsDarkMode()), LINK_STYLE_OTHER
	if md.IsDarkMode() {
		style += LINK_STYLE_DARK
	} else {
		style += LINK_STYLE_LIGHT
	}
	body += md.filesBody()
	return Html(FILES_TITLE, style, body)
}

//########################################
// Server methods

// Handler for FileView
func (md *MdServer) FileViewHandler(w http.ResponseWriter, r *http.Request) {
	file_path := r.RequestURI[len(FILEVIEW_EP)-1:]
	log.Printf("Serving file %v", file_path)
	fmt.Fprintf(w, string(md.serveFile(file_path)))
}

// Handler for FileListView
func (md *MdServer) FileListViewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, md.filesHtml())
}

func (md *MdServer) DarkThemeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Switching theme to dark")
	md.SetTheme("dark")
	fmt.Fprintf(w, md.filesHtml())
}
func (md *MdServer) LightThemeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Switching theme to light")
	md.SetTheme("light")
	fmt.Fprintf(w, md.filesHtml())
}
func (md *MdServer) RescanHandler(w http.ResponseWriter, r *http.Request) {
	md.ReloadFiles()
}

// Mount all endpoints and serve...
func (md *MdServer) Serve() {
	fs := http.FileServer(http.Dir("./static"))
	http.HandleFunc(FILELISTVIEW_EP, md.FileListViewHandler)
	http.HandleFunc(FILEVIEW_EP, md.FileViewHandler)
	http.HandleFunc(THEME_DARK_EP, md.DarkThemeHandler)
	http.HandleFunc(THEME_LIGHT_EP, md.LightThemeHandler)
	http.HandleFunc(RESCAN_FILES_EP, md.RescanHandler)
	http.Handle(STATIC_EP, http.StripPrefix(STATIC_EP, fs))
	go md.ReloadFiles()
	log.Fatal(http.ListenAndServe(md.bindAddr(), nil))
}
