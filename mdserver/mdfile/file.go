package mdfile

/********************************************************************************/
/*                                  MdFile                                      */
/*                                                                              */
/********************************************************************************/

import (
	u "gomd/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

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

//LoadMdFiles - Walks through a specified directory and finds md files
func LoadMdFiles(path string) []MdFile {
	files := []MdFile{}

	paths, err := filepath.Glob(path + "/*")
	if err != nil {
		return files
	}
	for _, p := range paths {
		f, err := os.Open(p)
		if err != nil {
			continue
		}
		info, err := f.Stat()
		if err != nil {
			continue
		}
		if !info.IsDir() {
			file, err := LoadMdFile(p)
			if err != nil {
				u.Logln(u.Error, "Failed to load file - ", err)
				continue
			}
			files = append(files, file)
		}
	}
	return files
}

//IsHidden check whether this file is hidden
func (f *MdFile) IsHidden() bool {
	return f.Path[0] == '.'
}
