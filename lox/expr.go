package lox

// literal        → NUMBER | STRING | "true" | "false" | "nil" ;
// grouping       → "(" expression ")" ;
// unary          → ( "-" | "!" ) expression ;
// binary         → expression operator expression ;
// ternary        → expression operator expression operator expression ;
// operator       → "==" | "!=" | "<" | "<=" | ">" | ">="

type Expr interface {
	exprNode()
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

func (b Binary) exprNode() {
}
func (g Grouping) exprNode() {
}
func (l Literal) exprNode() {
}
func (u Unary) exprNode() {
}
func (t Ternary) exprNode() {
}
