package tests

import (
	"testing"

	"github.com/ViniciusTei/viniciustei-blog/markdown"
)

func TestParseMarkdownToHTML(t *testing.T) {
	markdownStr := `# Header
This is a paragraph.

- Item 1
- Item 2`

	expected := `<h1>Header</h1>
<p>This is a paragraph.</p>
<ul>
<li>Item 1</li>
<li>Item 2</li>
</ul>
`

	html := markdown.MarkdownToHTML(markdownStr)
	if html != expected {
		t.Errorf("Expected HTML \n '%s' \n, got \n '%s' \n", expected, html)
	}
}
