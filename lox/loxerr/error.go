package loxerr

import "golox/lox/token"

type RuntimeError struct { //TODO 之后抽出
	Token   token.Token
	Message string
}

func NewRuntimeError(t token.Token, msg string) *RuntimeError {
	return &RuntimeError{
		Token:   t,
		Message: msg,
	}
}

func (e *RuntimeError) Error() string {
	return e.Message
}
