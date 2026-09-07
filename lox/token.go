package lox

import (
	"fmt"
	"strings"
	"text/tabwriter"
)

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

func PrintTokens(tokens []Token) string {
	var b strings.Builder

	w := tabwriter.NewWriter(&b, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "TYPE\tLEXEME\tLITERAL\tLINE")
	fmt.Fprintln(w, "----\t------\t-------\t----")

	for _, t := range tokens {
		fmt.Fprintf(
			w,
			"%v\t%q\t%v\t%d\n",
			t.Type,
			t.Lexeme,
			t.Literal,
			t.Line,
		)
	}

	w.Flush()

	return b.String()
}
