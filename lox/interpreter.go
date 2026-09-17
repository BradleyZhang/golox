package lox

import (
	"fmt"
	"strconv"
)

type Interpreter struct {
}

type RuntimeError struct {
	Token   Token
	Message string
}

func (e *RuntimeError) Error() string {
	return e.Message
}
func (a *Interpreter) Interpret(statements []Stmt) {
	for _, stmt := range statements {
		if err := a.excute(stmt); err != nil {
			GlobalLox.RuntimeError(err)
		}
	}

}

type evalResult struct { //evaluate result
	value any
	err   *RuntimeError
}

func (a *Interpreter) excute(stmt Stmt) *RuntimeError {
	if err, ok := stmt.Accept(a).(*RuntimeError); ok {
		return err
	}
	return nil
}

func (a *Interpreter) evaluate(expr Expr) evalResult {
	return expr.Accept(a).(evalResult)
}

// impl StmtVisitor
func (a *Interpreter) VisitExpression(e *Expression) any {
	evalResult := a.evaluate(e.Expression)
	return evalResult.err
}

func (a *Interpreter) VisitPrintStmt(p *PrintStmt) any {
	evalResult := a.evaluate(p.Expression)
	if evalResult.err != nil {
		return evalResult.err
	}
	fmt.Println(stringify(evalResult.value))
	return nil
}

// impl ExprVisitor
func (i *Interpreter) VisitBinary(b *Binary) any {
	left := i.evaluate(b.Left)
	if left.err != nil {
		return evalResult{nil, left.err}
	}
	right := i.evaluate(b.Right)
	if right.err != nil {
		return evalResult{nil, right.err}
	}
	switch b.Operator.Type {
	case Minus, Slash, Star, Greater, GreaterEqual, Less, LessEqual:
		l, okL := left.value.(float64)
		r, okR := right.value.(float64)
		if !okL || !okR {
			return evalResult{nil, &RuntimeError{b.Operator, "Operand must be a number."}}
		}
		switch b.Operator.Type {
		case Minus:
			return evalResult{l - r, nil}
		case Slash:
			if r == 0 {
				return evalResult{nil, &RuntimeError{b.Operator, "Division by zero"}}
			}
			return evalResult{l / r, nil}
		case Star:
			return evalResult{l * r, nil}
		case Greater:
			return evalResult{l > r, nil}
		case GreaterEqual:
			return evalResult{l >= r, nil}
		case Less:
			return evalResult{l < r, nil}
		case LessEqual:
			return evalResult{l <= r, nil}
		}
	case Plus:
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
	case BangEqual:
		return evalResult{!isEqual(left.value, right.value), nil}
	case EqualEqual:
		return evalResult{isEqual(left.value, right.value), nil}

	}
	// Unreachable
	return evalResult{nil, &RuntimeError{Token{}, "interpreter binary unreachable"}}
}
func (i *Interpreter) VisitGrouping(g *Grouping) any {
	return i.evaluate(g.Expression)
}
func (i *Interpreter) VisitLiteral(l *Literal) any {
	return evalResult{l.Value, nil}
}
func (i *Interpreter) VisitUnary(u *Unary) any {
	right := i.evaluate(u.Right)
	if right.err != nil {
		return evalResult{nil, right.err}
	}
	switch u.Operator.Type {
	case Minus:
		r, ok := right.value.(float64)
		if !ok {
			return evalResult{nil, &RuntimeError{u.Operator, "Operand must be a number."}}
		}
		return evalResult{-(r), nil}
	case Bang:
		return evalResult{!isTruthy(right.value), nil}
	}
	// Unreachable
	return evalResult{nil, &RuntimeError{Token{}, "interpreter unary unreachable"}}
}
func (i *Interpreter) VisitTernary(t *Ternary) any {
	if t.OperatorL.Type == QuestionMark && t.OperatorR.Type == Colon {
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
	return evalResult{nil, &RuntimeError{Token{}, "interpreter ternary unreachable"}}
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
