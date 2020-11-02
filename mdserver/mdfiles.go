package mdserver

import (
	. "gomd/mdserver/mdfile"
	u "gomd/util"
	"os"
	"path/filepath"
	"time"
)

/// MdFiles is a control wrapper for multiple MdFiles
type MdFiles struct {
	Path       string
	Files      []MdFile
	Links      map[string]string
	ShowHidden bool
}

//NewMdFiles returns an instance of MdFiles
func NewMdFiles(path string, showHidden bool) MdFiles {
	f := loadMdFiles(path)
	return MdFiles{
		Path:       path,
		Files:      f,
		Links:      linksFromFiles(f, showHidden),
		ShowHidden: showHidden,
	}
}

//linksFromFiles returns a map of mdfile names as keys and their
//coresponding full path as values
func linksFromFiles(files []MdFile, showHidden bool) map[string]string {
	links := make(map[string]string)
	for _, f := range files {
		if f.IsHidden() && !showHidden {
			continue
		}
		links[f.Filename] = fileviewEp + f.Path
	}
	return links
}

//loadMdFiles - Walks through a specified directory and finds md files
func loadMdFiles(path string) []MdFile {
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

//checkIfFilesChanged loops over all current files and checks if modification
//time changed if it changed sends a reload message to a hub
func (md *MdFiles) checkIfFilesChanged(changed chan bool) {
	for i := 0; i < len(md.Files); i++ {
		f := &md.Files[i]
		if hasChanged, _ := f.HasModTimeChanged(); hasChanged {
			u.Logf(u.Info, "File %v changed", f.Filename)
			err := f.ReloadMdFile()
			if err != nil {
				u.Logln(u.Warn, "Failed to reload file - ", err)
			}
			changed <- true
		}
	}
}

//findNewFiles checks for new files in md.Path.
//If it finds a new file a reload message is sent to the hub.
func (md *MdFiles) findNewFiles(newFound chan bool) error {
	paths, err := filepath.Glob(md.Path + "/*")
	if err != nil {
		return err
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
			if !md.isFileInFiles(p) {
				u.Logf(u.Info, "New file found - '%v'", p)
				file, err := LoadMdFile(p)
				if err != nil {
					u.Logln(u.Error, "Failed to load file - ", err)
					continue
				}
				md.Files = append(md.Files, file)
				newFound <- true
			}

		}
	}

	return nil
}

//watchFiles loops endlessly checking whether any mdfile changed
//and looking for new files on each loop
func (md *MdFiles) Watch(changed chan bool, newFound chan bool) {
	defer func() {
		changed <- false
		newFound <- false
	}()
	for {

		md.checkIfFilesChanged(changed)
		md.findNewFiles(newFound)
		time.Sleep(sleepDuration * time.Millisecond)
	}
}

//RegenerateLinks recreates links after file change
func (md *MdFiles) RegenerateLinks() {
	md.Links = linksFromFiles(md.Files, md.ShowHidden)
}

//isFileInFiles checks if specified path is part of this server's files
func (md *MdFiles) isFileInFiles(path string) bool {
	for _, f := range md.Files {
		if f.Path == path {
			return true
		}
	}
	return false
}
