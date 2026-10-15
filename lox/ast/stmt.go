package ast

import (
	"golox/lox/token"
)

type Stmt interface {
	Accept(visitor StmtVisitor) any
}
type StmtVisitor interface {
	VisitPrintStmt(p *PrintStmt) any
	VisitBlock(b *Block) any
	VisitExpression(e *Expression) any
	VisitVarStmt(v *VarStmt) any
}
type PrintStmt struct {
	Expression Expr
}
type Block struct {
	Statements []Stmt
}
type Expression struct {
	Expression Expr
}
type VarStmt struct {
	Name        token.Token
	Initializer Expr
}

func (p PrintStmt) Accept(visitor StmtVisitor) any {
	return visitor.VisitPrintStmt(&p)
}
func (b Block) Accept(visitor StmtVisitor) any {
	return visitor.VisitBlock(&b)
}
func (e Expression) Accept(visitor StmtVisitor) any {
	return visitor.VisitExpression(&e)
}
func (v VarStmt) Accept(visitor StmtVisitor) any {
	return visitor.VisitVarStmt(&v)
}
