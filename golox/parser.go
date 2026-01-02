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

func (p *Parser) Parse() (Expr, error) {
	return p.expression()
}

func (p *Parser) expression() (Expr, error) {
	return p.ternary()
}

func (p *Parser) ternary() (Expr, error) {
	expr, err := p.equality()
	if err != nil {
		return nil, err
	}

	if p.match(QUESTION) {
		thenBranch, err := p.ternary()
		if err != nil {
			return nil, err
		}
		_, err = p.consume(COLON, "Expect ':' after then branch of ternary expression.")
		if err != nil {
			return nil, err
		}
		elseBranch, err := p.ternary()
		if err != nil {
			return nil, err
		}
		return NewTernaryExpr(expr, thenBranch, elseBranch), nil
	}

	return expr, nil
}

func (p *Parser) equality() (Expr, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match(BANG_EQUAL, EQUAL_EQUAL) {
		op := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		expr = NewBinaryExpr(op, expr, right)
	}

	return expr, nil
}

func (p *Parser) comparison() (Expr, error) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		op := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = NewBinaryExpr(op, expr, right)
	}

	return expr, nil
}

func (p *Parser) term() (Expr, error) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.match(MINUS, PLUS) {
		op := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expr = NewBinaryExpr(op, expr, right)
	}

	return expr, nil
}

func (p *Parser) factor() (Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.match(SLASH, STAR) {
		op := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = NewBinaryExpr(op, expr, right)
	}

	return expr, nil
}

func (p *Parser) unary() (Expr, error) {
	if p.match(BANG, MINUS) {
		op := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return NewUnaryExpr(op, right), nil
	}

	return p.primary()
}

func (p *Parser) primary() (Expr, error) {
	if p.match(FALSE) {
		return NewLiteralExpr(false), nil
	}
	if p.match(TRUE) {
		return NewLiteralExpr(true), nil
	}
	if p.match(NIL) {
		return NewLiteralExpr(nil), nil
	}
	if p.match(NUMBER, STRING) {
		return NewLiteralExpr(p.previous().Literal), nil
	}

	if p.match(LEFT_PAREN) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		_, err = p.consume(RIGHT_PAREN, "Expect ')' after expression.")
		if err != nil {
			return nil, err
		}
		return NewGroupingExpr(expr), nil
	}

	return nil, p.parseError(p.peek(), "Expect expression.")
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

	return Token{}, p.parseError(p.peek(), msg)
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

func (p *Parser) parseError(tok Token, msg string) error {
	ParseError(tok, msg)
	// return some arbitrary error
	return errors.New("parseError")
}

func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		switch p.peek().Type {
		case CLASS, FUN, VAR, FOR, IF, WHILE, PRINT, RETURN:
			return
		}
	}

	p.advance()
}
