package lox

import (
	"fmt"
	"strings"
)

type AstPrinter struct {
}

func (a *AstPrinter) Print(expr Expr) string {
	return expr.Accept(a).(string)
}

func (a *AstPrinter) VisitBinary(b Binary) any {
	return a.parenthesized(b.operator.Lexeme, b.left, b.right)
}
func (a *AstPrinter) VisitGrouping(g Grouping) any {
	return a.parenthesized("group", g.expression)
}
func (a *AstPrinter) VisitLiteral(l Literal) any {
	if l.value == nil {
		return "nil"
	}
	return fmt.Sprint(l.value)
}
func (a *AstPrinter) VisitUnary(u Unary) any {
	return a.parenthesized(u.operator.Lexeme, u.right)
}
func (a *AstPrinter) VisitTernary(t Ternary) any {
	return a.parenthesized(t.operatorL.Lexeme+t.operatorR.Lexeme, t.left, t.middle, t.right)
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
