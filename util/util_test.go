package util

import "testing"

func TestReplacesString(t *testing.T) {
	before := "this sentence <start>word<end> be valid"
	want := "this sentence should be valid"
	got := StrReplace(before, "should", 14, 30)

	if want != got {
		t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", want, got)
	}
}

func TestUnescapesHTML(t *testing.T) {
	before := "&quot;Life is short, the art long.&quot; - Hippocrates"
	want := `"Life is short, the art long." - Hippocrates`
	got := UnescapeHTML(before)

	if want != got {
		t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", want, got)
	}

	before = "&lt;pre&gt;&lt;code&gt;&lt;/code&gt;&lt;/pre&gt;"
	want = "<pre><code></code></pre>"
	got = UnescapeHTML(before)

	if want != got {
		t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", want, got)
	}

	before = "&amp;sometext&apos;"
	want = "&sometext'"
	got = UnescapeHTML(before)

	if want != got {
		t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", want, got)
	}
}

func TestExtractsFileExtension(t *testing.T) {
	cases := [][]string{
		{"README.md", ".md"},
		{"double.dot.rs", ".rs"},
		{"no_extension", ""},
		{".hiddenfile", ""},
		{".hiddenfile.ext", ".ext"},
	}

	for _, want := range cases {
		got := fileExtension(want[0])

		if want[1] != got {
			t.Errorf("Error: want \n`%v`\n, got \n`%v`\n", want[1], got)
		}
	}
}
