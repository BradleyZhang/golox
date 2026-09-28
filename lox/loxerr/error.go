package loxerr

import "golox/lox/token"

type RuntimeError struct { //TODO 之后抽出
	Token   token.Token
	Message string
}

func (e *RuntimeError) Error() string {
	return e.Message
}
