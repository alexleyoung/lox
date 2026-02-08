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

type IfStmt struct {
	Guard      Expr
	ThenBranch Stmt
	ElseBranch Stmt
}

type PrintStmt struct {
	Expr Expr
}

type WhileStmt struct {
	Condition Expr
	Body      Stmt
}

type BlockStmt struct {
	Statements []Stmt
}

func NewVariableStmt(name Token, initializer Expr) VariableStmt {
	return VariableStmt{Name: name, Initializer: initializer}
}

func NewExpressionStmt(expr Expr) ExpressionStmt {
	return ExpressionStmt{Expr: expr}
}

func NewIfStmt(guard Expr, thenBranch Stmt, elseBranch Stmt) IfStmt {
	return IfStmt{Guard: guard, ThenBranch: thenBranch, ElseBranch: elseBranch}
}

func NewPrintStmt(expr Expr) PrintStmt {
	return PrintStmt{Expr: expr}
}

func NewWhileStmt(condition Expr, body Stmt) WhileStmt {
	return WhileStmt{Condition: condition, Body: body}
}

func NewBlockStmt(statements []Stmt) BlockStmt {
	return BlockStmt{Statements: statements}
}

func (s VariableStmt) Accept(v StmtVisitor) error {
	return v.VisitVariableStmt(s)
}

func (s ExpressionStmt) Accept(v StmtVisitor) error {
	return v.VisitExpressionStmt(s)
}

func (s IfStmt) Accept(v StmtVisitor) error {
	return v.VisitIfStmt(s)
}

func (s PrintStmt) Accept(v StmtVisitor) error {
	return v.VisitPrintStmt(s)
}

func (s WhileStmt) Accept(v StmtVisitor) error {
	return v.VisitWhileStmt(s)
}

func (s BlockStmt) Accept(v StmtVisitor) error {
	return v.VisitBlockStmt(s)
}
