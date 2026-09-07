package lox

import (
	"fmt"
	"strings"
)

// literal        → NUMBER | STRING | "true" | "false" | "nil" ;
// grouping       → "(" expression ")" ;
// unary          → ( "-" | "!" ) expression ;
// binary         → expression operator expression ;
// ternary        → expression operator expression operator expression ;
// operator       → "==" | "!=" | "<" | "<=" | ">" | ">="

type Expr interface {
	String() string
}
type Binary struct {
	left     Expr
	operator Token
	right    Expr
}
type Grouping struct {
	expression Expr
}
type Literal struct {
	value any
}
type Unary struct {
	operator Token
	right    Expr
}
type Ternary struct {
	left      Expr
	operatorL Token
	middle    Expr
	operatorR Token
	right     Expr
}

func (b Binary) String() string {
	return parenthesized(b.operator.Lexeme, b.left, b.right)
}
func (g Grouping) String() string {
	return parenthesized("group", g.expression)
}
func (l Literal) String() string {
	if l.value == nil {
		return "nil"
	}
	return fmt.Sprint(l.value)
}
func (u Unary) String() string {
	return parenthesized(u.operator.Lexeme, u.right)
}
func (t Ternary) String() string {
	return parenthesized(t.operatorL.Lexeme+t.operatorR.Lexeme, t.left, t.middle, t.right)
}

func parenthesized(name string, exprs ...Expr) string {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString(name)
	for _, expr := range exprs {
		builder.WriteString(" ")
		builder.WriteString(expr.String())
	}
	builder.WriteString(")")
	return builder.String()
}
