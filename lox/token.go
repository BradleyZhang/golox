package lox

import "fmt"

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any //TODO 看换成什么适合，现在看字面量就是字符串或float
	Line    int
}

func (t *Token) Init(tokenType TokenType, lexeme string, literal any, line int) {
	t.Type = tokenType
	t.Lexeme = lexeme
	t.Literal = literal
	t.Line = line
}
func (t *Token) ToString() string {
	return fmt.Sprintf("%v %s %v %v", t.Type, t.Lexeme, t.Literal, t.Line)
}
