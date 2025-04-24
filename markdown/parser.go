package markdown

import (
	"fmt"
	"regexp"
	"strings"
)

type NodeType string

const (
	NodeHeader        NodeType = "header"
	NodeHeader2       NodeType = "header2"
	NodeHeader3       NodeType = "header3"
	NodeHeader4       NodeType = "header4"
	NodeHeader5       NodeType = "header5"
	NodeHeader6       NodeType = "header6"
	NodeParagraph     NodeType = "paragraph"
	NodeListItem      NodeType = "listitem"
	NodeOrderListItem NodeType = "orderlistitem"
	NodeCodeBlock     NodeType = "codeblock"
)

type Node struct {
	Type     NodeType
	Content  string
	Children []*Node
}

type Parser struct {
	lexer  *Lexer
	tokens []Token
	cursor int
}

func NewParser(lexer *Lexer) *Parser {
	var tokens []Token
	for {
		token := lexer.NextToken()
		tokens = append(tokens, token)
		if token.Type == TokenEOF {
			break
		}
	}

	return &Parser{
		lexer:  lexer,
		tokens: tokens,
		cursor: 0,
	}
}

func (p *Parser) Parse() []*Node {
	var nodes []*Node

	for {
		token := p.peek()
		if token.Type == TokenEOF {
			break
		}

		switch token.Type {
		case TokenHeader:
			nodes = append(nodes, &Node{Type: NodeHeader, Content: token.Literal})
			p.advance()
		case TokenHeader2:
			nodes = append(nodes, &Node{Type: NodeHeader2, Content: token.Literal})
			p.advance()
		case TokenHeader3:
			nodes = append(nodes, &Node{Type: NodeHeader3, Content: token.Literal})
			p.advance()
		case TokenHeader4:
			nodes = append(nodes, &Node{Type: NodeHeader4, Content: token.Literal})
			p.advance()
		case TokenHeader5:
			nodes = append(nodes, &Node{Type: NodeHeader5, Content: token.Literal})
			p.advance()
		case TokenHeader6:
			nodes = append(nodes, &Node{Type: NodeHeader6, Content: token.Literal})
			p.advance()
		case TokenParagraph:
			if token.Literal == "" {
				p.advance()
				continue
			}
			nodes = append(nodes, &Node{Type: NodeParagraph, Content: token.Literal})
			p.advance()
		case TokenListItem:
			listNode := &Node{Type: NodeListItem}
			for {
				token := p.peek()
				if token.Type != TokenListItem {
					break
				}
				listNode.Children = append(listNode.Children, &Node{Type: NodeListItem, Content: token.Literal})
				p.advance()
				// Stop if next token is a paragraph with empty content (blank line)
				if p.peek().Type == TokenParagraph && p.peek().Literal == "" {
					p.advance()
					break
				}
			}
			nodes = append(nodes, listNode)
		case TokenOrderListItem:
			listNode := &Node{Type: NodeOrderListItem}
			for {
				token := p.peek()
				if token.Type != TokenOrderListItem {
					break
				}
				listNode.Children = append(listNode.Children, &Node{Type: NodeOrderListItem, Content: token.Literal})
				p.advance()
				if p.peek().Type == TokenParagraph && p.peek().Literal == "" {
					p.advance()
					break
				}
			}
			nodes = append(nodes, listNode)
		case TokenCodeBlock:
			nodes = append(nodes, &Node{Type: NodeCodeBlock, Content: token.Literal})
			p.advance()
		default:
			p.advance()
		}
	}

	return nodes
}

func formatInline(text string) string {
	// Bold-italic: ***text*** or ___text___
	reBoldItalic := regexp.MustCompile(`(\*\*\*|___)(.*?)\1`)
	text = reBoldItalic.ReplaceAllString(text, "<strong><em>$2</em></strong>")
	// Bold: **text** or __text__
	reBold := regexp.MustCompile(`(\*\*|__)(.*?)\1`)
	text = reBold.ReplaceAllString(text, "<strong>$2</strong>")
	// Italic: *text* or _text_
	reItalic := regexp.MustCompile(`(\*|_)(.*?)\1`)
	text = reItalic.ReplaceAllString(text, "<em>$2</em>")
	// Links: [text](url)
	reLink := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	text = reLink.ReplaceAllString(text, `<a href="$2">$1</a>`)
	// Images: ![alt](url)
	reImage := regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`)
	text = reImage.ReplaceAllString(text, `<img src="$2" alt="$1">`)
	return text
}

func (p *Parser) ToHTML() string {
	nodes := p.Parse()
	var html string
	for _, node := range nodes {
		switch node.Type {
		case NodeHeader:
			html += fmt.Sprintf("<h1>%s</h1>\n", formatInline(node.Content))
		case NodeHeader2:
			html += fmt.Sprintf("<h2>%s</h2>\n", formatInline(node.Content))
		case NodeHeader3:
			html += fmt.Sprintf("<h3>%s</h3>\n", formatInline(node.Content))
		case NodeHeader4:
			html += fmt.Sprintf("<h4>%s</h4>\n", formatInline(node.Content))
		case NodeHeader5:
			html += fmt.Sprintf("<h5>%s</h5>\n", formatInline(node.Content))
		case NodeHeader6:
			html += fmt.Sprintf("<h6>%s</h6>\n", formatInline(node.Content))
		case NodeParagraph:
			html += fmt.Sprintf("<p>%s</p>\n", formatInline(node.Content))
		case NodeListItem:
			html += "<ul>\n"
			for _, child := range node.Children {
				html += fmt.Sprintf("<li>%s</li>\n", formatInline(child.Content))
			}
			html += "</ul>\n"
		case NodeOrderListItem:
			html += "<ol>\n"
			for _, child := range node.Children {
				html += fmt.Sprintf("<li>%s</li>\n", formatInline(child.Content))
			}
			html += "</ol>\n"
		case NodeCodeBlock:
			html += "<pre><code>" + escapeHTML(node.Content) + "</code></pre>\n"
		}
	}

	return html
}
func escapeHTML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return s
}

func (p *Parser) peek() Token {
	return p.tokens[p.cursor]
}

func (p *Parser) advance() {
	p.cursor++
}
