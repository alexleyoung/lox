package main

import (
	"errors"
	"slices"
)

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
	expr := p.term()

	for p.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		op := p.previous()
		right := p.term()
		expr = NewBinaryExpr(op, expr, right)
	}

	return expr
}

func (p *Parser) term() Expr {
	expr := p.factor()

	for p.match(MINUS, PLUS) {
		op := p.previous()
		right := p.factor()
		expr = NewBinaryExpr(op, expr, right)
	}

	return expr
}

func (p *Parser) factor() Expr {
	expr := p.unary()

	for p.match(SLASH, STAR) {
		op := p.previous()
		right := p.unary()
		expr = NewBinaryExpr(op, expr, right)
	}

	return expr
}

func (p *Parser) unary() Expr {
	for p.match(BANG, MINUS) {
		op := p.previous()
		right := p.unary()
		return NewUnaryExpr(op, right)
	}

	return p.primary()
}

func (p *Parser) primary() Expr {
	if p.match(FALSE) {
		return NewLiteralExpr(false)
	}
	if p.match(TRUE) {
		return NewLiteralExpr(true)
	}
	if p.match(NIL) {
		return NewLiteralExpr(nil)
	}
	if p.match(NUMBER, STRING) {
		return NewLiteralExpr(p.previous().Literal)
	}

	if p.match(LEFT_PAREN) {
		expr := p.expression()
		p.consume(RIGHT_PAREN, "Expect ')' after expression.")
		return expr
	}

	// unreachable ?
	return BinaryExpr{}
}

func (p *Parser) match(types ...TokenType) bool {
	if slices.ContainsFunc(types, p.check) {
		p.advance()
		return true
	}
	return false
}

func (p *Parser) consume(t TokenType, msg string) (Token, error) {
	if p.check(t) {
		return p.advance(), nil
	}

	return Token{}, parseError(p.peek(), msg)
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

func parseError(tok Token, msg string) error {
	ParseError(tok, msg)
	// return some arbitrary error
	return errors.New("parseError")
}
