package lexer

import "github.com/alexleyoung/golox/parser"

type Lexer struct {
	source string
}

func NewLexer(source string) *Lexer {
	return &Lexer{source: source}
}

func (s *Lexer) ScanTokens() []parser.Token {
	return nil
}
