package mdserver

import (
	"github.com/gomarkdown/markdown"
	log "github.com/sirupsen/logrus"
	. "gomd/mdserver/html"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

//################################################################################
// MdFile

type MdFile struct {
	Mod_time time.Time
	Path     string
	Filename string
	Size     int64
	Content  []byte
}

// Loads a markdown file from path loading all metadata and content
func LoadMdFile(path string) (MdFile, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return MdFile{}, err
	}

	info, err := os.Stat(path)
	if err != nil {
		return MdFile{}, err
	}

	_, file := filepath.Split(path)

	return MdFile{
		Mod_time: info.ModTime(),
		Path:     path,
		Filename: file,
		Size:     info.Size(),
		Content:  content,
	}, nil
}

func (f *MdFile) ReloadMdFile() error {
	content, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return err
	}

	info, err := os.Stat(f.Path)
	if err != nil {
		return err
	}

	f.Content = content
	f.Mod_time = info.ModTime()
	f.Size = info.Size()

	return nil
}

// Checks modification time. If changed updates Mod_time
func (f *MdFile) HasModTimeChanged() (bool, error) {
	info, err := os.Stat(f.Path)
	if err != nil {
		return false, err
	}

	if f.Mod_time != info.ModTime() {
		return true, nil
	}

	return false, nil
}

// Creates HTML string with this file contents
func (f *MdFile) AsHtml(isDarkMode bool, theme, bind_addr string) string {
	body, style := TopBar(isDarkMode), MdFileStyle(isDarkMode, theme)

	style += ReloadJs(bind_addr)
	body += string(markdown.ToHTML(f.Content, nil, nil))
	return Html(f.Filename, style, body)
}

// Walks through a specified directory and finds md files
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
