package mdserver

import (
	"bytes"
	"gomd/util"
	"regexp"
	"strings"

	h "github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

const (
	codeTagLangStart = "<pre><code class=\"language-"
	codeTagStart     = "<pre><code"
	codeTagEnd       = "</code></pre>"
	styleStart       = "<style type=\"text/css\">"
	bodyBegin        = "<body class=\"chroma\">"
)

//codeBlock represents code block with specified programming language
//found in html file
type codeBlock struct {
	code, lang               string
	codeStartIdx, codeEndIdx int
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
	formatter := h.New(h.LineNumbersInTable(true))

	iterator, err := lexer.Tokenise(nil, string(cb.code))
	if err != nil {
		return "", err
	}
	buf := bytes.Buffer{}
	err = formatter.Format(&buf, s, iterator)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func unescapeHTML(text string) string {
	text = strings.ReplaceAll(text, "&quot;", "\"")
	text = strings.ReplaceAll(text, "&gt;", ">")
	text = strings.ReplaceAll(text, "&lt;", "<")
	text = strings.ReplaceAll(text, "&amp;", "&")
	text = strings.ReplaceAll(text, "&apos;", "'")
	return text
}

func findBlocks(html string) []codeBlock {
	blocks := []codeBlock{}

	reg := regexp.MustCompile(`<pre><code class="language-(\w+)">((?:(.|\n)*?)+?)</code></pre>`)
	idxs := reg.FindAllStringSubmatchIndex(html, -1)
	blks := reg.FindAllStringSubmatch(html, -1)
	for i := 0; i < len(blks); i++ {
		blocks = append(blocks, codeBlock{code: unescapeHTML(blks[i][2]), lang: blks[i][1], codeStartIdx: idxs[i][4], codeEndIdx: idxs[i][5]})
	}

	return blocks
}

func highlightBlocks(html, style string, blocks []codeBlock) string {
	diff := 0
	for _, block := range blocks {
		highlighted, err := block.highlightBlock(style)
		if err != nil {
			continue
		}

		html = util.StrReplace(html, highlighted, block.codeStartIdx+diff, block.codeEndIdx+diff)
		diff += len(highlighted) - (block.codeEndIdx - block.codeStartIdx)
	}
	return html
}

//HighlightHTML extracts parsed markdown blocks from html and
//replaces them with highlighted with specified style html code
//with inlined style
func HighlightHTML(html, style string) string {
	blocks := findBlocks(html)
	return highlightBlocks(html, style, blocks)
}
