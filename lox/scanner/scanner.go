package scanner

import (
	"golox/lox/token"
	"strconv"
)

type Scanner struct {
	source  string
	tokens  []token.Token
	start   int
	current int
	line    int
}
type LineError struct {
	Line int
	Msg  string
}

func (e LineError) Error() string {
	return e.Msg
}

var (
	keywords = map[string]token.TokenType{
		"and":    token.And,
		"class":  token.Class,
		"else":   token.Else,
		"false":  token.False,
		"for":    token.For,
		"fun":    token.Fun,
		"if":     token.If,
		"nil":    token.Nil,
		"or":     token.Or,
		"print":  token.Print,
		"return": token.Return,
		"super":  token.Super,
		"this":   token.This,
		"true":   token.True,
		"var":    token.Var,
		"while":  token.While,
	}
)

func NewScanner(source string) *Scanner {
	s := &Scanner{source: source, line: 1}
	return s
}

func (s *Scanner) ScanTokens() ([]token.Token, []*LineError) {
	var tErrors []*LineError
	for !s.isAtEnd() {
		s.start = s.current
		err := s.scanToken()
		if err != nil {
			tErrors = append(tErrors, err)
		}
	}
	s.tokens = append(s.tokens, token.Token{token.EOF, "", nil, s.line})
	if len(tErrors) == 0 {
		return s.tokens, nil
	}
	return s.tokens, tErrors
}

func (s *Scanner) scanToken() *LineError {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(token.LeftParen, nil)
	case ')':
		s.addToken(token.RightParen, nil)
	case '{':
		s.addToken(token.LeftBrace, nil)
	case '}':
		s.addToken(token.RightBrace, nil)
	case ',':
		s.addToken(token.Comma, nil)
	case '.':
		s.addToken(token.Dot, nil)
	case '-':
		s.addToken(token.Minus, nil)
	case '+':
		s.addToken(token.Plus, nil)
	case ';':
		s.addToken(token.Semicolon, nil)
	case '*':
		s.addToken(token.Star, nil)
	case '?':
		s.addToken(token.QuestionMark, nil)
	case ':':
		s.addToken(token.Colon, nil)
	case '!':
		if s.match('=') {
			s.addToken(token.BangEqual, nil)
		} else {
			s.addToken(token.Bang, nil)
		}
	case '=':
		if s.match('=') {
			s.addToken(token.EqualEqual, nil)
		} else {
			s.addToken(token.Equal, nil)
		}
	case '<':
		if s.match('=') {
			s.addToken(token.LessEqual, nil)
		} else {
			s.addToken(token.Less, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(token.Greater, nil)
		} else {
			s.addToken(token.GreaterEqual, nil)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.Slash, nil)
		}
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		s.line++
	// string literal
	case '"':
		err := s.string()
		if err != nil {
			return err
		}
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.identifier()
		} else {
			return &LineError{s.line, "Unexpected character."}
		}
	}
	// Unreachable
	return nil
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	tokenType, ok := keywords[text]
	if !ok {
		tokenType = token.Identifier
		s.addToken(tokenType, text)
	} else {
		s.addToken(tokenType, nil)
	}
}
func (s *Scanner) string() *LineError {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		return &LineError{s.line, "Unterminated string."}
	}
	s.advance()
	value := s.source[s.start+1 : s.current-1]
	s.addToken(token.String, value)
	return nil
}
func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
	}
	f, _ := strconv.ParseFloat(s.source[s.start:s.current], 64)
	s.addToken(token.Number, f)
}
func (s *Scanner) advance() byte {
	s.current++
	return s.source[s.current-1]
}
func (s *Scanner) addToken(tokenType token.TokenType, literal any) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, token.Token{tokenType, text, literal, s.line})
}
func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}
	s.current++
	return true
}
func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}
func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	} else {
		return s.source[s.current+1]
	}
}
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}
func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}
