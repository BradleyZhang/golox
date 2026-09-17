package printer

import (
	"fmt"
	"golox/lox/ast"
	"strings"
)

type AstPrinter struct {
}

func (a *AstPrinter) Print(expr ast.Expr) string {
	return expr.Accept(a).(string)
}

func (a *AstPrinter) VisitBinary(b *ast.Binary) any {
	return a.parenthesized(b.Operator.Lexeme, b.Left, b.Right)
}
func (a *AstPrinter) VisitGrouping(g *ast.Grouping) any {
	return a.parenthesized("group", g.Expression)
}
func (a *AstPrinter) VisitLiteral(l *ast.Literal) any {
	if l.Value == nil {
		return "nil"
	}
	return fmt.Sprint(l.Value)
}
func (a *AstPrinter) VisitUnary(u *ast.Unary) any {
	return a.parenthesized(u.Operator.Lexeme, u.Right)
}
func (a *AstPrinter) VisitTernary(t *ast.Ternary) any {
	return a.parenthesized(t.OperatorL.Lexeme+t.OperatorR.Lexeme, t.Left, t.Middle, t.Right)
}

func (a *AstPrinter) parenthesized(name string, exprs ...ast.Expr) string {
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
