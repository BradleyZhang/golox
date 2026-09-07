package lox

import (
	"fmt"
	"strings"
)

type AstPrinter struct {
}

func (a *AstPrinter) Print(expr Expr) string {
	switch e := expr.(type) {

	case Binary:
		return a.parenthesized(e.operator.Lexeme, e.left, e.right)
	case Grouping:
		return a.parenthesized("group", e.expression)
	case Literal:
		if e.value == nil {
			return "nil"
		}
		return fmt.Sprint(e.value)
	case Unary:
		return a.parenthesized(e.operator.Lexeme, e.right)
	case Ternary:
		return a.parenthesized(e.operatorL.Lexeme+e.operatorR.Lexeme, e.left, e.middle, e.right)
	}
	return ""
}
func (a *AstPrinter) parenthesized(name string, exprs ...Expr) string {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString(name)
	for _, expr := range exprs {
		builder.WriteString(" ")
		builder.WriteString(a.Print(expr))
	}
	builder.WriteString(")")
	return builder.String()
}
