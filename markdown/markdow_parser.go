package markdown

import "fmt"

func MarkdownToHTML(markdown string) string {
	lexer := NewLexer(markdown)
	parser := NewParser(lexer)
	nodes := parser.Parse()

	var html string
	for _, node := range nodes {
		switch node.Type {
		case NodeHeader:
			html += fmt.Sprintf("<h1>%s</h1>\n", node.Content)
		case NodeHeader2:
			html += fmt.Sprintf("<h2>%s</h2>\n", node.Content)
		case NodeHeader3:
			html += fmt.Sprintf("<h3>%s</h3>\n", node.Content)
		case NodeHeader4:
			html += fmt.Sprintf("<h4>%s</h4>\n", node.Content)
		case NodeHeader5:
			html += fmt.Sprintf("<h5>%s</h5>\n", node.Content)
		case NodeHeader6:
			html += fmt.Sprintf("<h6>%s</h6>\n", node.Content)
		case NodeParagraph:
			html += fmt.Sprintf("<p>%s</p>\n", node.Content)
		case NodeList:
			html += "<ul>\n"
			for _, child := range node.Children {
				html += fmt.Sprintf("<li>%s</li>\n", child.Content)
			}
			html += "</ul>\n"
		default:
			html += fmt.Sprintf("<p>%s</p>\n", node.Content)
		}

	}

	return html
}
