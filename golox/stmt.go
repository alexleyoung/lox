package main

type Stmt interface {
	Accept(v StmtVisitor) error
}

type VariableStmt struct {
	Name        Token
	Initializer Expr
}

type ExpressionStmt struct {
	Expr Expr
}

type PrintStmt struct {
	Expr Expr
}

type BlockStmt struct {
	Statements []Stmt
}

type IfStmt struct {
	Guard      Expr
	ThenBranch Stmt
	ElseBranch Stmt
}

func NewVariableStmt(name Token, initializer Expr) VariableStmt {
	return VariableStmt{Name: name, Initializer: initializer}
}

func NewExpressionStmt(expr Expr) ExpressionStmt {
	return ExpressionStmt{Expr: expr}
}

func NewPrintStmt(expr Expr) PrintStmt {
	return PrintStmt{Expr: expr}
}

func NewBlockStmt(statements []Stmt) BlockStmt {
	return BlockStmt{Statements: statements}
}

func NewIfStmt(guard Expr, thenBranch Stmt, elseBranch Stmt) IfStmt {
	return IfStmt{Guard: guard, ThenBranch: thenBranch, ElseBranch: elseBranch}
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

func (s BlockStmt) Accept(v StmtVisitor) error {
	return v.VisitBlockStmt(s)
}

func (s IfStmt) Accept(v StmtVisitor) error {
	return v.VisitIfStmt(s)
}
