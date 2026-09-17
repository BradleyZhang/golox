package scanner

import (
	"golox/lox/token"
	"testing"
)

func TestScanTokens(t *testing.T) {
	source := "var average = (min + max) / 2;"

	scanner := NewScanner(source)
	tokens, _ := scanner.ScanTokens()
	wantType := []token.TokenType{
		token.Var,
		token.Identifier,
		token.Equal,
		token.LeftParen,
		token.Identifier,
		token.Plus,
		token.Identifier,
		token.RightParen,
		token.Slash,
		token.Number,
		token.Semicolon,
		token.EOF,
	}
	wantLiteral := []any{
		nil, "average", nil, nil, "min", nil, "max", nil, nil, 2.0, nil, nil,
	}
	if len(tokens) != len(wantType) {
		t.Fatalf("got %d tokens, want %d", len(tokens), len(wantType))
	}
	for i, token := range tokens {
		if token.Type != wantType[i] {
			t.Errorf(
				"tokens[%d].Type = %v, want %v",
				i,
				token.Type,
				wantType[i],
			)
		}
		if token.Literal != wantLiteral[i] {
			t.Errorf(
				"tokens[%d].Literal = %v, want %v",
				i,
				token.Literal,
				wantLiteral[i],
			)
		}
		if token.Line != 1 {
			t.Errorf(
				"tokens[%d].Line = %v, want %v",
				i,
				token.Line,
				1,
			)
		}

	}
}
