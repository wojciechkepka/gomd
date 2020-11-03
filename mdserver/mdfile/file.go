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
	u.Logf(u.Debug, "Reloading file '%v'", f.Filename)
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
	u.Logf(u.Debug, "Checking if file '%v' changed.", f.Filename)
	info, err := os.Stat(f.Path)
	if err != nil {
		return false, err
	}
	first := f.ModTime
	second := info.ModTime()
	if first != second {
		u.Logf(u.Debug, "Modtime changed from '%v' to '%v'", first, second)
		return true, nil
	}

	return false, nil
}

//IsHidden check whether this file is hidden
func (f *MdFile) IsHidden() bool {
	return f.Path[0] == '.'
}
