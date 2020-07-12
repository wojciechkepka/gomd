package mdserver

import (
	"fmt"
	. "gomd/mdserver/html"
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
const THEME_DARK_EP = "/theme/dark"
const THEME_LIGHT_EP = "/theme/light"
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
		path = "./"
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

func (md *MdServer) BindAddr() string {
	return fmt.Sprintf("%v:%v", md.bind_host, md.bind_port)
}

func (md *MdServer) Url() string {
	return HTTP + md.BindAddr()
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
	body, style := TopBarSlider(md.IsDarkMode()), FileListViewStyle(md.IsDarkMode())
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
}
func (md *MdServer) LightThemeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Switching theme to light")
	md.SetTheme("light")
}

// Mount all endpoints and serve...
func (md *MdServer) Serve() {
	log.Printf("Listening at %v", md.Url())
	log.Printf("Directory: %v", md.path)
	log.Printf("Theme: %v", md.theme)
	fs := http.FileServer(http.Dir("./static"))
	http.HandleFunc(FILELISTVIEW_EP, md.FileListViewHandler)
	http.HandleFunc(FILEVIEW_EP, md.FileViewHandler)
	http.HandleFunc(THEME_DARK_EP, md.DarkThemeHandler)
	http.HandleFunc(THEME_LIGHT_EP, md.LightThemeHandler)
	http.Handle(STATIC_EP, http.StripPrefix(STATIC_EP, fs))
	go md.WatchFiles()
	go md.OpenUrl()
	log.Fatal(http.ListenAndServe(md.BindAddr(), nil))
}
