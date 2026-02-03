package main

type Stmt interface {
	Accept(v StmtVisitor) error
}

type ExpressionStmt struct {
	Expr Expr
}

type PrintStmt struct {
	Expr Expr
}

type VariableStmt struct {
	Name        Token
	Initializer Expr
}

func NewExpressionStmt(expr Expr) ExpressionStmt {
	return ExpressionStmt{Expr: expr}
}

func NewPrintStmt(expr Expr) PrintStmt {
	return PrintStmt{Expr: expr}
}

func NewVariableStmt(name Token, initializer Expr) VariableStmt {
	return VariableStmt{Name: name, Initializer: initializer}
}

func (s ExpressionStmt) Accept(v StmtVisitor) error {
	return v.VisitExpressionStmt(s)
}

func (s PrintStmt) Accept(v StmtVisitor) error {
	return v.VisitPrintStmt(s)
}

func (s VariableStmt) Accept(v StmtVisitor) error {
	return v.VisitVariableStmt(s)
}
