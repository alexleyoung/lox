package main

import (
	"fmt"
)

type Interpreter struct {
	environment *Environment
	reporter    *ErrorReporter

	repl bool
}

func NewInterpreter() *Interpreter {
	return &Interpreter{environment: NewEnvironment(), reporter: NewErrorReporter()}
}

func (i *Interpreter) Interpret(statements []Stmt) error {
	for _, statement := range statements {
		err := i.execute(statement)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Interpreter) VisitVariableStmt(stmt VariableStmt) error {
	var value any = nil
	var err error

	if stmt.Initializer != nil {
		value, err = i.evaluate(stmt.Initializer)
		if err != nil {
			return err
		}
	}

	i.environment.define(stmt.Name.Lexeme, value)
	return nil
}

func (i *Interpreter) VisitPrintStmt(stmt PrintStmt) error {
	value, err := i.evaluate(stmt.Expr)
	if err != nil {
		return err
	}
	fmt.Print(i.stringify(value))
	return nil
}

func (i *Interpreter) VisitExpressionStmt(stmt ExpressionStmt) error {
	val, err := i.evaluate(stmt.Expr)
	if err != nil {
		return err
	}

	if i.repl {
		fmt.Print(val)
	}

	return nil
}

func (i *Interpreter) VisitBlockStmt(stmt BlockStmt) error {
	return i.executeBlock(stmt.statements, NewNestedEnvironment(i.environment))
}

func (i *Interpreter) executeBlock(stmts []Stmt, env *Environment) error {
	previous := i.environment
	i.environment = env
	for _, stmt := range stmts {
		err := i.execute(stmt)
		if err != nil {
			return err
		}
	}
	i.environment = previous
	return nil
}

func (i *Interpreter) VisitAssignmentExpr(expr AssignmentExpr) (any, error) {
	value, err := i.evaluate(expr.Expr)
	if err != nil {
		return nil, err
	}
	i.environment.assign(expr.Name, value)
	return value, nil
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
		leftNum, leftOk := left.(float64)
		rightNum, rightOk := right.(float64)
		if !leftOk || !rightOk {
			return nil, NewRuntimeError(expr.Op, "Operands must be numbers.")
		}
		if rightNum == 0 {
			return nil, NewRuntimeError(expr.Op, "Division by zero.")
		}
		return leftNum / rightNum, nil
	case MINUS:
		leftNum, leftOk := left.(float64)
		rightNum, rightOk := right.(float64)
		if !leftOk || !rightOk {
			return nil, NewRuntimeError(expr.Op, "Operands must be numbers.")
		}
		return leftNum - rightNum, nil
	case STAR:
		leftNum, leftOk := left.(float64)
		rightNum, rightOk := right.(float64)
		if !leftOk || !rightOk {
			return nil, NewRuntimeError(expr.Op, "Operands must be numbers.")
		}
		return leftNum * rightNum, nil
	// use + for string concat and number addition
	case PLUS:
		switch left := left.(type) {
		case float64:
			switch right := right.(type) {
			case string:
				return i.stringify(left) + right, nil
			case float64:
				return left + right, nil
			}
		case string:
			return left + i.stringify(right), nil
		}
		return nil, NewRuntimeError(expr.Op, "Operands must be two numbers or two strings.")

	// only supported between numbers
	case LESS:
		leftNum, leftOk := left.(float64)
		rightNum, rightOk := right.(float64)
		if !leftOk || !rightOk {
			return nil, NewRuntimeError(expr.Op, "Operands must be numbers.")
		}
		return leftNum < rightNum, nil
	case LESS_EQUAL:
		leftNum, leftOk := left.(float64)
		rightNum, rightOk := right.(float64)
		if !leftOk || !rightOk {
			return nil, NewRuntimeError(expr.Op, "Operands must be numbers.")
		}
		return leftNum <= rightNum, nil
	case GREATER:
		leftNum, leftOk := left.(float64)
		rightNum, rightOk := right.(float64)
		if !leftOk || !rightOk {
			return nil, NewRuntimeError(expr.Op, "Operands must be numbers.")
		}
		return leftNum > rightNum, nil
	case GREATER_EQUAL:
		leftNum, leftOk := left.(float64)
		rightNum, rightOk := right.(float64)
		if !leftOk || !rightOk {
			return nil, NewRuntimeError(expr.Op, "Operands must be numbers.")
		}
		return leftNum >= rightNum, nil

	case EQUAL_EQUAL:
		return left == right, nil
	case BANG_EQUAL:
		return !(left == right), nil
	}

	return nil, nil
}

func (i *Interpreter) VisitGroupingExpr(expr GroupingExpr) (any, error) {
	return i.evaluate(expr.Expr)
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
		if num, ok := v.(float64); ok {
			return -num, nil
		}
		return nil, NewRuntimeError(expr.Op, "Operand must be a number.")
	}

	return nil, nil
}

func (i *Interpreter) VisitLiteralExpr(expr LiteralExpr) (any, error) {
	return expr.Value, nil
}

func (i *Interpreter) VisitVariableExpr(expr VariableExpr) (any, error) {
	value, err := i.environment.get(expr.Name)
	if err != nil {
		return nil, NewRuntimeError(expr.Name, err.Error())
	}
	return value, nil
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

func (i *Interpreter) execute(stmt Stmt) error {
	return stmt.Accept(i)
}
