package main

type Visitor interface {
	VisitAssignmentExpr(expr AssignmentExpr) (any, error)
	VisitBinaryExpr(expr BinaryExpr) (any, error)
	VisitGroupingExpr(expr GroupingExpr) (any, error)
	VisitLiteralExpr(expr LiteralExpr) (any, error)
	VisitUnaryExpr(expr UnaryExpr) (any, error)
	VisitVariableExpr(expr VariableExpr) (any, error)
}

type StmtVisitor interface {
	VisitExpressionStmt(stmt ExpressionStmt) error
	VisitPrintStmt(stmt PrintStmt) error
	VisitVariableStmt(stmt VariableStmt) error
	VisitBlockStmt(stmt BlockStmt) error
	VisitIfStmt(stmt IfStmt) error
}
