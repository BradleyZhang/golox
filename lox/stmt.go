package lox

type Stmt interface {
	Accept(visitor StmtVisitor) any
}
type StmtVisitor interface {
	VisitPrintStmt(p *PrintStmt) any
	VisitExpression(e *Expression) any
}
type PrintStmt struct {
	expression Expr
}
type Expression struct {
	expression Expr
}

func (p PrintStmt) Accept(visitor StmtVisitor) any {
	return visitor.VisitPrintStmt(&p)
}
func (e Expression) Accept(visitor StmtVisitor) any {
	return visitor.VisitExpression(&e)
}
