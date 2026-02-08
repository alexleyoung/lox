package main

type Expr interface {
	Accept(v Visitor) (any, error)
}

type AssignmentExpr struct {
	Name Token
	Expr Expr
}

type LogicalExpr struct {
	Left, Right Expr
	Op          Token
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

type VariableExpr struct {
	Name Token
}

func NewAssignmentExpr(name Token, expr Expr) AssignmentExpr {
	return AssignmentExpr{Name: name, Expr: expr}
}

func NewLogicalExpr(op Token, left, right Expr) LogicalExpr {
	return LogicalExpr{Op: op, Left: left, Right: right}
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

func NewVariableExpr(name Token) VariableExpr {
	return VariableExpr{Name: name}
}

func (e AssignmentExpr) Accept(v Visitor) (any, error) {
	return v.VisitAssignmentExpr(e)
}

func (e LogicalExpr) Accept(v Visitor) (any, error) {
	return v.VisitLogicalExpr(e)
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

func (e VariableExpr) Accept(v Visitor) (any, error) {
	return v.VisitVariableExpr(e)
}
