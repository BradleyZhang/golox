package lox

type TokenType int

const (
	// Single-character tokens.
	LeftParen  TokenType = 0  // (
	RightParen TokenType = 1  // )
	LeftBrace  TokenType = 2  // {
	RightBrace TokenType = 3  // }
	Comma      TokenType = 4  // ,
	Dot        TokenType = 5  // .
	Minus      TokenType = 6  // -
	Plus       TokenType = 7  // +
	Semicolon  TokenType = 8  // ;
	Slash      TokenType = 9  // /
	Star       TokenType = 10 // *

	// One or two character tokens.
	Bang         TokenType = 11 // !
	BangEqual    TokenType = 12 // !=
	Equal        TokenType = 13 // =
	EqualEqual   TokenType = 14 // ==
	Greater      TokenType = 15 // >
	GreaterEqual TokenType = 16 // >=
	Less         TokenType = 17 // <
	LessEqual    TokenType = 18 // <=

	// Literals.
	Identifier TokenType = 19
	String     TokenType = 20
	Number     TokenType = 21

	// Keywords.
	And    TokenType = 22
	Class  TokenType = 23
	Else   TokenType = 24
	False  TokenType = 25
	Fun    TokenType = 26
	For    TokenType = 27
	If     TokenType = 28
	Nil    TokenType = 29
	Or     TokenType = 30
	Print  TokenType = 31
	Return TokenType = 32
	Super  TokenType = 33
	This   TokenType = 34
	True   TokenType = 35
	Var    TokenType = 36
	While  TokenType = 37

	EOF TokenType = 38
)
