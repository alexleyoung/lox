package main

import "fmt"

type Expr interface {
	parenthesize() string
}

type BinaryExpr struct {
	Left, Right Expr
	Op          Token
}

type GroupingExpr struct {
	Expr Expr
}

type LiteralExpr struct {
	Value any
}

type UnaryExpr struct {
	Expr Expr
	Op   Token
}

func NewBinaryExpr(op Token, left, right Expr) BinaryExpr {
	return BinaryExpr{Op: op, Left: left, Right: right}
}

func NewGroupingExpr(expr Expr) GroupingExpr {
	return GroupingExpr{Expr: expr}
}

func NewLiteralExpr(value any) LiteralExpr {
	return LiteralExpr{Value: value}
}

func NewUnaryExpr(op Token, expr Expr) UnaryExpr {
	return UnaryExpr{Op: op, Expr: expr}
}

func (e BinaryExpr) parenthesize() string {
	return "(" + e.Op.String() + " " + e.Left.parenthesize() + " " + e.Right.parenthesize() + ")"
}

func (e GroupingExpr) parenthesize() string {
	return "(group " + e.Expr.parenthesize() + ")"
}

func (e LiteralExpr) parenthesize() string {
	if e.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", e.Value)
}

func (e UnaryExpr) parenthesize() string {
	return "(" + e.Op.String() + " " + e.Expr.parenthesize() + ")"
}
