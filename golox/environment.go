package main

type Environment struct {
	values map[string]any
}

func NewEnvironment() Environment {
	return Environment{values: make(map[string]any)}
}

func (e *Environment) define(name string, value any) {
	// not checking if "name" in `values` allows var redefinition
	// could err if we want to disallow and force assignment instead
	// doing so is not so friendly to REPLs though
	e.values[name] = value
}

func (e Environment) get(name Token) (any, error) {
	val, ok := e.values[name.Lexeme]
	if !ok {
		return nil, NewEnvironmentError(name.Lexeme, "")
	}
	return val, nil
}
