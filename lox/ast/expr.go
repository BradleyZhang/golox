package ast

import (
	"golox/lox/token"
)

// literal        → NUMBER | STRING | "true" | "false" | "nil" ;
// grouping       → "(" expression ")" ;
// unary          → ( "-" | "!" ) expression ;
// binary         → expression operator expression ;
// ternary        → expression operator expression operator expression ;
// operator       → "==" | "!=" | "<" | "<=" | ">" | ">="

type Expr interface {
	Accept(visitor ExprVisitor) any
}
type ExprVisitor interface {
	VisitAssign(a *Assign) any
	VisitBinary(b *Binary) any
	VisitGrouping(g *Grouping) any
	VisitLiteral(l *Literal) any
	VisitLogical(l *Logical) any
	VisitUnary(u *Unary) any
	VisitTernary(t *Ternary) any
	VisitVariable(v *Variable) any
}
type Assign struct {
	Name  token.Token
	Value Expr
}
type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}
type Grouping struct {
	Expression Expr
}
type Literal struct {
	Value any
}
type Logical struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}
type Unary struct {
	Operator token.Token
	Right    Expr
}
type Ternary struct {
	Left      Expr
	OperatorL token.Token
	Middle    Expr
	OperatorR token.Token
	Right     Expr
}
type Variable struct {
	Name token.Token
}

func (a Assign) Accept(visitor ExprVisitor) any {
	return visitor.VisitAssign(&a)
}
func (b Binary) Accept(visitor ExprVisitor) any {
	return visitor.VisitBinary(&b)
}
func (g Grouping) Accept(visitor ExprVisitor) any {
	return visitor.VisitGrouping(&g)
}
func (l Literal) Accept(visitor ExprVisitor) any {
	return visitor.VisitLiteral(&l)
}
func (l Logical) Accept(visitor ExprVisitor) any {
	return visitor.VisitLogical(&l)
}
func (u Unary) Accept(visitor ExprVisitor) any {
	return visitor.VisitUnary(&u)
}
func (t Ternary) Accept(visitor ExprVisitor) any {
	return visitor.VisitTernary(&t)
}
func (v Variable) Accept(visitor ExprVisitor) any {
	return visitor.VisitVariable(&v)
}
