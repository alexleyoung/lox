package main

import (
	"fmt"
)

type Interpreter struct{}

func (i *Interpreter) VisitLiteralExpr(expr LiteralExpr) (any, error) {
	return expr.Value, nil
}

func (i *Interpreter) VisitUnaryExpr(expr UnaryExpr) (any, error) {
	v, err := i.evaluate(expr.Expr)
	if err != nil {
		return nil, err
	}

	switch expr.Op.Type {
	case BANG:
		return !i.isTruthy(v), nil
	case MINUS:
		return -v.(float64), nil
	}

	return nil, nil
}

func (i *Interpreter) VisitGroupingExpr(expr GroupingExpr) (any, error) {
	return i.evaluate(expr.Expr)
}

func (i *Interpreter) VisitBinaryExpr(expr BinaryExpr) (any, error) {
	left, err := i.evaluate(expr.Left)
	if err != nil {
		return nil, err
	}
	right, err := i.evaluate(expr.Right)
	if err != nil {
		return nil, err
	}

	switch expr.Op.Type {
	case SLASH:
		return left.(float64) / right.(float64), nil
	case MINUS:
		return left.(float64) - right.(float64), nil
	case STAR:
		return left.(float64) * right.(float64), nil
	// use + for string concat and number addition
	case PLUS:
		switch left.(type) {
		case string:
			return left.(string) + right.(string), nil
		case float64:
			return left.(float64) + right.(float64), nil
		}

	// only supported between numbers
	case LESS:
		return left.(float64) < right.(float64), nil
	case LESS_EQUAL:
		return left.(float64) <= right.(float64), nil
	case GREATER:
		return left.(float64) > right.(float64), nil
	case GREATER_EQUAL:
		return left.(float64) >= right.(float64), nil

	case EQUAL_EQUAL:
		return left == right, nil
	case BANG_EQUAL:
		return !(left == right), nil
	}

	return nil, nil
}

func (i *Interpreter) VisitTernaryExpr(expr TernaryExpr) (any, error) {
	condition, err := i.evaluate(expr.Guard)
	if err != nil {
		return nil, err
	}
	thenBranch, err := i.evaluate(expr.Then)
	if err != nil {
		return nil, err
	}
	elseBranch, err := i.evaluate(expr.Else)
	if err != nil {
		return nil, err
	}

	if i.isTruthy(condition) {
		return thenBranch, nil
	}
	return elseBranch, nil
}

func (i *Interpreter) isTruthy(v any) bool {
	switch v := v.(type) {
	case bool:
		return v
	case int:
		return v != 0
	case float64:
		return v != 0
	case string:
		return v != ""
	case nil:
		return false
	default:
		return true
	}

}

func (i *Interpreter) evaluate(expr Expr) (any, error) {
	return expr.Accept(i)
}

func (i *Interpreter) stringify(obj any) string {
	if obj == nil {
		return "nil"
	}

	if num, ok := obj.(float64); ok {
		text := fmt.Sprintf("%v", num)
		return text
	}

	return fmt.Sprintf("%v", obj)
}

func (i *Interpreter) Interpret(expr Expr) (any, error) {
	value, err := i.evaluate(expr)
	if err != nil {
		return nil, err
	}
	print(i.stringify(value))
	return value, nil
}
