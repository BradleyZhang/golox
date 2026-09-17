package interpreter

import (
	"fmt"
	"golox/lox/ast"
	"golox/lox/token"
	"strconv"
)

type Interpreter struct {
}

type RuntimeError struct {
	Token   token.Token
	Message string
}

func (e *RuntimeError) Error() string {
	return e.Message
}
func (a *Interpreter) Interpret(statements []ast.Stmt) *RuntimeError {
	for _, stmt := range statements {
		if err := a.excute(stmt); err != nil {
			return err
		}
	}
	return nil
}

type evalResult struct { //evaluate result
	value any
	err   *RuntimeError
}

func (a *Interpreter) excute(stmt ast.Stmt) *RuntimeError {
	if err, ok := stmt.Accept(a).(*RuntimeError); ok {
		return err
	}
	return nil
}

func (a *Interpreter) evaluate(expr ast.Expr) evalResult {
	return expr.Accept(a).(evalResult)
}

// impl StmtVisitor
func (a *Interpreter) VisitExpression(e *ast.Expression) any {
	evalResult := a.evaluate(e.Expression)
	return evalResult.err
}

func (a *Interpreter) VisitPrintStmt(p *ast.PrintStmt) any {
	evalResult := a.evaluate(p.Expression)
	if evalResult.err != nil {
		return evalResult.err
	}
	fmt.Println(stringify(evalResult.value))
	return nil
}

// impl ExprVisitor
func (i *Interpreter) VisitBinary(b *ast.Binary) any {
	left := i.evaluate(b.Left)
	if left.err != nil {
		return evalResult{nil, left.err}
	}
	right := i.evaluate(b.Right)
	if right.err != nil {
		return evalResult{nil, right.err}
	}
	switch b.Operator.Type {
	case token.Minus, token.Slash, token.Star, token.Greater, token.GreaterEqual, token.Less, token.LessEqual:
		l, okL := left.value.(float64)
		r, okR := right.value.(float64)
		if !okL || !okR {
			return evalResult{nil, &RuntimeError{b.Operator, "Operand must be a number."}}
		}
		switch b.Operator.Type {
		case token.Minus:
			return evalResult{l - r, nil}
		case token.Slash:
			if r == 0 {
				return evalResult{nil, &RuntimeError{b.Operator, "Division by zero"}}
			}
			return evalResult{l / r, nil}
		case token.Star:
			return evalResult{l * r, nil}
		case token.Greater:
			return evalResult{l > r, nil}
		case token.GreaterEqual:
			return evalResult{l >= r, nil}
		case token.Less:
			return evalResult{l < r, nil}
		case token.LessEqual:
			return evalResult{l <= r, nil}
		}
	case token.Plus:
		{
			l, okL := left.value.(float64)
			r, okR := right.value.(float64)
			if okL && okR {
				return evalResult{l + r, nil}
			}
		}
		{
			l, okL := left.value.(string)
			r, okR := right.value.(string)
			if okL && okR {
				return evalResult{l + r, nil}
			}
			if okL || okR {
				return evalResult{fmt.Sprint(left.value) + fmt.Sprint(right.value), nil}
			}
		}
		return evalResult{nil, &RuntimeError{b.Operator, "Operands must be two numbers or two strings."}}
	case token.BangEqual:
		return evalResult{!isEqual(left.value, right.value), nil}
	case token.EqualEqual:
		return evalResult{isEqual(left.value, right.value), nil}

	}
	// Unreachable
	return evalResult{nil, &RuntimeError{token.Token{}, "interpreter binary unreachable"}}
}
func (i *Interpreter) VisitGrouping(g *ast.Grouping) any {
	return i.evaluate(g.Expression)
}
func (i *Interpreter) VisitLiteral(l *ast.Literal) any {
	return evalResult{l.Value, nil}
}
func (i *Interpreter) VisitUnary(u *ast.Unary) any {
	right := i.evaluate(u.Right)
	if right.err != nil {
		return evalResult{nil, right.err}
	}
	switch u.Operator.Type {
	case token.Minus:
		r, ok := right.value.(float64)
		if !ok {
			return evalResult{nil, &RuntimeError{u.Operator, "Operand must be a number."}}
		}
		return evalResult{-(r), nil}
	case token.Bang:
		return evalResult{!isTruthy(right.value), nil}
	}
	// Unreachable
	return evalResult{nil, &RuntimeError{token.Token{}, "interpreter unary unreachable"}}
}
func (i *Interpreter) VisitTernary(t *ast.Ternary) any {
	if t.OperatorL.Type == token.QuestionMark && t.OperatorR.Type == token.Colon {
		left := i.evaluate(t.Left)
		if left.err != nil {
			return evalResult{nil, left.err}
		}

		if isTruthy(left.value) {
			return i.evaluate(t.Middle)
		}
		return i.evaluate(t.Right)
	}
	// Unreachable
	return evalResult{nil, &RuntimeError{token.Token{}, "interpreter ternary unreachable"}}
}
func isTruthy(v any) bool {
	if v == nil {
		return false
	}
	if b, ok := v.(bool); ok {
		return b
	}
	return true
}
func isEqual(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return false
	}

	// TODO
	return a == b
}
func stringify(value any) string {
	if value == nil {
		return "nil"
	}
	if v, ok := value.(float64); ok {
		return strconv.FormatFloat(v, 'f', -1, 64)
	}
	if v, ok := value.(bool); ok {
		if v {
			return "true"
		} else {
			return "false"
		}
	}
	if v, ok := value.(string); ok {
		return v
	}
	// TODO literal 都有哪些类型，补完
	return ""
}
