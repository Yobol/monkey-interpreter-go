package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL" // 遇到未知的词法单元
	EOF     TokenType = "EOF"     // 文件结尾

	IDENT TokenType = "IDENT" // 变量名
	INT   TokenType = "INT"   // 整型

	// 运算符
	ASSIGN TokenType = "=" // 赋值运算符
	PLUS   TokenType = "+" // 加法运算符

	// 分隔符
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	// 关键字
	FUNCTION TokenType = "fn"
	LET      TokenType = "let"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func IdentifyType(ident string) TokenType {
	if tt, ok := keywords[ident]; ok {
		return tt
	}
	return IDENT
}
