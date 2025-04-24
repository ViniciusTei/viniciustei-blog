package markdown

func MarkdownToHTML(markdown string) string {
	lexer := NewLexer(markdown)
	parser := NewParser(lexer)
	return parser.ToHTML()
}
