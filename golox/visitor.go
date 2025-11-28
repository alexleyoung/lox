package main

// Visitor is an interface for the visitor pattern.
type Visitor interface {
	VisitBinaryExpr(expr BinaryExpr) (any, error)
	VisitGroupingExpr(expr GroupingExpr) (any, error)
	VisitLiteralExpr(expr LiteralExpr) (any, error)
	VisitUnaryExpr(expr UnaryExpr) (any, error)
	VisitTernaryExpr(expr TernaryExpr) (any, error)
}
