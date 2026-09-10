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
func (a *Interpreter) Interpret(expr Expr) {
	result := a.evaluate(expr)
	if result.err != nil {
		GlobalLox.RuntimeError(result.err)
		return
	}
	fmt.Println(stringify(result.value))
}

type evalResult struct { //evaluate result
	value any
	err   *RuntimeError
}

func (a *Interpreter) evaluate(expr Expr) evalResult {
	return expr.Accept(a).(evalResult)
}

func (i *Interpreter) visitBinary(b *Binary) any {
	left := i.evaluate(b.left)
	if left.err != nil {
		return evalResult{nil, left.err}
	}
	right := i.evaluate(b.right)
	if right.err != nil {
		return evalResult{nil, right.err}
	}
	switch b.operator.Type {
	case Minus, Slash, Star, Greater, GreaterEqual, Less, LessEqual:
		l, okL := left.value.(float64)
		r, okR := right.value.(float64)
		if !okL || !okR {
			return evalResult{nil, &RuntimeError{b.operator, "Operand must be a number."}}
		}
		switch b.operator.Type {
		case Minus:
			return evalResult{l - r, nil}
		case Slash:
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
		}
		return evalResult{nil, &RuntimeError{b.operator, "Operands must be two numbers or two strings."}}
	case BangEqual:
		return evalResult{!isEqual(left.value, right.value), nil}
	case EqualEqual:
		return evalResult{isEqual(left.value, right.value), nil}

	}
	// Unreachable
	return evalResult{nil, &RuntimeError{Token{}, "interpreter binary unreachable"}}
}
func (i *Interpreter) visitGrouping(g *Grouping) any {
	return i.evaluate(g.expression)
}
func (i *Interpreter) visitLiteral(l *Literal) any {
	return evalResult{l.value, nil}
}
func (i *Interpreter) visitUnary(u *Unary) any {
	right := i.evaluate(u.right)
	if right.err != nil {
		return evalResult{nil, right.err}
	}
	switch u.operator.Type {
	case Minus:
		r, ok := right.value.(float64)
		if !ok {
			return evalResult{nil, &RuntimeError{u.operator, "Operand must be a number."}}
		}
		return evalResult{-(r), nil}
	case Bang:
		return evalResult{!isTruthy(right.value), nil}
	}
	// Unreachable
	return evalResult{nil, &RuntimeError{Token{}, "interpreter unary unreachable"}}
}
func (i *Interpreter) visitTernary(t *Ternary) any {
	if t.operatorL.Type == QuestionMark && t.operatorR.Type == Colon {
		left := i.evaluate(t.left)
		if left.err != nil {
			return evalResult{nil, left.err}
		}

		if isTruthy(left.value) {
			return i.evaluate(t.middle)
		}
		return i.evaluate(t.right)
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
