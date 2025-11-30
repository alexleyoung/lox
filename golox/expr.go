package main

type Expr interface {
	Accept(v Visitor) (any, error)
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
	Guard Expr
	Then  Expr
	Else  Expr
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

func NewTernaryExpr(guard, thenBranch, elseBranch Expr) TernaryExpr {
	return TernaryExpr{Guard: guard, Then: thenBranch, Else: elseBranch}
}

func (e BinaryExpr) Accept(v Visitor) (any, error) {
	return v.VisitBinaryExpr(e)
}

func (e GroupingExpr) Accept(v Visitor) (any, error) {
	return v.VisitGroupingExpr(e)
}

func (e LiteralExpr) Accept(v Visitor) (any, error) {
	return v.VisitLiteralExpr(e)
}

func (e UnaryExpr) Accept(v Visitor) (any, error) {
	return v.VisitUnaryExpr(e)
}

func (e TernaryExpr) Accept(v Visitor) (any, error) {
	return v.VisitTernaryExpr(e)
}
