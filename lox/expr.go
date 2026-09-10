package lox

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
	visitBinary(b *Binary) any
	visitGrouping(g *Grouping) any
	visitLiteral(l *Literal) any
	visitUnary(u *Unary) any
	visitTernary(t *Ternary) any
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

func (b Binary) Accept(visitor ExprVisitor) any {
	return visitor.visitBinary(&b)
}
func (g Grouping) Accept(visitor ExprVisitor) any {
	return visitor.visitGrouping(&g)
}
func (l Literal) Accept(visitor ExprVisitor) any {
	return visitor.visitLiteral(&l)
}
func (u Unary) Accept(visitor ExprVisitor) any {
	return visitor.visitUnary(&u)
}
func (t Ternary) Accept(visitor ExprVisitor) any {
	return visitor.visitTernary(&t)
}
