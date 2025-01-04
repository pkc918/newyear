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
	ASSIGN = "="
	PLUS   = "+"

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
