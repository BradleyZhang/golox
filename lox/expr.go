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
	VisitBinary(b *Binary) any
	VisitGrouping(g *Grouping) any
	VisitLiteral(l *Literal) any
	VisitUnary(u *Unary) any
	VisitTernary(t *Ternary) any
}
type Binary struct {
	Left     Expr
	Operator Token
	Right    Expr
}
type Grouping struct {
	Expression Expr
}
type Literal struct {
	Value any
}
type Unary struct {
	Operator Token
	Right    Expr
}
type Ternary struct {
	Left      Expr
	OperatorL Token
	Middle    Expr
	OperatorR Token
	Right     Expr
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
func (u Unary) Accept(visitor ExprVisitor) any {
	return visitor.VisitUnary(&u)
}
func (t Ternary) Accept(visitor ExprVisitor) any {
	return visitor.VisitTernary(&t)
}
