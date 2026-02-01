package main

type Stmt interface {
	Accept(v StmtVisitor) (any, error)
}

type ExpressionStmt struct {
	Expr Expr
}

type PrintStmt struct {
	Expr Expr
}

func NewExpressionStmt(expr Expr) ExpressionStmt {
	return ExpressionStmt{Expr: expr}
}

func NewPrintStmt(expr Expr) PrintStmt {
	return PrintStmt{Expr: expr}
}

func (s ExpressionStmt) Accept(v StmtVisitor) (any, error) {
	return v.VisitExpressionStmt(s)
}

func (s PrintStmt) Accept(v StmtVisitor) (any, error) {
	return v.VisitPrintStmt(s)
}
