package enviroment

import (
	"golox/lox/loxerr"
	"golox/lox/token"
)

type Env struct {
	values map[string]any
}

func NewEnv() *Env {
	return &Env{
		values: map[string]any{},
	}
}
func (e *Env) Define(name string, value any) {
	e.values[name] = value
}
func (e *Env) Get(name token.Token) (any, error) { //TODO 之后把RuntimeError抽出来，这里返回Runtime Error
	if v, ok := e.values[name.Lexeme]; ok {
		return v, nil
	}
	return nil, loxerr.NewRuntimeError(name, "Undefined variable '"+name.Lexeme+"'.")
}
func (e *Env) Assign(name token.Token, value any) error {
	if _, ok := e.values[name.Lexeme]; ok {
		e.values[name.Lexeme] = value
		return nil
	} else {
		return loxerr.NewRuntimeError(name, "Undefined variable '"+name.Lexeme+"'.")
	}
}
