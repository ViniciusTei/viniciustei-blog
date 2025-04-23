package markdown

import "strings"

type TokenType string

const (
	TokenHeader    TokenType = "header"
	TokenHeader2   TokenType = "header2"
	TokenHeader3   TokenType = "header3"
	TokenHeader4   TokenType = "header4"
	TokenHeader5   TokenType = "header5"
	TokenHeader6   TokenType = "header6"
	TokenParagraph TokenType = "paragraph"
	TokenList      TokenType = "list"
	TokenEOF       TokenType = "EOF"
)

type Token struct {
	Type    TokenType
	Literal string
}

type Lexer struct {
	input  string
	lines  []string
	cursor int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:  input,
		lines:  strings.Split(input, "\n"),
		cursor: 0,
	}
}

func (l *Lexer) NextToken() Token {
	if l.cursor >= len(l.lines) {
		return Token{Type: TokenEOF}
	}

	line := l.lines[l.cursor]
	l.cursor++

	if strings.HasPrefix(line, "#") {
		level := strings.Count(line, "#")

		if level == 5 {
			return Token{Type: TokenHeader5, Literal: strings.TrimSpace(line[level:])}
		}

		if level == 4 {
			return Token{Type: TokenHeader4, Literal: strings.TrimSpace(line[level:])}
		}

		if level == 3 {
			return Token{Type: TokenHeader3, Literal: strings.TrimSpace(line[level:])}
		}

		if level == 2 {
			return Token{Type: TokenHeader2, Literal: strings.TrimSpace(line[level:])}
		}

		if level == 1 {
			return Token{Type: TokenHeader, Literal: strings.TrimSpace(line[level:])}
		}

		return Token{Type: TokenHeader6, Literal: strings.TrimSpace(line[level:])}
	} else if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") {
		return Token{Type: TokenList, Literal: strings.TrimSpace(line[2:])}
	} else if line == "" {
		return l.NextToken() // skip empty line
	}

	return Token{Type: TokenParagraph, Literal: line}
}
