package file

/********************************************************************************/
/*                                    File                                      */
/*                                                                              */
/********************************************************************************/

import (
	"github.com/gomarkdown/markdown"
	diff "github.com/sergi/go-diff/diffmatchpatch"
	"gomd/mdserver/highlight"
	u "gomd/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	Markdown   = ".md"
	Go         = ".go"
	Rust       = ".rs"
	Shell      = ".sh"
	Python     = ".py"
	C          = ".c"
	CPP        = ".cpp"
	JavaScript = ".js"
	Java       = ".java"
	Ruby       = ".rb"
	YAML       = ".yaml"
	JSON       = ".json"
)

type Language struct {
	Name string
}

func NewLanguage(extension string) *Language {
	switch extension {
	case Markdown:
		return &Language{"markdown"}
	case Go:
		return &Language{"go"}
	case Rust:
		return &Language{"rust"}
	case Shell:
		return &Language{"shell"}
	case C:
		return &Language{"c"}
	case CPP:
		return &Language{"cpp"}
	case JavaScript:
		return &Language{"javascript"}
	case Java:
		return &Language{"java"}
	case Ruby:
		return &Language{"ruby"}
	case YAML:
		return &Language{"yaml"}
	case JSON:
		return &Language{"json"}
	}
	return &Language{""}
}

func (l *Language) Highlight(code, style string) (string, error) {
	return highlight.HighlightCode(code, l.Name, style)
}

// File - structure representing a file held by MdServer
type File struct {
	ModTime        time.Time
	Path           string
	Filename       string
	Size           int64
	InitialContent string
	Lang           *Language
	Content        []byte
}

// NewFile - Loads a markdown file from path loading all metadata and content
func NewFile(path string) (File, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return File{}, err
	}
	info, err := os.Stat(path)
	if err != nil {
		return File{}, err
	}

	_, file := filepath.Split(path)

	return File{
		ModTime:        info.ModTime(),
		Path:           path,
		Filename:       file,
		Size:           info.Size(),
		Lang:           NewLanguage(u.FileExtension(info)),
		InitialContent: string(content),
		Content:        content,
	}, nil
}

//Reload - Reloads this file returning error if read failed or failed
//to read metadata
func (f *File) Reload() error {
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
func (f *File) HasModTimeChanged() (bool, error) {
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
func (f *File) IsHidden() bool {
	//TODO: Add windows impl
	return f.Path[0] == '.'
}

//Diff creates a diff between InitialContent and Content returned encapsuled in pre tags.
//If splitLines is enabled wraps each line with span of class token-<del|add>.
func (f *File) Diff(splitLines bool) string {
	dmp := diff.New()
	d := dmp.DiffMain(f.InitialContent, string(f.Content), false)

	s := ""
	for _, token := range d {
		lines := strings.Split(token.Text, "\n")
		for i, line := range lines {
			switch token.Type {
			case diff.DiffEqual:
				s += line
			case diff.DiffDelete:
				s += `<span class="token-del">` + line + `</span>`
			case diff.DiffInsert:
				s += `<span class="token-add">` + line + `</span>`
			}
			if len(lines) > 1 && i != len(lines)-1 {
				s += "\n"
			}
		}
	}

	if splitLines {
		lines := strings.Split(s, "\n")
		s = ""

		for _, line := range lines {
			s += `<code class="diff-line">` + line + "</code>\n"
		}
	}

	return `<pre>` + s + `</pre>`
}

//Highlight highlights Content of this file in selected style if
//there is a lexer for it. Otherwise returns untouched content.
func (f *File) Highlight(style string) (string, error) {
	return f.Lang.Highlight(string(f.Content), style)
}

//renderMdAsHTML renders this files content as HTML using a markdown parser.
func (f *File) renderMdAsHTML() string {
	return string(markdown.ToHTML(f.Content, nil, nil))
}

//RenderHTML renders this file as HTML according to its language.
func (f *File) RenderHTML(style string) string {
	if f.IsMarkdown() {
		return highlight.HighlightHTML(f.renderMdAsHTML(), style)
	} else {
		val, _ := f.Highlight(style)
		return val
	}
}

//IsMarkdown returns true if this files language is markdown
func (f *File) IsMarkdown() bool {
	return f.Lang.Name == "markdown"
}
