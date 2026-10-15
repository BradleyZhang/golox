package enviroment

import (
	"golox/lox/loxerr"
	"golox/lox/token"
)

type Env struct {
	enclosing *Env
	values    map[string]any
}

func NewEnvWithEnclosing(enclosing *Env) *Env {
	return &Env{
		enclosing: enclosing,
		values:    make(map[string]any),
	}
}
func NewEnv() *Env {
	return &Env{values: make(map[string]any)}
}
func (e *Env) Define(name string, value any) {
	e.values[name] = value
}
func (e *Env) Get(name token.Token) (any, error) { //TODO 之后把RuntimeError抽出来，这里返回Runtime Error
	if v, ok := e.values[name.Lexeme]; ok {
		return v, nil
	}
	if e.enclosing != nil {
		return e.enclosing.Get(name)
	}
	return nil, loxerr.NewRuntimeError(name, "Undefined variable '"+name.Lexeme+"'.")
}
func (e *Env) Assign(name token.Token, value any) error {
	if _, ok := e.values[name.Lexeme]; ok {
		e.values[name.Lexeme] = value
		return nil
	} else if e.enclosing != nil {
		return e.enclosing.Assign(name, value)
	} else {
		return loxerr.NewRuntimeError(name, "Undefined variable '"+name.Lexeme+"'.")
	}
}
