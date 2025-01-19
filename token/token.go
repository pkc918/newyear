package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// 特殊标识
	ILLEGAL = "ILLEGAL" // 非法字符
	EOF     = "EOF"     // End Of File

	// 标识符 + 字面量
	IDENT = "IDENT"
	INT   = "INT"

	// 运算符
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"function": FUNCTION,
	"let":      LET,
}

// LookupIdent 通过检查关键字表来判断给定的标识符是否是关键字。如果是，返回关键字的 TokenType 常量。如果不是，但会 token.IDENT
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
