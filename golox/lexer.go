package main

type Lexer struct {
	source string
	tokens []Token

	start, current, line int
}

func NewLexer(source string) *Lexer {
	return &Lexer{source: source}
}

func (s *Lexer) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	return s.tokens
}

func (s *Lexer) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Lexer) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LEFT_PAREN, nil)
	case ')':
		s.addToken(RIGHT_PAREN, nil)
	case '{':
		s.addToken(LEFT_BRACE, nil)
	case '}':
		s.addToken(RIGHT_BRACE, nil)
	case ',':
		s.addToken(COMMA, nil)
	case '.':
		s.addToken(DOT, nil)
	case '-':
		s.addToken(MINUS, nil)
	case '+':
		s.addToken(PLUS, nil)
	case ';':
		s.addToken(SEMICOLON, nil)
	case '*':
		s.addToken(STAR, nil)

	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL, nil)
		} else {
			s.addToken(BANG, nil)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL, nil)
		} else {
			s.addToken(EQUAL, nil)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL, nil)
		} else {
			s.addToken(LESS, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL, nil)
		} else {
			s.addToken(GREATER, nil)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH, nil)
		}

	case ' ':
		fallthrough
	case '\t':
		fallthrough
	case '\r':

	case '\n':
		s.line++

	default:
		Error(s.line, "Unexpected character: '"+string(c)+"'")
	}
}

func (s *Lexer) advance() byte {
	curr := s.source[s.current]
	s.current++
	return curr
}

func (s *Lexer) addToken(tok TokenType, literal *Object) {
	text := string(s.source[s.start:s.current])
	s.tokens = append(s.tokens, Token{Type: tok, Lexeme: string(text), Literal: literal, Line: s.line})
}

func (s *Lexer) match(expected byte) bool {
	if s.isAtEnd() || s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *Lexer) peek() byte {
	if s.isAtEnd() {
		return 0
	}

	return s.source[s.current]
}
