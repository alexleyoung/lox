package main

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

func (i *Interpreter) isTruthy(v any) bool {
	switch v.(type) {
	case bool:
		return v.(bool)
	case int:
		return v.(int) != 0
	case float64:
		return v.(float64) != 0
	case string:
		return v.(string) != ""
	case nil:
		return false
	default:
		return true
	}
}

func (i *Interpreter) VisitGroupingExpr(expr GroupingExpr) (any, error) {
	return i.evaluate(expr.Expr)
}

func (i *Interpreter) evaluate(expr Expr) (any, error) {
	return expr.Accept(i)
}

func (i *Interpreter) VisitBinaryExpr(expr BinaryExpr) (any, error) {

}
