package lib

import (
	"fmt"
	"strconv"
)

type TokenType int8;
type Token struct {
	Type TokenType
	Value string
	Literal *float64
}
const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	DOT
	MINUS
	PLUS
	SLASH
	STAR
	CARAT
	BANG
	BAR
	NUMBER
	SEMICOLON
	COMMA
	IDENTIFIER
	EQUALS
	EOF

	PRINT 
	ROOT 
)

func (t TokenType) String() string {
	switch t {
	case LEFT_PAREN:
		return "LParen"
	case RIGHT_PAREN:
		return "RParen"
	case DOT:
		return "Dot"
	case MINUS:
		return "Minus"
	case PLUS:
		return "RightParen"
	case SLASH:
		return "Slash"
	case STAR:
		return "Star"
	case CARAT:
		return "RightParen"
	case BANG:
		return "Dot"
	case BAR:
		return "Bar"
	case NUMBER:
		return "Number"
	case PRINT:
		return "Print"
	case ROOT:
		return "Root"
	case SEMICOLON:
		return "Semicolon"
	case COMMA:
		return "Comma"
	case IDENTIFIER:
		return "Identifier"
	case EQUALS:
		return "Equals"
	case EOF:
		return "Eof"
	}
	return "Unknown";
}

type UnexpectedChar struct {
	Char string
	Index int
}

type Tokenizer struct {
	start uint32
	current uint32
	tokens []Token
	str string
}

func NewTokenizer(s string) *Tokenizer {
	return &Tokenizer{
		start: 0,
		current: 0,
		tokens: []Token{},
		str: s,
	};
}

type TokenizeError struct {
	msg string
}
func (t *TokenizeError) Error() string {
	return "Tokenizer: " + t.msg;
}
func newTokenizeError(msg string) *TokenizeError {
	return &TokenizeError{ msg: msg };
}

func (t *Tokenizer) Parse() (*[]Token, error) {
	for (!t.is_at_end()) {
		t.start = t.current;
		char := t.advance();
		switch char {
		case ' ':
		case '\r':
		case '\t':
		case '\n':
			break;
		case '(': t.add_token(LEFT_PAREN, nil); break;
		case ')': t.add_token(RIGHT_PAREN, nil); break;
		case '.': t.add_token(DOT, nil); break;
		case '-': t.add_token(MINUS, nil); break;
		case '+': t.add_token(PLUS, nil); break;
		case '*': t.add_token(STAR, nil); break;
		case '/': t.add_token(SLASH, nil); break;
		case ';': t.add_token(SEMICOLON, nil); break;
		case '=': t.add_token(EQUALS, nil); break;
		case ',': t.add_token(COMMA, nil); break;
		case '^': t.add_token(CARAT, nil); break;
		case '!': t.add_token(BANG, nil); break;
		case '|': t.add_token(BAR, nil); break;
		default:
			if t.is_digit(char) {
				err := t.number();
				if err != nil {
					return nil, err;
				}
			} else if t.is_alpha(char) {
				t.identifier();
			}
		}
	}
	return &t.tokens, nil;
}

func (t *Tokenizer) identifier() {
	for (t.is_alphanumeric(t.peek())) {
		t.advance();
	}
	text := t.str[t.start:t.current];
	switch text {
	case "p":
		t.add_token(PRINT, nil);
		break;
	case "root":
		t.add_token(ROOT, nil);
		break;
	default:
		t.add_token(IDENTIFIER, nil);
	}
}

func (t *Tokenizer) number() (error) {
	for t.is_digit(t.peek()) {
		t.advance();
	}
	if (t.peek() == '.' && t.is_digit(t.peek_next())) {
		t.advance();
		for t.is_digit(t.peek()) {
			t.advance();
		}
	}
	num, err := strconv.ParseFloat(t.str[t.start:t.current], 64);
	if err == nil {
		t.add_token(NUMBER, &num);
		return nil;
	} else {
		return newTokenizeError("Failed to parse float.");
	}
}
func (t *Tokenizer) is_digit(char byte) bool {
	return char >= '0' && char <= '9';
}
func (t *Tokenizer) is_alpha(char byte) bool {
	return (char >= 'a' && char <= 'z') || 
		(char >= 'A' && char <= 'Z') || 
		char == '_';
}
func (t *Tokenizer) is_alphanumeric(char byte) bool {
	return t.is_alpha(char) || t.is_digit(char);
}

func (t *Tokenizer) is_at_end() bool {
	return t.current >= uint32(len(t.str));
}
func (t *Tokenizer) advance() byte {
	c := t.current;
	t.current += 1;
	return t.str[c];
}
func (t *Tokenizer) peek() byte {
	if t.is_at_end() {
		return byte(0);
	}
	return t.str[t.current];
}
func (t *Tokenizer) peek_next() byte {
	if t.is_at_end() {
		return byte(0);
	}
	return t.str[t.current + 1];
}
func (t *Tokenizer) add_token(token_type TokenType, literal *float64) {
	token := Token {
		Type: token_type,
		Value: t.str[t.start:t.current],
		Literal: literal,
	};
	t.tokens = append(t.tokens, token);
}

func (t *Tokenizer) Scan() {
	fmt.Println("I printed something~");
}
