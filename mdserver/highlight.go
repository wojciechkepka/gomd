package mdserver

import (
	"bufio"
	"bytes"
	"gomd/util"
	"strings"

	h "github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

const (
	codeTagStart = "<pre><code class=\"language-"
	codeTagEnd   = "</code></pre>"
	styleStart   = "<style type=\"text/css\">"
	bodyBegin    = "<body class=\"chroma\">"
)

//codeBlock represents code block with specified programming language
//found in html file
type codeBlock struct {
	code, lang string
}

func (cb *codeBlock) highlightBlock(style string) (string, error) {
	lexer := lexers.Get(cb.lang)
	if lexer == nil {
		lexer = lexers.Fallback
	}

	s := styles.Get(style)
	if s == nil {
		s = styles.Fallback
	}
	formatter := h.New()

	iterator, err := lexer.Tokenise(nil, string(cb.code))
	if err != nil {
		return "", err
	}
	buf := bytes.Buffer{}
	writer := bufio.NewWriter(&buf)
	err = formatter.Format(writer, s, iterator)
	if err != nil {
		return "", err
	}
	writer.Flush()
	return buf.String(), nil
}

func scanLang(html string) string {
	lang := strings.Builder{}
	for _, r := range html {
		if r == '"' {
			break
		}
		lang.WriteRune(r)
	}
	return lang.String()
}

func scanCode(html string) string {
	code := strings.Builder{}
	for i, r := range html {
		if util.IsSubStrAtIdx(html, codeTagEnd, i) {
			break
		}
		code.WriteRune(r)
	}
	out := code.String()
	out = strings.ReplaceAll(out, "&quot;", "\"")
	out = strings.ReplaceAll(out, "&gt;", ">")
	out = strings.ReplaceAll(out, "&lt;", "<")
	out = strings.ReplaceAll(out, "&amp;", "&")
	out = strings.ReplaceAll(out, "&apos;", "'")
	return out[2:]
}

func findCodeBlocks(html string) []codeBlock {
	blocks := []codeBlock{}
	for i := range html {
		if util.IsSubStrAtIdx(html, codeTagStart, i) {
			i += len(codeTagStart)
			lang := scanLang(html[i:])
			i += len(lang)
			code := scanCode(html[i:])
			i += len(code)
			blocks = append(blocks, codeBlock{code: code, lang: lang})

		}
	}
	return blocks
}

func extractBody(out string) string {
	for i := range out {
		if util.IsSubStrAtIdx(out, "<pre style", i) {
			return out[i:]
		}
	}
	return out
}

func extractCodeBlocks(html, style string) []string {
	blocks := findCodeBlocks(html)
	finalBlocks := []string{}
	for _, block := range blocks {
		out, err := block.highlightBlock(style)
		if err != nil {
			util.Logln(util.Info, err)
		}
		finalBlocks = append(finalBlocks, extractBody(out))
	}

	return finalBlocks
}

func HighlightHtml(html, style string) string {
	out := strings.Builder{}
	blocks := extractCodeBlocks(html, style)
	util.Logln(util.Info, blocks)
	blockIdx := 0
	push := true
	for i, r := range html {
		if util.IsSubStrAtIdx(html, codeTagStart, i) {
			push = false
		}
		if util.IsSubStrAtIdx(html, codeTagEnd, i) {
			out.WriteString(blocks[blockIdx])
			i += len(codeTagEnd) + 5
			blockIdx++
			push = true
		}
		if push {
			out.WriteRune(r)
		}
	}
	return out.String()
}
