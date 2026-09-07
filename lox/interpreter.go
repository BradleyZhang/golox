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
	v, err := a.value(expr)
	if err != nil {
		GlobalLox.RuntimeError(err)
		return
	}
	fmt.Println(stringify(v))
}

// TODO：断言失败处理
func (a *Interpreter) value(expr Expr) (any, *RuntimeError) {
	switch e := expr.(type) {
	case Binary:
		left, err := a.value(e.left)
		if err != nil {
			return nil, err
		}
		right, _err := a.value(e.right)
		if _err != nil {
			return nil, _err
		}
		switch e.operator.Type {
		case Minus, Slash, Star, Greater, GreaterEqual, Less, LessEqual:
			l, okL := left.(float64)
			r, okR := right.(float64)
			if !okL || !okR {
				return nil, &RuntimeError{e.operator, "Operand must be a number."}
			}
			switch e.operator.Type {
			case Minus:
				return l - r, nil
			case Slash:
				return l / r, nil
			case Star:
				return l * r, nil
			case Greater:
				return l > r, nil
			case GreaterEqual:
				return l >= r, nil
			case Less:
				return l < r, nil
			case LessEqual:
				return l <= r, nil
			}
		case Plus:
			{
				l, okL := left.(float64)
				r, okR := right.(float64)
				if okL && okR {
					return l + r, nil
				}
			}
			{
				l, okL := left.(string)
				r, okR := right.(string)
				if okL && okR {
					return l + r, nil
				}
			}
			return nil, &RuntimeError{e.operator, "Operands must be two numbers or two strings."}

		case BangEqual:
			return !isEqual(left, right), nil
		case EqualEqual:
			return isEqual(left, right), nil

		}
	case Grouping:
		return a.value(e.expression)
	case Literal:
		return e.value, nil
	case Unary:
		right, err := a.value(e.right)
		if err != nil {
			return nil, err
		}
		switch e.operator.Type {
		case Minus:
			r, ok := right.(float64)
			if !ok {
				return nil, &RuntimeError{e.operator, "Operand must be a number."}
			}
			return -(r), nil
		case Bang:
			return !isTruthy(right), nil
		}

	case Ternary:
		if e.operatorL.Type == QuestionMark && e.operatorR.Type == Colon {
			left, err := a.value(e.left)
			if err != nil {
				return nil, err
			}
			middle, err := a.value(e.middle)
			if err != nil {
				return nil, err
			}
			right, err := a.value(e.right)
			if err != nil {
				return nil, err
			}
			if isTruthy(left) {
				return middle, nil
			} else {
				return right, nil
			}
		}
	}
	// Unreachable
	return nil, &RuntimeError{Token{}, "unreachable"} // TODO 错误信息
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
