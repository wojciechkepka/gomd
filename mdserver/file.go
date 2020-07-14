package mdserver

import (
	"gomd/mdserver/html"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/gomarkdown/markdown"
	log "github.com/sirupsen/logrus"
)

//################################################################################
// MdFile

// MdFile - structure representing a markdown file
type MdFile struct {
	ModTime  time.Time
	Path     string
	Filename string
	Size     int64
	Content  []byte
}

// LoadMdFile - Loads a markdown file from path loading all metadata and content
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
		ModTime:  info.ModTime(),
		Path:     path,
		Filename: file,
		Size:     info.Size(),
		Content:  content,
	}, nil
}

//ReloadMdFile - Reloads md file returning error if read failed or failed
//to read metadata
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
	f.ModTime = info.ModTime()
	f.Size = info.Size()

	return nil
}

//HasModTimeChanged - Checks modification time. If changed updates ModTime
func (f *MdFile) HasModTimeChanged() (bool, error) {
	info, err := os.Stat(f.Path)
	if err != nil {
		return false, err
	}

	if f.ModTime != info.ModTime() {
		return true, nil
	}

	return false, nil
}

//AsHTML - Creates HTML string with this file contents
func (f *MdFile) AsHTML(isDarkMode bool, theme, bindAddr string) string {
	body, style := html.TopBar(isDarkMode), html.MdFileStyle(isDarkMode, theme)

	style += html.ReloadJs(bindAddr)
	body += string(markdown.ToHTML(f.Content, nil, nil))
	return html.HTML(f.Filename, style, body)
}

//LoadFiles - Walks through a specified directory and finds md files
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

//IsHidden check whether this file is hidden
func (f *MdFile) IsHidden() bool {
	return f.Path[0] == '.'
}
