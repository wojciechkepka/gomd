package util

import (
	"fmt"
	"os"
	"os/exec"
	rt "runtime"
	"strings"
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

//StrReplace replaces a chunk of string str, starting at chunkStart
//and ending at chunkEnd, with newChunk
func StrReplace(str, newChunk string, chunkStart, chunkEnd int) string {
	first := str[:chunkStart]
	second := str[chunkEnd:]
	return first + newChunk + second
}

//UnescapeHTML replaces html escaped symbols with their
//rendered equivalent like `&quot;` -> `"`
func UnescapeHTML(text string) string {
	text = strings.ReplaceAll(text, "&quot;", "\"")
	text = strings.ReplaceAll(text, "&gt;", ">")
	text = strings.ReplaceAll(text, "&lt;", "<")
	text = strings.ReplaceAll(text, "&amp;", "&")
	text = strings.ReplaceAll(text, "&apos;", "'")
	return text
}

func fileExtension(name string) string {
	if len(name) > 0 {
		startsWithDot := name[0] == '.'
		elems := strings.Split(name, ".")

		if (len(elems) == 2 && startsWithDot) || len(elems) == 1 {
			return ""
		}

		return "." + elems[len(elems)-1]
	}
	return ""
}

// FileExtension returns file extension of passed os.FileInfo if such exists
func FileExtension(info os.FileInfo) string {
	return fileExtension(info.Name())
}
