package expressionparser

import (
	"fmt"
	"unicode"
)

// Token types
const (
	EOF TokenType = iota
	NUMBER
	PLUS
	MINUS
	MULT
	DIV
	LPAREN
	RPAREN
	INVALID
)

type TokenType int

// Token structure
type Token struct {
	Type  TokenType
	Value string
}

// Lexer converts input string into tokens
type Lexer struct {
	input  string
	pos    int
	ch     rune
}

// NewLexer creates a new Lexer
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar advances the position in the string and sets the current character.
func (l *Lexer) readChar() {
	if l.pos >= len(l.input) {
		l.ch = 0 // EOF
	} else {
		l.ch = rune(l.input[l.pos])
	}
	l.pos++
}

// NextToken returns the next token in the input.
func (l *Lexer) NextToken() Token {
	var tok Token

// Skip whitespace
for unicode.IsSpace(l.ch) {
	l.readChar()
}

// Handle EOF
if l.ch == 0 {
	return Token{Type: EOF}
}

// Handle numbers
if unicode.IsDigit(l.ch) {
	tok.Type = NUMBER
	tok.Value = l.readNumber()
	return tok
}

// Handle operators and parentheses
switch l.ch {
	case '+':
		tok = Token{Type: PLUS, Value: "+"}
	case '-':
		tok = Token{Type: MINUS, Value: "-"}
	case '*':
		tok = Token{Type: MULT, Value: "*"}
	case '/':
		tok = Token{Type: DIV, Value: "/"}
	case '(':
		tok = Token{Type: LPAREN, Value: "("}
	case ')':
		tok = Token{Type: RPAREN, Value: ")"}
	default:
		tok = Token{Type: INVALID, Value: fmt.Sprintf("Invalid character: %c", l.ch)}
}

l.readChar()
return tok
}

// readNumber reads a complete number (integer) from the input.
func (l *Lexer) readNumber() string {
	start := l.pos - 1
	for unicode.IsDigit(l.ch) {
	l.readChar()
}

return l.input[start:l.pos-1]

}

// Expression tree node types
type Expr interface{}

type Number struct {
	Value float64
}

type BinaryOp struct {
	Left  Expr
	Op    Token
	Right Expr
}

// Parser structure
type Parser struct {
	lexer *Lexer
	curr  Token
}

// NewParser creates a new parser instance
func NewParser(lexer *Lexer) *Parser {
	p := &Parser{lexer: lexer}
	p.nextToken()
	return p
}

// nextToken advances to the next token
func (p *Parser) nextToken() {
	p.curr = p.lexer.NextToken()
}

// Parse expression entry point
func (p *Parser) Parse() (Expr, error) {
	return p.parseExpr()
}

// parseExpr handles the parsing of the expression
func (p *Parser) parseExpr() (Expr, error) {

	// Start with parsing a term (handles operator precedence)
	left, err := p.parseTerm()
	if err != nil {
	return nil, err
	}

	// Handle addition and subtraction
	for p.curr.Type == PLUS || p.curr.Type == MINUS {
	op := p.curr
	p.nextToken()
	right, err := p.parseTerm()
	if err != nil {
		return nil, err
	}

	left = &BinaryOp{Left: left, Op: op, Right: right}
}

return left, nil
}

// parseTerm handles multiplication and division
func (p *Parser) parseTerm() (Expr, error) {
	left, err := p.parseFactor()
	if err != nil {
		return nil, err
	}

	// Handle multiplication and division
	for p.curr.Type == MULT || p.curr.Type == DIV {
		op := p.curr
		p.nextToken()
		right, err := p.parseFactor()

		if err != nil {
			return nil, err
	}
	left = &BinaryOp{Left: left, Op: op, Right: right}
}

return left, nil
}

// parseFactor handles numbers and parenthesized expressions
func (p *Parser) parseFactor() (Expr, error) {
	switch p.curr.Type {
		case NUMBER:
			value := p.curr.Value
			p.nextToken()
			return &Number{Value: parseNumber(value)}, nil
		case LPAREN:
			p.nextToken()
			expr, err := p.parseExpr()
			if err != nil {
				return nil, err
			}

			if p.curr.Type != RPAREN {
				return nil, fmt.Errorf("expected closing parenthesis")
		}

		p.nextToken()
		return expr, nil
		default:
			return nil, fmt.Errorf("expected a number or parenthesis, got %v", p.curr.Type)
	}
}

// parseNumber converts string to float64
func parseNumber(s string) float64 {
	var num float64
	
	fmt.Sscanf(s, "%f", &num)
	return num
}

// Eval evaluates an expression
func Eval(expr Expr) (float64, error) {
	switch v := expr.(type) {
		case *Number:
			return v.Value, nil
		case *BinaryOp:
			left, err := Eval(v.Left)
			if err != nil {
				return 0, err
			}
			right, err := Eval(v.Right)
			if err != nil {
				return 0, err
			}
	switch v.Op.Type {
		case PLUS:
			return left + right, nil
		case MINUS:
			return left - right, nil
		case MULT:
			return left * right, nil
		case DIV:
			if right == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			return left / right, nil
		}
	default:
		return 0, fmt.Errorf("unsupported expression type")
}

return 0, fmt.Errorf("invalid expression")
}

// end of file