package main

import "fmt"

type LoxError interface {
	error
	Line() int
}

type LexerError struct {
	line    int
	Message string
}

func (e LexerError) Error() string {
	return fmt.Sprintf("[line %d] Error: %s", e.line, e.Message)
}

func (e LexerError) Line() int {
	return e.line
}

type ParserError struct {
	Token   Token
	Message string
}

func (e ParserError) Error() string {
	where := " at 'token'"
	if e.Token.Type == EOF {
		where = " at end"
	} else {
		where = fmt.Sprintf(" at '%s'", e.Token.Lexeme)
	}
	return fmt.Sprintf("[line %d] Error%s: %s", e.Token.Line, where, e.Message)
}

func (e ParserError) Line() int {
	return e.Token.Line
}

type RuntimeError struct {
	Token   Token
	Message string
}

func (e RuntimeError) Error() string {
	return fmt.Sprintf("[line %d] RuntimeError: %s", e.Token.Line, e.Message)
}

func (e RuntimeError) Line() int {
	return e.Token.Line
}

type EnvironmentError struct {
	Name    string
	Message string
}

func (e EnvironmentError) Error() string {
	if e.Message == "" {
		return fmt.Sprintf("Undefined variable '%s'.", e.Name)
	}
	return fmt.Sprintf("Undefined variable '%s'. %s", e.Name, e.Message)
}

func (e EnvironmentError) Line() int {
	return -1
}

func NewLexerError(line int, message string) LexerError {
	return LexerError{line: line, Message: message}
}

func NewParserError(token Token, message string) ParserError {
	return ParserError{Token: token, Message: message}
}

func NewRuntimeError(token Token, message string) RuntimeError {
	return RuntimeError{Token: token, Message: message}
}

func NewEnvironmentError(name string, message string) EnvironmentError {
	return EnvironmentError{Name: name, Message: message}
}

type ErrorReporter struct {
	hadError bool
}

func NewErrorReporter() *ErrorReporter {
	return &ErrorReporter{hadError: false}
}

func (r *ErrorReporter) Report(err error) {
	if loxErr, ok := err.(LoxError); ok {
		fmt.Println(loxErr.Error())
		r.hadError = true
	} else {
		fmt.Println(err.Error())
		r.hadError = true
	}
}

func (r *ErrorReporter) HadError() bool {
	return r.hadError
}

func (r *ErrorReporter) Reset() {
	r.hadError = false
}
