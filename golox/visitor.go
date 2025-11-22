package main

// Visitor is an interface for the visitor pattern.
type Visitor interface {
	visitBinaryExpr(expr BinaryExpr) (any, error)
	visitGroupingExpr(expr GroupingExpr) (any, error)
	visitLiteralExpr(expr LiteralExpr) (any, error)
	visitUnaryExpr(expr UnaryExpr) (any, error)
	visitTernaryExpr(expr TernaryExpr) (any, error)
}
