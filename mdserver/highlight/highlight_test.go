package highlight

import "testing"

const (
	TestHTMLSingle = `
<html>
<head>
</head>
<body>
	<pre><code class="language-python">
import sys

print("test123")
sys.exit(1)
	</code></pre>
	<pre><code>
This is just a codeblock</code></pre>
</body>
</html>`
	TestHTMLSingleAfter = `
<html>
<head>
</head>
<body>
	<pre><code class="language-python"><pre style="color:#e5e5e5;background-color:#000">
<span style="color:#fff;font-weight:bold">import</span> sys

<span style="color:#fff;font-weight:bold">print</span>(<span style="color:#0ff;font-weight:bold"></span><span style="color:#0ff;font-weight:bold">&#34;</span><span style="color:#0ff;font-weight:bold">test123</span><span style="color:#0ff;font-weight:bold">&#34;</span>)
sys.exit(<span style="color:#ff0;font-weight:bold">1</span>)
	</pre></code></pre>
	<pre><code>
This is just a codeblock</code></pre>
</body>
</html>`
	TestHTMLMultiple = `
<html>
<head>
</head>
<body>
<pre><code class="language-go">
import "fmt"

x := 15
fmt.Printf("%v", x)</code></pre>
<div><p> some random html </p></div>
<pre><code>not a codeblock</pre></code>
<pre><code class="language-rust">
fn main() {
	let s = String::new("");
	println!("it works! {}", s);
}
</code></pre>
</body>
</html>`
	TestHTMLMultipleAfter = `
<html>
<head>
</head>
<body>
<pre><code class="language-go"><pre style="color:#e5e5e5;background-color:#000">
<span style="color:#fff;font-weight:bold">import</span> <span style="color:#0ff;font-weight:bold">&#34;fmt&#34;</span>

x := <span style="color:#ff0;font-weight:bold">15</span>
fmt.Printf(<span style="color:#0ff;font-weight:bold">&#34;%v&#34;</span>, x)</pre></code></pre>
<div><p> some random html </p></div>
<pre><code>not a codeblock</pre></code>
<pre><code class="language-rust"><pre style="color:#e5e5e5;background-color:#000">
<span style="color:#fff;font-weight:bold">fn</span> main() {
	<span style="color:#fff;font-weight:bold">let</span> s = <span style="color:#fff;font-weight:bold">String</span>::new(<span style="color:#0ff;font-weight:bold">&#34;&#34;</span>);
	println!(<span style="color:#0ff;font-weight:bold">&#34;it works! {}&#34;</span>, s);
}
</pre></code></pre>
</body>
</html>`
)

func TestFindsCodeblock(t *testing.T) {
	want := []codeBlock{
		{
			lang: "python",
			code: `
import sys

print("test123")
sys.exit(1)
	`,
			codeStartIdx: 66,
			codeEndIdx:   109,
		},
	}

	got := findBlocks(TestHTMLSingle)

	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", want[i], got[i])
		}
	}
}

func TestFindsCodeblocks(t *testing.T) {
	want := []codeBlock{
		{
			lang: "go",
			code: `
import "fmt"

x := 15
fmt.Printf("%v", x)`,
			codeStartIdx: 61,
			codeEndIdx:   103,
		},
		{
			lang: "rust",
			code: `
fn main() {
	let s = String::new("");
	println!("it works! {}", s);
}
`,
			codeStartIdx: 227,
			codeEndIdx:   298,
		},
	}

	got := findBlocks(TestHTMLMultiple)

	if len(want) != len(got) {
		t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", want, got)
	}

	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", want[i], got[i])
		}
	}
}

func TestHighlightsSingleCodeblock(t *testing.T) {
	blocks := findBlocks(TestHTMLSingle)
	got := highlightBlocks(TestHTMLSingle, "solarized", blocks)

	for i, c := range got {
		if got[i] != TestHTMLSingleAfter[i] {
			t.Logf("got `%c` want `%c` at index %v", c, TestHTMLSingleAfter[i], i)
			t.Logf("`%v`", got[i-10:i+10])
		}
	}

	if got != TestHTMLSingleAfter {
		t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", TestHTMLSingleAfter, got)
	}
}

func TestHighlightsMultipleCodeblocks(t *testing.T) {
	blocks := findBlocks(TestHTMLMultiple)
	got := highlightBlocks(TestHTMLMultiple, "solarized", blocks)

	for i, c := range got {
		if got[i] != TestHTMLMultipleAfter[i] {
			t.Logf("got `%c` want `%c` at index %v", c, TestHTMLMultipleAfter[i], i)
		}
	}

	if got != TestHTMLMultipleAfter {
		t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", TestHTMLMultipleAfter, got)
	}
}

func TestHighlightsHTML(t *testing.T) {
	got := HighlightHTML(TestHTMLMultiple, "solarized")

	if got != TestHTMLMultipleAfter {
		t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", TestHTMLMultipleAfter, got)
	}
}
