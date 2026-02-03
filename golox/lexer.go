package main

import "strconv"

var keywords = map[string]TokenType{
	"and":    AND,
	"or":     OR,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"for":    FOR,
	"while":  WHILE,
	"nil":    NIL,
	"print":  PRINT,
	"return": RETURN,
	"fun":    FUN,
	"class":  CLASS,
	"var":    VAR,
	"super":  SUPER,
	"this":   THIS,
}

type Lexer struct {
	source string
	tokens []Token

	start, current, line int
	errors               []error
	reporter             *ErrorReporter
}

func NewLexer(source string) *Lexer {
	return &Lexer{source: source, reporter: NewErrorReporter(), errors: make([]error, 0), line: 1}
}

func (s *Lexer) ScanTokens() ([]Token, []error) {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, Token{Type: EOF, Lexeme: "", Literal: nil, Line: s.line})
	return s.tokens, s.errors
}

func (s *Lexer) scanToken() {
	c := s.advance()
	switch c {
	// single character tokens
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
	case '?':
		s.addToken(QUESTION, nil)
	case ':':
		s.addToken(COLON, nil)

	// two character tokens
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

	// string literals
	case '"':
		s.scanString()

	// whitespace
	case ' ', '\t', '\r':

	case '\n':
		s.line++

	default:
		// number literals
		if isDigit(c) {
			s.scanNumber()
		} else if isAlpha(c) {
			s.scanIdentifier()
		} else {
			err := NewLexerError(s.line, "Unexpected character: '"+string(c)+"'")
			s.reporter.Report(err)
			s.errors = append(s.errors, err)
		}
	}
}

func (s *Lexer) scanIdentifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := string(s.source[s.start:s.current])
	keyword, ok := keywords[text]
	if !ok {
		s.addToken(IDENTIFIER, nil)
		return
	}
	s.addToken(keyword, nil)
}

func (s *Lexer) scanString() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		err := NewLexerError(s.line, "Unterminated string")
		s.reporter.Report(err)
		s.errors = append(s.errors, err)
		return
	}

	// consume closing "
	s.advance()

	// trim quotes
	value := string(s.source[s.start+1 : s.current-1])
	s.addToken(STRING, value)
}

func (s *Lexer) scanNumber() {
	for isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
	}

	value, err := strconv.ParseFloat((s.source[s.start:s.current]), 10)
	if err != nil {
		err2 := NewLexerError(s.line, "Invalid number: "+s.source[s.start:s.current])
		s.reporter.Report(err2)
		s.errors = append(s.errors, err2)
		return
	}
	s.addToken(NUMBER, value)
}

// create and add token, start to current, to tokens
func (s *Lexer) addToken(tok TokenType, literal any) {
	text := string(s.source[s.start:s.current])
	s.tokens = append(s.tokens, Token{Type: tok, Lexeme: text, Literal: literal, Line: s.line})
}

// consumes character iff current matches expected
func (s *Lexer) match(expected byte) bool {
	if s.isAtEnd() || s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

// peek but dont consume current character
func (s *Lexer) peek() byte {
	if s.isAtEnd() {
		return 0
	}

	return s.source[s.current]
}

// peek next character
func (s *Lexer) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

// consume and advance one character
func (s *Lexer) advance() byte {
	curr := s.source[s.current]
	s.current++
	return curr
}

func (s *Lexer) isAtEnd() bool {
	return s.current >= len(s.source)
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
