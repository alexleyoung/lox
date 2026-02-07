package main

type Environment struct {
	enclosing *Environment
	values    map[string]any
}

func NewEnvironment() *Environment {
	return &Environment{enclosing: nil, values: make(map[string]any)}
}

func NewNestedEnvironment(enclosing *Environment) *Environment {
	return &Environment{enclosing: enclosing, values: make(map[string]any)}
}

func (e *Environment) define(name string, value any) {
	// not checking if "name" in `values` allows var redefinition
	// could err if we want to disallow and force assignment instead
	// doing so is not so friendly to REPLs though
	e.values[name] = value
}

func (e *Environment) get(name Token) (any, error) {
	val, ok := e.values[name.Lexeme]
	if !ok {
		if e.enclosing != nil {
			return e.enclosing.get(name)
		}
		return nil, NewEnvironmentError(name, "")
	}
	return val, nil
}

func (e *Environment) assign(name Token, value any) error {
	_, ok := e.values[name.Lexeme]
	if !ok {
		if e.enclosing != nil {
			return e.enclosing.assign(name, value)
		}
		return NewEnvironmentError(name, "Undefined variable: '"+name.Lexeme+"'.")
	}

	e.values[name.Lexeme] = value
	return nil
}
