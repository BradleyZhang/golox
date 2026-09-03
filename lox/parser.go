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
// expression     → equality ;
// equality       → comparison ( ( "!=" | "==" ) comparison )* ; (等于)
// comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ; （比较）
// term           → factor ( ( "-" | "+" ) factor )* ; （加减）
// factor         → unary ( ( "/" | "*" ) unary )* ; （乘除）
// unary          → ( "!" | "-" ) unary
//                | primary ; （一元运算符）
// primary        → NUMBER | STRING | "true" | "false" | "nil"
//                | "(" expression ")" ; （字面量和括号表达式）

package lox

type Parser struct {
	tokens  []Token
	current int
}
type ParseError struct {
	Message string
}

func (e ParseError) Error() string {
	return e.Message
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) Parse() Expr {
	expr, err := p.expression()
	if err != nil {
		return nil
	}
	return expr
}

func (p *Parser) expression() (Expr, error) {
	return p.equality()
}

func (p *Parser) equality() (Expr, error) {
	expr, err := p.comparision()
	if err != nil {
		return nil, err
	}
	for p.match(BangEqual, EqualEqual) {
		operator := p.previous()
		right, err := p.comparision()
		if err != nil {
			return nil, err
		}
		expr = Binary{expr, operator, right}
	}
	return expr, nil
}
func (p *Parser) comparision() (Expr, error) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}
	for p.match(Greater, GreaterEqual, Less, LessEqual) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = Binary{expr, operator, right}
	}
	return expr, nil
}
func (p *Parser) term() (Expr, error) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}
	for p.match(Minus, Plus) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expr = Binary{expr, operator, right}
	}
	return expr, nil
}
func (p *Parser) factor() (Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}
	for p.match(Slash, Star) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = Binary{expr, operator, right}
	}
	return expr, nil
}
func (p *Parser) unary() (Expr, error) {
	if p.match(Bang, Minus) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return Unary{operator, right}, nil
	}
	return p.primary()
}

func (p *Parser) primary() (Expr, error) {
	if p.match(False) {
		return Literal{false}, nil
	}
	if p.match(True) {
		return Literal{true}, nil
	}
	if p.match(Nil) {
		return Literal{nil}, nil
	}
	if p.match(Number, String) {
		return Literal{p.previous().Literal}, nil
	}
	if p.match(LeftParen) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		if _, err := p.consume(RightParen, "Expect ')' after expression."); err != nil {
			return nil, err
		}
		return Grouping{expr}, nil
	}

	return nil, p.Error(p.peek(), "Expect expression.")
}

func (p *Parser) match(types ...TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) consume(tType TokenType, msg string) (Token, error) {
	if p.check(tType) {
		return p.advance(), nil
	}
	return Token{}, p.Error(p.peek(), msg)
}

func (p *Parser) Error(token Token, msg string) error {
	GlobalLox.TokenError(token, msg)
	return ParseError{Message: msg}
}
func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().Type == Semicolon {
			return
		}

		switch p.peek().Type {
		case Class, Fun, Var, For, If, While, Print, Return:
			return
		}
		p.advance()
	}
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
		return p.previous()
	}
	return Token{}
}
func (p *Parser) check(tType TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tType
}
func (p *Parser) isAtEnd() bool {
	return p.peek().Type == EOF
}
func (p *Parser) peek() Token {
	return p.tokens[p.current]
}
func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}
