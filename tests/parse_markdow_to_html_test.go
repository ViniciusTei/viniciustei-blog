package tests

import (
	"testing"

	"github.com/ViniciusTei/viniciustei-blog/markdown"
)

func TestParseMarkdownToHTML(t *testing.T) {
	markdownStr := `# Header

  This is a paragraph.
  - List item 1
  - List item 2

  * List item 1
  * List item 2

  + List item 1
  + List item 2

  1. List item 1
  2. List item 2
  3. List item 1
`

	expected := `<h1>Header</h1>
<p>This is a paragraph.</p>
<ul>
<li>List item 1</li>
<li>List item 2</li>
</ul>
<ul>
<li>List item 1</li>
<li>List item 2</li>
</ul>
<ul>
<li>List item 1</li>
<li>List item 2</li>
</ul>
<ol>
<li>List item 1</li>
<li>List item 2</li>
<li>List item 1</li>
</ol>
`

	html := markdown.MarkdownToHTML(markdownStr)
	if html != expected {
		t.Errorf("Expected HTML \n '%s' \n, got \n '%s' \n", expected, html)
	}
}
