package highlight

import (
	"bytes"
	"gomd/util"
	"regexp"

	h "github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
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

func findBlocks(html string) []codeBlock {
	blocks := []codeBlock{}

	reg := regexp.MustCompile(`<pre><code class="language-(\w+)">((?:(.|\n)*?)+?)</code></pre>`)
	idxs := reg.FindAllStringSubmatchIndex(html, -1)
	blks := reg.FindAllStringSubmatch(html, -1)
	for i := 0; i < len(blks); i++ {
		blocks = append(blocks,
			codeBlock{
				code:         util.UnescapeHTML(blks[i][2]),
				lang:         blks[i][1],
				codeStartIdx: idxs[i][4],
				codeEndIdx:   idxs[i][5],
			},
		)
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

//HighlightHTML finds code blocks in html with language specified as class
//and replaces them with highlighted html.
func HighlightHTML(html, style string) string {
	blocks := findBlocks(html)
	return highlightBlocks(html, style, blocks)
}
