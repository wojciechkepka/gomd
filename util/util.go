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

//IsSubStrAtIdx checks if subStr is a substring of str at index idx
//For example:
//	IsSubStrAtIdx("hello world", "world", 6) == true
func IsSubStrAtIdx(str, subStr string, idx int) bool {
	slice := str[idx:]
	for i := 0; i < len(subStr); i++ {
		if slice[i] != subStr[i] {
			return false
		}
	}
	return true
}
