package main

import "slices"

type Parser struct {
	tokens  []Token
	current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens, 0}
}

func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	expr := p.comparison()

	for p.match(BANG_EQUAL, EQUAL_EQUAL) {
		op := p.previous()
		right := p.comparison()
		expr = NewBinaryExpr(op, expr, right)
	}

	return expr
}

func (p *Parser) comparison() Expr {
	return BinaryExpr{}
}

func (p *Parser) term() Expr {
	return BinaryExpr{}
}

func (p *Parser) factor() Expr {
	return BinaryExpr{}
}

func (p *Parser) unary() Expr {
	return BinaryExpr{}
}

func (p *Parser) primary() Expr {
	return BinaryExpr{}
}

func (p *Parser) match(types ...TokenType) bool {
	if slices.ContainsFunc(types, p.check) {
		p.advance()
		return true
	}
	return false
}

func (p *Parser) check(t TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == EOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}
