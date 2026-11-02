package ast

import (
	"golox/lox/token"
)

type Stmt interface {
	Accept(visitor StmtVisitor) any
}
type StmtVisitor interface {
	VisitIfStmt(i *IfStmt) any
	VisitPrintStmt(p *PrintStmt) any
	VisitBlock(b *Block) any
	VisitExpression(e *Expression) any
	VisitVarStmt(v *VarStmt) any
	VisitWhileStmt(w *WhileStmt) any
}
type IfStmt struct {
	Condition  Expr
	ThenBranch Stmt
	ElseBranch Stmt
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
type WhileStmt struct {
	Condition Expr
	Body      Stmt
}

func (i IfStmt) Accept(visitor StmtVisitor) any {
	return visitor.VisitIfStmt(&i)
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
func (w WhileStmt) Accept(visitor StmtVisitor) any {
	return visitor.VisitWhileStmt(&w)
}
