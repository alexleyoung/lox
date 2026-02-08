package main

import (
	"fmt"
	"strings"
)

type AstPrinter struct{}

func (p *AstPrinter) Print(expr Expr) (string, error) {
	result, err := expr.Accept(p)
	if err != nil {
		return "", err
	}
	s, ok := result.(string)
	if !ok {
		return "", fmt.Errorf("unexpected type returned from accept: %T", result)
	}
	return s, nil
}

func (p *AstPrinter) VisitAssignmentExpr(expr AssignmentExpr) (any, error) {
	return p.parenthesize("=", VariableExpr{Name: expr.Name}, expr.Expr)
}

func (p *AstPrinter) VisitLogicalExpr(expr LogicalExpr) (any, error) {
	return p.parenthesize(expr.Op.Lexeme, expr.Left, expr.Right)
}

func (p *AstPrinter) VisitBinaryExpr(expr BinaryExpr) (any, error) {
	return p.parenthesize(expr.Op.Lexeme, expr.Left, expr.Right)
}

func (p *AstPrinter) VisitGroupingExpr(expr GroupingExpr) (any, error) {
	return p.parenthesize("group", expr.Expr)
}

func (p *AstPrinter) VisitLiteralExpr(expr LiteralExpr) (any, error) {
	if expr.Value == nil {
		return "nil", nil
	}
	return fmt.Sprintf("%v", expr.Value), nil
}

func (p *AstPrinter) VisitUnaryExpr(expr UnaryExpr) (any, error) {
	return p.parenthesize(expr.Op.Lexeme, expr.Expr)
}

func (p *AstPrinter) VisitVariableExpr(expr VariableExpr) (any, error) {
	return expr.Name.Lexeme, nil
}

func (p *AstPrinter) parenthesize(name string, exprs ...Expr) (string, error) {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString(name)
	for _, expr := range exprs {
		builder.WriteString(" ")
		s, err := expr.Accept(p)
		if err != nil {
			return "", err
		}
		str, ok := s.(string)
		if !ok {
			return "", fmt.Errorf("unexpected type returned from accept: %T", s)
		}
		builder.WriteString(str)
	}
	builder.WriteString(")")
	return builder.String(), nil
}
