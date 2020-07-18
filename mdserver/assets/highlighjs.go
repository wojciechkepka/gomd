package assets

import "gomd/html"

//Location of hljs
const (
	HljsCSS = "//cdnjs.cloudflare.com/ajax/libs/highlight.js/10.1.1/styles/default.min.css"
	HljsJS  = "//cdnjs.cloudflare.com/ajax/libs/highlight.js/10.1.1/highlight.min.js"
)

//AddHighlightJsToHTML adds highlightjs to html
func AddHighlightJsToHTML(h *html.HTML) {
	h.AddLink("stylesheet", HljsCSS)
	h.AddScriptSrc(HljsJS)
	h.AddScript("hljs.initHighlightingOnLoad();")
}
