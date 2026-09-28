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
	return nil, &loxerr.RuntimeError{Token: name, Message: "Undefined variable '" + name.Lexeme + "'."}
}
