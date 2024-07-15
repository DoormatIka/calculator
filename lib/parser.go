package lib;

import "slices"

type RecursiveDescentParser struct {
	current uint16
	tokens []Token
	measurements []string
}

type ParserError struct {
	msg string
}
func (t *ParserError) Error() string {
	return "Parser: " + t.msg;
}
func newParseError(msg string) *ParserError {
	return &ParserError{ msg: msg };
}

func NewRecursiveDescentParser(measurements []string) RecursiveDescentParser {
	return RecursiveDescentParser{
		current: 0,
		tokens: []Token{},
		measurements: measurements,
	};
}
func (p *RecursiveDescentParser) Parse(tokens []Token) ([]Stmt, *ParserError) {
	p.tokens = tokens;

	statements := []Stmt{};
	for !p.isAtEnd() {
		decl, err := p.printStatement();
		if err != nil {
			return nil, err;
		}
		statements = append(statements, decl);
	}
	return statements, nil;
}

func (p *RecursiveDescentParser) printStatement() (Stmt, *ParserError) {
	peeked_token := p.peek().Type;
	unallowed_token_types := []TokenType{EQUALS, PLUS, SLASH, STAR};
	for _, v := range unallowed_token_types {
		if peeked_token == v {
			return nil, newParseError("You can't operate on a print statement.");
		}
	}
	if (p.peek().Type == SEMICOLON) {
		return nil, newParseError("You need to put an expression like 1 or 1 + 2 into it.");
	}
	expr, err := p.expression();
	if err != nil {
		return nil, err;
	}
	_, err = p.consume(SEMICOLON, "Expect ';' after value.");
	if err != nil {
		return nil, err;
	}
	p_stmt := PrintStmt {expression: &expr};
	return &p_stmt, nil;
}

func (p *RecursiveDescentParser) expression() (Expr, *ParserError) {
	val, err := p.unary();
	if err != nil {
		return nil, err;
	}
	return val, nil;
}

func (p *RecursiveDescentParser) unary() (Expr, *ParserError) {
	if (p.match_and_advance(MINUS)) {
		operator := p.previous();
		right, err := p.primary();
		if err != nil {
			return nil, err;
		}
		unary_obj := UnaryExpr {
			operator: *operator,
			right: &right,
		};
		return &unary_obj, nil;
	}
	return p.primary();
}
func (p *RecursiveDescentParser) primary() (Expr, *ParserError) {
	if p.match_and_advance(NUMBER) {
		value := p.previous().Literal;
		var number_label string;
		if (p.match_and_advance(IDENTIFIER)) {
			number_label = p.previous().Text;
		}
		if number_label != "" && slices.Index(p.measurements, number_label) == -1 {
			return nil, newParseError("Unknown measurement type.");
		}
		literal := LiteralExpr {
			value: *value,
			label: &number_label,
		};
		return &literal, nil;
	}
	if p.match_and_advance(LEFT_PAREN, BAR) {
		prev := p.previous();
		expr, err := p.expression();
		if err != nil {
			return nil, err;
		}
		switch prev.Type {
			case LEFT_PAREN:
				p.consume(RIGHT_PAREN, "Expected ')' after expression.");
				break;
			case BAR:
				p.consume(BAR, "Expected '|' after expression.");
				break;
			default:
				return nil, newParseError("Something went horribly wrong when parsing groupings.");
		}
		grouping := GroupingExpr {
			operator: *prev,
			expression: &expr,
		}
		return &grouping, nil;
	}
	return nil, newParseError("Expected an expression.");
}

//// HELPERS ////
func (p *RecursiveDescentParser) consume(Type TokenType, message string) (*Token, *ParserError) {
	if (p.check(Type)) {
		return p.advance(), nil;
	}
	return nil, newParseError(message);
}
func (p *RecursiveDescentParser) isAtEnd() bool {
	return p.peek().Type == EOF;
}
func (p *RecursiveDescentParser) peek() *Token {
	return &p.tokens[p.current]; // out of range. this problems opens pandora's box.
}
func (p *RecursiveDescentParser) peek_next() *Token {
	if p.isAtEnd() {
		return nil;
	}
	return &p.tokens[p.current + 1];
}
func (p *RecursiveDescentParser) previous() *Token {
	if p.isAtEnd() {
		return nil;
	}
	return &p.tokens[p.current - 1];
}
func (p *RecursiveDescentParser) advance() *Token {
	if (!p.isAtEnd()) {
		p.current += 1;
	}
	return p.previous();
}
func (p *RecursiveDescentParser) check(Type TokenType) bool {
	if (p.isAtEnd()) {
		return false;
	}
	return p.peek().Type == Type;
}
func (p *RecursiveDescentParser) match_and_advance(types ...TokenType) bool {
	for _, v := range types {
		if (p.check(v)) {
			p.advance();
			return true;
		}
	}
	return false;
}

func (p *RecursiveDescentParser) synchronize() {
	p.advance();
	for (!p.isAtEnd()) {
		if (p.previous().Type == SEMICOLON) {
			return;
		}
		switch (p.peek().Type) {
			case PRINT:
			case IDENTIFIER:
				return;
		}
		p.advance();
	}
}
