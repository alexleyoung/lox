package main

type Visitor interface {
	VisitAssignmentExpr(expr AssignmentExpr) (any, error)
	VisitLogicalExpr(expr LogicalExpr) (any, error)
	VisitBinaryExpr(expr BinaryExpr) (any, error)
	VisitGroupingExpr(expr GroupingExpr) (any, error)
	VisitLiteralExpr(expr LiteralExpr) (any, error)
	VisitUnaryExpr(expr UnaryExpr) (any, error)
	VisitVariableExpr(expr VariableExpr) (any, error)
}

type StmtVisitor interface {
	VisitVariableStmt(stmt VariableStmt) error
	VisitExpressionStmt(stmt ExpressionStmt) error
	VisitPrintStmt(stmt PrintStmt) error
	VisitIfStmt(stmt IfStmt) error
	VisitWhileStmt(stmt WhileStmt) error
	VisitBlockStmt(stmt BlockStmt) error
}
