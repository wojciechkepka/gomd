package mdserver

import (
	"github.com/gomarkdown/markdown"
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
		Content:  content,
	}, nil
}

// Checks modification time. If changed updates Mod_time
func (f *MdFile) CheckModTime() error {
	info, err := os.Stat(f.Path)
	if err != nil {
		return err
	}

	if f.Mod_time != info.ModTime() {
		f.Mod_time = info.ModTime()
	}

	return nil
}

// Creates HTML string with this file contents
func (f *MdFile) AsHtml(isDarkMode bool) string {
	body, style := TopBar(isDarkMode), MdFileStyle(isDarkMode)

	body += string(markdown.ToHTML(f.Content, nil, nil))
	return Html(f.Filename, style, body)
}
