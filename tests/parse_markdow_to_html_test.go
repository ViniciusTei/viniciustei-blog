package tests

import (
	"testing"

	"github.com/ViniciusTei/viniciustei-blog/markdown"
)

func TestParseMarkdownToHTML(t *testing.T) {
	markdownStr := `# Header

  This is a paragraph.
`

	expected := `<h1>Header</h1>
<p>This is a paragraph.</p>
`

	html := markdown.MarkdownToHTML(markdownStr)
	if html != expected {
		t.Errorf("Expected HTML \n '%s' \n, got \n '%s' \n", expected, html)
	}
}
