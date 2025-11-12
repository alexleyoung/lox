package lexer

type Lexer struct {
	source string
}

type Token struct{}

func NewLexer(source string) *Lexer {
	return &Lexer{source: source}
}

func (s *Lexer) ScanTokens() []Token {
	return nil
}
