package markdown

import "fmt"

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
			nodes = append(nodes, &Node{Type: NodeParagraph, Content: token.Literal})
			p.advance()
		case TokenListItem:
			listnode := &Node{Type: NodeListItem, Content: token.Literal}
			listnode.Children = append(listnode.Children, listnode)

			for {
				if token.Type != TokenListItem {
					break
				}
				listnode.Children = append(listnode.Children, &Node{Type: NodeParagraph, Content: token.Literal})
				token = p.peek()
				p.advance()
			}
		case TokenOrderListItem:
			listnode := &Node{Type: NodeOrderListItem, Content: token.Literal}
			nodes = append(nodes, listnode)

			for {
				if token.Type != TokenOrderListItem {
					break
				}
				listnode.Children = append(listnode.Children, &Node{Type: NodeParagraph, Content: token.Literal})
				token = p.peek()
				p.advance()
			}
		}
	}

	return nodes
}

func (p *Parser) ToHTML() string {
	nodes := p.Parse()
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
		case NodeListItem:
			/* TODO:
			 * The list is redenring wrong, it is wrapping every li inside its
			 * own ul. Inteads, every li should be inside the same ul.
			 */
			html += "<ul>\n"
			for _, child := range node.Children {
				html += fmt.Sprintf("<li>%s</li>\n", child.Content)
			}
			html += "</ul>\n"
		}

	}

	return html
}

func (p *Parser) peek() Token {
	return p.tokens[p.cursor]
}

func (p *Parser) advance() {
	p.cursor++
}
