package lexer

import (
	"github.com/newyear/token"
	"regexp"
)

type Lexer struct {
	input        string
	position     int  // 所输入字符中的当前位置
	readPosition int  // 所输入字符串中的当前读取位置（指向当前字符之后的一个字符）
	ch           byte // 当前正在查看的字符
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		tok = newToken(token.BANG, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// skipWhitespace 跳过空白字符（作为词法单元的分隔符）
func (l *Lexer) skipWhitespace() {
	re := regexp.MustCompile(`[ \t\n\r]`)
	for matched := re.MatchString(string(l.ch)); matched; {
		l.readChar()
	}
}

// isLetter 判断是否是字母
func isLetter(ch byte) bool {
	matched, err := regexp.Match(`^[a-zA-Z_]+`, []byte(string(ch)))
	if err != nil {
		panic(err)
	}
	return matched
}

// readIdentifier 从一个有效token开始读取到一个分隔符之前，代表了一个单元
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 判断是否是数字
func isDigit(ch byte) bool {
	matched, err := regexp.Match(`[0-9]`, []byte(string(ch)))
	if err != nil {
		panic(err)
	}
	return matched
}

// readNumber 从一个有效token开始读取到一个分隔符之前，代表了一个单元，读取数字
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
