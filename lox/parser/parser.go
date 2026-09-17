// expression     → literal
//                | unary
//                | binary
//                | grouping ;

// literal        → NUMBER | STRING | "true" | "false" | "nil" ;
// grouping       → "(" expression ")" ;
// unary          → ( "-" | "!" ) expression ;
// binary         → expression operator expression ;
// operator       → "==" | "!=" | "<" | "<=" | ">" | ">="
//                | "+"  | "-"  | "*" | "/" ;

// 优先级由低到高，每个规则只匹配>=自身优先级的表达式
// 以 1+2*3为例， term 要求左右是高一级的完整factor，这样把表达式切成 (1) + (2*3) ，可以理解为低优先级的决定在哪里切开，然后把左右交给高优先级
// expression     → equality ;
// comma          → conditional ((",") conditional )* ; (逗号运算符)
// conditional    → equality ( "?" equality ":" conditional )?
// equality       → comparison ( ( "!=" | "==" ) comparison )* ; (等于)
// comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ; （比较）
// term           → factor ( ( "-" | "+" ) factor )* ; （加减）
// factor         → unary ( ( "/" | "*" ) unary )* ; （乘除）
// unary          → ( "!" | "-" ) unary
//                | binaryWithoutLeftOperand ; （一元运算符）
// binaryWithoutLeftOperand → 二元运算符 错误处理
//				  | primary
// primary        → NUMBER | STRING | "true" | "false" | "nil"
//                | "(" expression ")" ; （字面量和括号表达式）

// 优先级从低到高
// | Name              | Operators         | Associates    |
// | ----------------- | ----------------- | ------------- |
// | Comma 逗号运算符    | `,`               | Left 左结合   |
// | Conditional 条件运算符 | `:?`           | Right 右结合   |
// | Equality  等于    | `==` `!=`         | Left  左结合  |
// | Comparison  比较  | `>` `>=` `<` `<=` | Left  左结合  |
// | Term  加减运算    | `-` `+`           | Left  左结合  |
// | Factor   乘除运算 | `/` `*`           | Left  左结合  |
// | Unary  一元运算符 | `!` `-`           | Right  右结合 |

// 语句
// program        → statement* EOF ;
// statement      → exprStmt
//                | printStmt ;

// exprStmt       → expression ";" ;
// printStmt      → "print" expression ";" ;
package parser

import (
	"golox/lox/ast"
	"golox/lox/token"
)

type Parser struct {
	tokens  []token.Token
	current int
}
type ParseError struct {
	Message string
}

func (e ParseError) Error() string {
	return e.Message
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{tokens: tokens}
}

type TokenError struct {
	Token token.Token
	Msg   string
}

func (p *Parser) Parse() ([]ast.Stmt, *TokenError) {
	statements := []ast.Stmt{}
	for !p.isAtEnd() {
		statement, err := p.statement()
		if err != nil {
			return nil, err
		}
		statements = append(statements, statement)
	}
	return statements, nil
}
func (p *Parser) statement() (ast.Stmt, *TokenError) {
	if p.match(token.Print) {
		return p.printStatement()
	}
	return p.expressionStatement()
}
func (p *Parser) printStatement() (ast.Stmt, *TokenError) {
	value, err := p.expression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(token.Semicolon, "Expect ';' after value.")
	if err != nil {
		return nil, err
	}
	return ast.PrintStmt{value}, nil
}
func (p *Parser) expressionStatement() (ast.Stmt, *TokenError) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(token.Semicolon, "Expect ';' after value.")
	if err != nil {
		return nil, err
	}
	return ast.Expression{expr}, nil

}
func (p *Parser) expression() (ast.Expr, *TokenError) {
	return p.comma()
}

func (p *Parser) comma() (ast.Expr, *TokenError) {
	expr, err := p.conditional()
	if err != nil {
		return nil, err
	}
	for p.match(token.Comma) {
		operator := p.previous()
		right, err := p.conditional()
		if err != nil {
			return nil, err
		}
		expr = ast.Binary{expr, operator, right}
	}
	return expr, nil
}
func (p *Parser) conditional() (ast.Expr, *TokenError) {
	expr, err := p.equality()
	if err != nil {
		return nil, err
	}
	if p.match(token.QuestionMark) {
		operatorL := p.previous()
		middle, err := p.equality()
		if err != nil {
			return nil, err
		}
		p.consume(token.Colon, "Expect ':' after expression.") // TODO 错误提示
		operatorR := p.previous()
		right, err := p.conditional()
		if err != nil {
			return nil, err
		}
		expr = ast.Ternary{
			Left:      expr,
			OperatorL: operatorL,
			Middle:    middle,
			OperatorR: operatorR,
			Right:     right,
		}
	}
	return expr, nil
}

func (p *Parser) equality() (ast.Expr, *TokenError) {
	expr, err := p.comparision()
	if err != nil {
		return nil, err
	}
	for p.match(token.BangEqual, token.EqualEqual) {
		operator := p.previous()
		right, err := p.comparision()
		if err != nil {
			return nil, err
		}
		expr = ast.Binary{expr, operator, right}
	}
	return expr, nil
}
func (p *Parser) comparision() (ast.Expr, *TokenError) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}
	for p.match(token.Greater, token.GreaterEqual, token.Less, token.LessEqual) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = ast.Binary{expr, operator, right}
	}
	return expr, nil
}
func (p *Parser) term() (ast.Expr, *TokenError) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}
	for p.match(token.Minus, token.Plus) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expr = ast.Binary{expr, operator, right}
	}
	return expr, nil
}
func (p *Parser) factor() (ast.Expr, *TokenError) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}
	for p.match(token.Slash, token.Star) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = ast.Binary{expr, operator, right}
	}
	return expr, nil
}
func (p *Parser) unary() (ast.Expr, *TokenError) {
	if p.match(token.Bang, token.Minus) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return ast.Unary{operator, right}, nil
	}
	return p.binaryWithoutLeftOperand()
}

// 错误处理
// 没有左操作数的二元操作符
func (p *Parser) binaryWithoutLeftOperand() (ast.Expr, *TokenError) {
	switch {
	case p.match(token.Comma):
		err := p.Error(p.previous(), "Expect left-hand operand before binary operator")
		p.comma() // TODO 同步处理暂时无法测试，之后判断是应该comma 还是高一级，后面同理
		return nil, err
	case p.match(token.BangEqual, token.EqualEqual):
		err := p.Error(p.previous(), "Expect left-hand operand before binary operator")
		p.equality()
		return nil, err
	case p.match(token.Less, token.LessEqual, token.Greater, token.GreaterEqual):
		err := p.Error(p.previous(), "Expect left-hand operand before binary operator")
		p.comparision()
		return nil, err
	case p.match(token.Minus, token.Plus):
		err := p.Error(p.previous(), "Expect left-hand operand before binary operator")
		p.term()
		return nil, err
	case p.match(token.Slash, token.Star):
		err := p.Error(p.previous(), "Expect left-hand operand before binary operator")
		p.factor()
		return nil, err
	}
	return p.primary()
}

func (p *Parser) primary() (ast.Expr, *TokenError) {
	if p.match(token.False) {
		return ast.Literal{false}, nil
	}
	if p.match(token.True) {
		return ast.Literal{true}, nil
	}
	if p.match(token.Nil) {
		return ast.Literal{nil}, nil
	}
	if p.match(token.Number, token.String) {
		return ast.Literal{p.previous().Literal}, nil
	}
	if p.match(token.LeftParen) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		if _, err := p.consume(token.RightParen, "Expect ')' after expression."); err != nil {
			return nil, err
		}
		return ast.Grouping{expr}, nil
	}

	return nil, p.Error(p.peek(), "Expect expression.")
}

func (p *Parser) match(types ...token.TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) consume(tType token.TokenType, msg string) (token.Token, *TokenError) {
	if p.check(tType) {
		return p.advance(), nil
	}
	return token.Token{}, p.Error(p.peek(), msg)
}

func (p *Parser) Error(token token.Token, msg string) *TokenError {
	// lox.GlobalLox.TokenError(token, msg)
	return &TokenError{Token: token, Msg: msg}
}
func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().Type == token.Semicolon {
			return
		}

		switch p.peek().Type {
		case token.Class, token.Fun, token.Var, token.For, token.If, token.While, token.Print, token.Return:
			return
		}
		p.advance()
	}
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.current++
		return p.previous()
	}
	return token.Token{}
}
func (p *Parser) check(tType token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tType
}
func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}
func (p *Parser) peek() token.Token {
	return p.tokens[p.current]
}
func (p *Parser) previous() token.Token {
	return p.tokens[p.current-1]
}
