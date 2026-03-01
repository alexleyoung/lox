package main

import "time"

type Callable interface {
	Arity() int
	Call(interpreter *Interpreter, args []any) (any, error)
}

type ClockNativeFn struct{}

func (c ClockNativeFn) Arity() int { return 0 }

func (c ClockNativeFn) Call(interpreter *Interpreter, args []any) (any, error) {
	return float64(time.Now().UnixMilli()) / 1000.0, nil
}

func (c ClockNativeFn) String() string { return "<native fn>" }
