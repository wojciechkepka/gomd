package mdserver

import (
	u "gomd/util"
	"os"
	"path/filepath"
	"time"
)

//checkIfFilesChanged loops over all current files and checks if modification time changed
//if it changed sends a reload message to a hub
func (md *MdServer) checkIfFilesChanged() {
	for i := 0; i < len(md.Files); i++ {
		f := &md.Files[i]
		if hasChanged, _ := f.HasModTimeChanged(); hasChanged {
			u.Logf(u.Info, "File %v changed. Reloading.", f.Filename)
			err := f.ReloadMdFile()
			if err != nil {
				u.Logln(u.Warn, "Failed to reload file - ", err)
			}
			md.sendReload()
		}
	}
}

//findNewFiles checks for new files in md.Path. If it finds a new file sends a reload message
//to a hub
func (md *MdServer) findNewFiles() {
	err := filepath.Walk(md.path, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() && !u.IsSubDirPath(md.path, p) {
			if !md.isFileInFiles(p) {
				u.Logf(u.Info, "New file found - '%v'", p)
				file, err := LoadMdFile(p)
				if err != nil {
					u.Logln(u.Error, "Failed to load file - ", err)
					return nil
				}
				md.Files = append(md.Files, file)
				md.sendReload()
			}

		}
		return nil
	})
	if err != nil {
		u.Logf(u.Error, "Error: failed to read directory %v - %v", md.path, err)
	}

}

//watchFiles - Loops endlessly checking all md.Files whether they changed
//also runs findNewFiles on each loop
func (md *MdServer) watchFiles() {
	for {
		md.checkIfFilesChanged()
		md.findNewFiles()
		time.Sleep(sleepDuration * time.Millisecond)
	}
}

func (md *MdServer) isFileInFiles(path string) bool {
	for _, f := range md.Files {
		if f.Path == path {
			return true
		}
	}
	return false
}
