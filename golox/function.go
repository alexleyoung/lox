package main

import "fmt"

type Function struct {
	declaration FunctionStmt
}

func NewFunction(declaration FunctionStmt) Function {
	return Function{declaration}
}

func (f *Function) Arity() int {
	return len(f.declaration.Params)
}

func (f *Function) Call(i *Interpreter, args []any) (any, error) {
	env := NewNestedEnvironment(i.Globals)
	for i, param := range f.declaration.Params {
		env.define(param.Lexeme, args[i])
	}

	i.executeBlock(f.declaration.Body, env)
	return nil, nil
}

func (f *Function) String() string {
	return fmt.Sprintf("<fn %s>", f.declaration.Name)
}
