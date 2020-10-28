package html

import "testing"

func TestRendersEmptyHTML(t *testing.T) {
	want := `<!DOCTYPE html><html><head><meta charset="utf-8" ></head><body></body></html>`
	h := New()
	got := h.Render()

	if got != want {
		t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", want, got)
	}
}

func TestSetsCharset(t *testing.T) {
	want := `<!DOCTYPE html><html><head><meta charset="ISO-8859-1" ></head><body></body></html>`

	h := New()
	h.SetCharset("ISO-8859-1")
	got := h.Render()

	if got != want {
		t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", want, got)
	}
}

func TestAddsStyle(t *testing.T) {
	want := `<!DOCTYPE html><html><head><meta charset="utf-8" ><style>.btn{display:inline-block;box-sizing:border-box;color:#1c1c1c}</style></head><body></body></html>`

	h := New()
	h.AddStyle(".btn{display:inline-block;box-sizing:border-box;color:#1c1c1c}")
	got := h.Render()

	if got != want {
		t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", want, got)
	}
}

func TestAddsScript(t *testing.T) {
	want := `<!DOCTYPE html><html><head><meta charset="utf-8" ><script>function themeChange(cb){var rq=new XMLHttpRequest;if(cb.checked){rq.open("GET","/theme/light",true)}else{rq.open("GET","/theme/dark",true)}</script></head><body></body></html>`

	h := New()
	h.AddScript(`function themeChange(cb){var rq=new XMLHttpRequest;if(cb.checked){rq.open("GET","/theme/light",true)}else{rq.open("GET","/theme/dark",true)}`)
	got := h.Render()

	if got != want {
		t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", want, got)
	}
}

func TestAddsBodyItem(t *testing.T) {
	want := `<!DOCTYPE html><html><head><meta charset="utf-8" ></head><body><div><p>test123</p></div><ul><li></li><li></li></ul></body></html>`

	h := New()
	h.AddBodyItem("<div><p>test123</p></div>")
	h.AddBodyItem("<ul><li></li><li></li></ul>")
	got := h.Render()

	if got != want {
		t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", want, got)
	}
}

func TestAddsLink(t *testing.T) {
	want := `<!DOCTYPE html><html><head><meta charset="utf-8" ><link rel="stylesheet" href="style.css" ></head><body></body></html>`

	h := New()
	h.AddLink("stylesheet", "style.css")
	got := h.Render()

	if got != want {
		t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", want, got)
	}
}

func TestAddsScriptSrc(t *testing.T) {
	want := `<!DOCTYPE html><html><head><meta charset="utf-8" ><script src="assets/script.js" ></script></head><body></body></html>`

	h := New()
	h.AddScriptSrc("assets/script.js")
	got := h.Render()

	if got != want {
		t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", want, got)
	}
}

func TestRendersFullHTML(t *testing.T) {
	want := `<!DOCTYPE html><html><head><meta charset="utf-8" ><link rel="stylesheet" href="style.css" ><style>.btn{display:inline-block;box-sizing:border-box;color:#1c1c1c}</style><script>function themeChange(cb){var rq=new XMLHttpRequest;if(cb.checked){rq.open("GET","/theme/light",true)}else{rq.open("GET","/theme/dark",true)}</script><script src="static/script.js" ></script></head><body><h1>Title</h1><div><p>test123</p></div><ul><li></li><li></li></ul></body></html>`

	h := New()
	h.AddLink("stylesheet", "style.css")
	h.AddScriptSrc("static/script.js")
	h.AddScript(`function themeChange(cb){var rq=new XMLHttpRequest;if(cb.checked){rq.open("GET","/theme/light",true)}else{rq.open("GET","/theme/dark",true)}`)
	h.AddStyle(".btn{display:inline-block;box-sizing:border-box;color:#1c1c1c}")
	h.AddBodyItem("<h1>Title</h1>")
	h.AddBodyItem("<div><p>test123</p></div>")
	h.AddBodyItem("<ul><li></li><li></li></ul>")
	got := h.Render()

	if got != want {
		t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", want, got)
	}
}
