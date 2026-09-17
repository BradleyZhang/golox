package lox

type Stmt interface {
	Accept(visitor StmtVisitor) any
}
type StmtVisitor interface {
	VisitPrintStmt(p *PrintStmt) any
	VisitExpression(e *Expression) any
}
type PrintStmt struct {
	Expression Expr
}
type Expression struct {
	Expression Expr
}

func (p PrintStmt) Accept(visitor StmtVisitor) any {
	return visitor.VisitPrintStmt(&p)
}
func (e Expression) Accept(visitor StmtVisitor) any {
	return visitor.VisitExpression(&e)
}
