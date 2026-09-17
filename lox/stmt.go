package lox

type Stmt interface {
	Accept(visitor StmtVisitor) any
}
type StmtVisitor interface {
	visitPrintStmt(p *PrintStmt) any
	visitExpression(e *Expression) any
}
type PrintStmt struct {
	expression Expr
}
type Expression struct {
	expression Expr
}

func (p PrintStmt) Accept(visitor StmtVisitor) any {
	return visitor.visitPrintStmt(&p)
}
func (e Expression) Accept(visitor StmtVisitor) any {
	return visitor.visitExpression(&e)
}
