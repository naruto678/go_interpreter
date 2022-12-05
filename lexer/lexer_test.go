package lexer

import (
	"interpreter/token"
	"testing"
)

func TestToken(t *testing.T) {
	input := "+(){},;"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for _, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test went wrong buddy")

		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Something went wrong again")
		}

	}

}
