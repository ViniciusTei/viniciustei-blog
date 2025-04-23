package markdown

type NodeType string

const (
	NodeHeader    NodeType = "header"
	NodeHeader2   NodeType = "header2"
	NodeHeader3   NodeType = "header3"
	NodeHeader4   NodeType = "header4"
	NodeHeader5   NodeType = "header5"
	NodeHeader6   NodeType = "header6"
	NodeParagraph NodeType = "paragraph"
	NodeList      NodeType = "list"
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
		case TokenHeader2:
			nodes = append(nodes, &Node{Type: NodeHeader2, Content: token.Literal})
		case TokenHeader3:
			nodes = append(nodes, &Node{Type: NodeHeader3, Content: token.Literal})
		case TokenHeader4:
			nodes = append(nodes, &Node{Type: NodeHeader4, Content: token.Literal})
		case TokenHeader5:
			nodes = append(nodes, &Node{Type: NodeHeader5, Content: token.Literal})
		case TokenHeader6:
			nodes = append(nodes, &Node{Type: NodeHeader6, Content: token.Literal})
		case TokenParagraph:
			nodes = append(nodes, &Node{Type: NodeParagraph, Content: token.Literal})
		case TokenList:
			listNode := &Node{Type: NodeList}
			listNode.Children = append(listNode.Children, &Node{Type: NodeParagraph, Content: token.Literal})
			nodes = append(nodes, listNode)

		}
		p.advance()
	}
	return nodes
}

func (p *Parser) peek() Token {
	return p.tokens[p.cursor]
}

func (p *Parser) advance() {
	p.cursor++
}
