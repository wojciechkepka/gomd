package util

import (
	"fmt"
	"os/exec"
	rt "runtime"
)

//URLOpen - tries to open a url based on OS
func URLOpen(url string) error {
	var cmd *exec.Cmd

	switch os := rt.GOOS; os {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("start", url)
	default:
		return fmt.Errorf("Unsupported os %v", os)
	}

	return cmd.Run()
}

//IsSubDirPath - checks if given path contains any '/'
func IsSubDirPath(basePath, path string) bool {
	return CountChInStr('/', path) > 0
}

//CountChInStr - counts ch character occurances in str string
func CountChInStr(ch rune, str string) int {
	count := 0
	for _, c := range str {
		if c == ch {
			count++
		}
	}

	return count
}
