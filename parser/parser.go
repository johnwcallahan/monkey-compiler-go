package parser

import (
	"github.com/johnwcallahan/monkey-go/ast"
	"github.com/johnwcallahan/monkey-go/lexer"
	"github.com/johnwcallahan/monkey-go/token"
)

// Parser ...
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New ...
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram ...
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
