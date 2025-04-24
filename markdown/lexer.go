package markdown

import (
	"fmt"
	"strings"
)

type TokenType string

const (
	TokenHeader        TokenType = "header"
	TokenHeader2       TokenType = "header2"
	TokenHeader3       TokenType = "header3"
	TokenHeader4       TokenType = "header4"
	TokenHeader5       TokenType = "header5"
	TokenHeader6       TokenType = "header6"
	TokenParagraph     TokenType = "paragraph"
	TokenListItem      TokenType = "listitem"
	TokenOrderListItem TokenType = "orderlistitem"
	TokenEOF           TokenType = "EOF"
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
	for l.cursor < len(l.lines) {
		line := strings.TrimSpace(l.lines[l.cursor]) // Trim spaces here!
		l.cursor++

		if strings.HasPrefix(line, "#") {
			level := strings.Count(line, "#")
			return Token{
				Type: TokenType("header" + func() string {
					if level == 1 {
						return ""
					} else {
						return fmt.Sprintf("%d", level)
					}
				}()),
				Literal: strings.TrimSpace(line[level:]),
			}
		} else if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") || strings.HasPrefix(line, "+ ") {
			return Token{Type: TokenListItem, Literal: strings.TrimSpace(line[2:])}
		} else if isOrderedListItem(line) {
			// e.g. "1. List item"
			dotIdx := strings.Index(line, ".")
			return Token{Type: TokenOrderListItem, Literal: strings.TrimSpace(line[dotIdx+1:])}
		}
		return Token{Type: TokenParagraph, Literal: strings.TrimSpace(line)}
	}

	return Token{Type: TokenEOF}
}

// Helper function for ordered list detection
func isOrderedListItem(line string) bool {
	if len(line) < 3 {
		return false
	}
	// Find a number followed by a dot and a space
	for i := 0; i < len(line); i++ {
		if line[i] < '0' || line[i] > '9' {
			if line[i] == '.' && i+1 < len(line) && line[i+1] == ' ' {
				return i > 0
			}
			return false
		}
	}
	return false
}
