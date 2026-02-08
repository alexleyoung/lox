package main

import (
	"slices"
)

type Parser struct {
	tokens   []Token
	current  int
	reporter *ErrorReporter
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, current: 0, reporter: NewErrorReporter()}
}

func (p *Parser) Parse() ([]Stmt, error) {
	statements := make([]Stmt, 0)

	for !p.isAtEnd() {
		statement, err := p.declaration()
		if err != nil {
			p.synchronize()
		}
		statements = append(statements, statement)
	}

	return statements, nil
}

func (p *Parser) HadError() bool {
	return p.reporter.HadError()
}

func (p *Parser) declaration() (Stmt, error) {
	if p.match(VAR) {
		stmt, err := p.varDeclaration()
		if err != nil {
			p.synchronize()
			return nil, err
		}
		return stmt, nil
	}

	return p.statement()
}

func (p *Parser) varDeclaration() (Stmt, error) {
	name, err := p.consume(IDENTIFIER, "Expect variable name.")
	if err != nil {
		return nil, err
	}

	var initializer Expr = nil
	if p.match(EQUAL) {
		initializer, err = p.expression()
		if err != nil {
			return nil, err
		}
	}

	_, err = p.consume(SEMICOLON, "Expect ';' after variable declaration.")
	if err != nil {
		return nil, err
	}
	return NewVariableStmt(name, initializer), nil
}

func (p *Parser) statement() (Stmt, error) {
	// Check non-expression statements first and leave as a fallthrough
	// "Hard to proactively recognize an expression from its first token"
	if p.match(IF) {
		return p.ifStatement()
	}
	if p.match(PRINT) {
		return p.printStatement()
	}
	if p.match(LEFT_BRACE) {
		return p.blockStatement()
	}
	return p.expressionStatement()
}

func (p *Parser) ifStatement() (Stmt, error) {
	_, err := p.consume(LEFT_PAREN, "Expect '(' after 'if'.")
	if err != nil {
		return nil, err
	}

	expr, err := p.expression()
	_, err = p.consume(RIGHT_PAREN, "Expect ')' after if condition.")
	if err != nil {
		return nil, err
	}

	thenBranch, err := p.statement()
	var elseBranch Stmt = nil
	if p.match(ELSE) {
		elseBranch, err = p.statement()
		if err != nil {
			return nil, err
		}
	}

	return NewIfStmt(expr, thenBranch, elseBranch), nil
}

func (p *Parser) printStatement() (Stmt, error) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(SEMICOLON, "Expect ';' after value.")
	if err != nil {
		return nil, err
	}
	return NewPrintStmt(expr), nil
}

func (p *Parser) blockStatement() (Stmt, error) {
	stmts := make([]Stmt, 0)
	for !p.check(RIGHT_BRACE) && !p.isAtEnd() {
		stmt, err := p.declaration()
		if err != nil {
			return nil, err
		}
		stmts = append(stmts, stmt)
	}
	_, err := p.consume(RIGHT_BRACE, "Expected closing brace '}'.")
	return NewBlockStmt(stmts), err
}

func (p *Parser) expressionStatement() (Stmt, error) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(SEMICOLON, "Expect ';' after value.")
	if err != nil {
		return nil, err
	}
	return NewExpressionStmt(expr), nil
}

func (p *Parser) expression() (Expr, error) {
	return p.assignment()
}

func (p *Parser) assignment() (Expr, error) {
	expr, err := p.equality()
	if err != nil {
		return nil, err
	}

	if p.match(EQUAL) {
		equals := p.previous()
		value, err := p.assignment()
		if err != nil {
			reporter.Report(err)
		}

		expr, ok := expr.(VariableExpr)
		if ok {
			name := expr.Name
			return NewAssignmentExpr(name, value), nil
		}

		reporter.Report(NewParserError(equals, "Invalid assignment target."))
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

	if p.match(IDENTIFIER) {
		return NewVariableExpr(p.previous()), nil
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

	return nil, p.parseError(p.peek(), "Failed to parse")
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

func (p *Parser) parseError(tok Token, msg string) ParserError {
	err := NewParserError(tok, msg)
	p.reporter.Report(err)
	return err
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
