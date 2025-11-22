package main

type Expr interface {
	accept(v Visitor) (any, error)
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

type TernaryExpr struct {
	Condition  Expr
	ThenBranch Expr
	ElseBranch Expr
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

func NewTernaryExpr(condition, thenBranch, elseBranch Expr) TernaryExpr {
	return TernaryExpr{Condition: condition, ThenBranch: thenBranch, ElseBranch: elseBranch}
}

func (e BinaryExpr) accept(v Visitor) (any, error) {
	return v.visitBinaryExpr(e)
}

func (e GroupingExpr) accept(v Visitor) (any, error) {
	return v.visitGroupingExpr(e)
}

func (e LiteralExpr) accept(v Visitor) (any, error) {
	return v.visitLiteralExpr(e)
}

func (e UnaryExpr) accept(v Visitor) (any, error) {
	return v.visitUnaryExpr(e)
}

func (e TernaryExpr) accept(v Visitor) (any, error) {
	return v.visitTernaryExpr(e)
}
