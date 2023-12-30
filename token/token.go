package token

type Token struct {
	Type    TokenType // 保存词法单元的类型
	Literal string    // 保存词法单元的字面量
}
