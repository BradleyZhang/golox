package scanner

import (
	"golox/lox/token"
	"testing"
)

var tests = map[string][]token.Token{
	"var average = (min + max) / 2;": {
		{Type: token.Var, Lexeme: "var", Literal: nil, Line: 1},
		{Type: token.Identifier, Lexeme: "average", Literal: "average", Line: 1},
		{Type: token.Equal, Lexeme: "=", Literal: nil, Line: 1},
		{Type: token.LeftParen, Lexeme: "(", Literal: nil, Line: 1},
		{Type: token.Identifier, Lexeme: "min", Literal: "min", Line: 1},
		{Type: token.Plus, Lexeme: "+", Literal: nil, Line: 1},
		{Type: token.Identifier, Lexeme: "max", Literal: "max", Line: 1},
		{Type: token.RightParen, Lexeme: ")", Literal: nil, Line: 1},
		{Type: token.Slash, Lexeme: "/", Literal: nil, Line: 1},
		{Type: token.Number, Lexeme: "2", Literal: 2.0, Line: 1},
		{Type: token.Semicolon, Lexeme: ";", Literal: nil, Line: 1},
		{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
	},

	"var x = 123;": {
		{Type: token.Var, Lexeme: "var", Literal: nil, Line: 1},
		{Type: token.Identifier, Lexeme: "x", Literal: "x", Line: 1},
		{Type: token.Equal, Lexeme: "=", Literal: nil, Line: 1},
		{Type: token.Number, Lexeme: "123", Literal: 123.0, Line: 1},
		{Type: token.Semicolon, Lexeme: ";", Literal: nil, Line: 1},
		{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
	},

	"var pi = 3.14159;": {
		{Type: token.Var, Lexeme: "var", Literal: nil, Line: 1},
		{Type: token.Identifier, Lexeme: "pi", Literal: "pi", Line: 1},
		{Type: token.Equal, Lexeme: "=", Literal: nil, Line: 1},
		{Type: token.Number, Lexeme: "3.14159", Literal: 3.14159, Line: 1},
		{Type: token.Semicolon, Lexeme: ";", Literal: nil, Line: 1},
		{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
	},

	`print("hello world");`: {
		{Type: token.Print, Lexeme: "print", Literal: nil, Line: 1},
		{Type: token.LeftParen, Lexeme: "(", Literal: nil, Line: 1},
		{Type: token.String, Lexeme: `"hello world"`, Literal: "hello world", Line: 1},
		{Type: token.RightParen, Lexeme: ")", Literal: nil, Line: 1},
		{Type: token.Semicolon, Lexeme: ";", Literal: nil, Line: 1},
		{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
	},

	"x >= 10;": {
		{Type: token.Identifier, Lexeme: "x", Literal: "x", Line: 1},
		{Type: token.GreaterEqual, Lexeme: ">=", Literal: nil, Line: 1},
		{Type: token.Number, Lexeme: "10", Literal: 10.0, Line: 1},
		{Type: token.Semicolon, Lexeme: ";", Literal: nil, Line: 1},
		{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
	},

	"if (x == 10) print(x);": {
		{Type: token.If, Lexeme: "if", Literal: nil, Line: 1},
		{Type: token.LeftParen, Lexeme: "(", Literal: nil, Line: 1},
		{Type: token.Identifier, Lexeme: "x", Literal: "x", Line: 1},
		{Type: token.EqualEqual, Lexeme: "==", Literal: nil, Line: 1},
		{Type: token.Number, Lexeme: "10", Literal: 10.0, Line: 1},
		{Type: token.RightParen, Lexeme: ")", Literal: nil, Line: 1},
		{Type: token.Print, Lexeme: "print", Literal: nil, Line: 1},
		{Type: token.LeftParen, Lexeme: "(", Literal: nil, Line: 1},
		{Type: token.Identifier, Lexeme: "x", Literal: "x", Line: 1},
		{Type: token.RightParen, Lexeme: ")", Literal: nil, Line: 1},
		{Type: token.Semicolon, Lexeme: ";", Literal: nil, Line: 1},
		{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
	},

	"if (x > 10) { return true; }": {
		{Type: token.If, Lexeme: "if", Literal: nil, Line: 1},
		{Type: token.LeftParen, Lexeme: "(", Literal: nil, Line: 1},
		{Type: token.Identifier, Lexeme: "x", Literal: "x", Line: 1},
		{Type: token.Greater, Lexeme: ">", Literal: nil, Line: 1},
		{Type: token.Number, Lexeme: "10", Literal: 10.0, Line: 1},
		{Type: token.RightParen, Lexeme: ")", Literal: nil, Line: 1},
		{Type: token.LeftBrace, Lexeme: "{", Literal: nil, Line: 1},
		{Type: token.Return, Lexeme: "return", Literal: nil, Line: 1},
		{Type: token.True, Lexeme: "true", Literal: nil, Line: 1},
		{Type: token.Semicolon, Lexeme: ";", Literal: nil, Line: 1},
		{Type: token.RightBrace, Lexeme: "}", Literal: nil, Line: 1},
		{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
	},

	"for (var i = 0; i < 10; i = i + 1) print(i);": {
		{Type: token.For, Lexeme: "for", Literal: nil, Line: 1},
		{Type: token.LeftParen, Lexeme: "(", Literal: nil, Line: 1},
		{Type: token.Var, Lexeme: "var", Literal: nil, Line: 1},
		{Type: token.Identifier, Lexeme: "i", Literal: "i", Line: 1},
		{Type: token.Equal, Lexeme: "=", Literal: nil, Line: 1},
		{Type: token.Number, Lexeme: "0", Literal: 0.0, Line: 1},
		{Type: token.Semicolon, Lexeme: ";", Literal: nil, Line: 1},
		{Type: token.Identifier, Lexeme: "i", Literal: "i", Line: 1},
		{Type: token.Less, Lexeme: "<", Literal: nil, Line: 1},
		{Type: token.Number, Lexeme: "10", Literal: 10.0, Line: 1},
		{Type: token.Semicolon, Lexeme: ";", Literal: nil, Line: 1},
		{Type: token.Identifier, Lexeme: "i", Literal: "i", Line: 1},
		{Type: token.Equal, Lexeme: "=", Literal: nil, Line: 1},
		{Type: token.Identifier, Lexeme: "i", Literal: "i", Line: 1},
		{Type: token.Plus, Lexeme: "+", Literal: nil, Line: 1},
		{Type: token.Number, Lexeme: "1", Literal: 1.0, Line: 1},
		{Type: token.RightParen, Lexeme: ")", Literal: nil, Line: 1},
		{Type: token.Print, Lexeme: "print", Literal: nil, Line: 1},
		{Type: token.LeftParen, Lexeme: "(", Literal: nil, Line: 1},
		{Type: token.Identifier, Lexeme: "i", Literal: "i", Line: 1},
		{Type: token.RightParen, Lexeme: ")", Literal: nil, Line: 1},
		{Type: token.Semicolon, Lexeme: ";", Literal: nil, Line: 1},
		{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
	},

	"true and false or nil;": {
		{Type: token.True, Lexeme: "true", Literal: nil, Line: 1},
		{Type: token.And, Lexeme: "and", Literal: nil, Line: 1},
		{Type: token.False, Lexeme: "false", Literal: nil, Line: 1},
		{Type: token.Or, Lexeme: "or", Literal: nil, Line: 1},
		{Type: token.Nil, Lexeme: "nil", Literal: nil, Line: 1},
		{Type: token.Semicolon, Lexeme: ";", Literal: nil, Line: 1},
		{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
	},

	"!= == = ! >= > <= <": {
		{Type: token.BangEqual, Lexeme: "!=", Literal: nil, Line: 1},
		{Type: token.EqualEqual, Lexeme: "==", Literal: nil, Line: 1},
		{Type: token.Equal, Lexeme: "=", Literal: nil, Line: 1},
		{Type: token.Bang, Lexeme: "!", Literal: nil, Line: 1},
		{Type: token.GreaterEqual, Lexeme: ">=", Literal: nil, Line: 1},
		{Type: token.Greater, Lexeme: ">", Literal: nil, Line: 1},
		{Type: token.LessEqual, Lexeme: "<=", Literal: nil, Line: 1},
		{Type: token.Less, Lexeme: "<", Literal: nil, Line: 1},
		{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1},
	},

	`// comment
var x = 10;`: {
		{Type: token.Var, Lexeme: "var", Literal: nil, Line: 2},
		{Type: token.Identifier, Lexeme: "x", Literal: "x", Line: 2},
		{Type: token.Equal, Lexeme: "=", Literal: nil, Line: 2},
		{Type: token.Number, Lexeme: "10", Literal: 10.0, Line: 2},
		{Type: token.Semicolon, Lexeme: ";", Literal: nil, Line: 2},
		{Type: token.EOF, Lexeme: "", Literal: nil, Line: 2},
	},
}

func TestScanTokens(t *testing.T) {
	for source, test := range tests {
		scanner := NewScanner(source)
		tokens, _ := scanner.ScanTokens()

		if len(tokens) != len(test) {
			t.Errorf("%s:got %d tokens, want %d", source, len(tokens), len(test))
			return
		}
		for i, token := range tokens {
			if token.Type != test[i].Type {
				t.Errorf(
					"tokens[%d].Type = %v, want %v (%s in \"%s\")",
					i,
					token.Type,
					test[i].Type,
					test[i].Lexeme,
					source,
				)
			}
			if token.Literal != test[i].Literal {
				t.Errorf(
					"tokens[%d].Literal = %v, want %v (%s in \"%s\")",
					i,
					token.Literal,
					test[i].Literal,
					test[i].Lexeme,
					source,
				)
			}
			if token.Line != test[i].Line {
				t.Errorf(
					"tokens[%d].Line = %v, want %v (%s in \"%s\")",
					i,
					token.Line,
					test[i].Line,
					test[i].Lexeme,
					source,
				)
			}
		}
	}
}
