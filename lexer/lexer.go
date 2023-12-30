package lexer

import (
	"monkey/token"
)

func New(source string) *Lexer {
	l := &Lexer{
		source: source,
	}
	l.readChar()
	return l
}

/**
 * 词法分析器会将源码作为输入，并输出对应的词法单元。
 *
 * 词法分析器会遍历输入的字符，然后逐个输出识别到的词法单元。
 */
type Lexer struct {
	source       string
	position     int  // 所输入源代码的当前位置（指向当前字符）
	ch           byte // 当前遍历字符，对应 position
	readPosition int  // 所输入源代码的当前读取位置（指向下一个要读取字符）
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	var tk token.Token
	switch l.ch {
	case 0: // 特殊处理文件末尾
		tk.Type = token.EOF
		tk.Literal = ""
	case '=':
		tk = newToken(token.ASSIGN, l.ch)
	case '+':
		tk = newToken(token.PLUS, l.ch)
	case '(':
		tk = newToken(token.LPAREN, l.ch)
	case ')':
		tk = newToken(token.RPAREN, l.ch)
	case '{':
		tk = newToken(token.LBRACE, l.ch)
	case '}':
		tk = newToken(token.RBRACE, l.ch)
	case ',':
		tk = newToken(token.COMMA, l.ch)
	case ';':
		tk = newToken(token.SEMICOLON, l.ch)
	default:
		if isLetter(l.ch) { // 这里决定了变量名不能以数字开头
			tk.Literal = l.readIdentifier()
			tk.Type = token.IdentifyType(tk.Literal)
			return tk
		} else if isDigit(l.ch) {
			tk.Literal = l.readNumber()
			tk.Type = token.INT
			return tk
		} else {
			tk = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tk
}
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

/**
 * readChar 用于读取输入源码中的下一个字符，并前移其位置。
 */
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.source) {
		l.ch = 0
	} else {
		l.ch = l.source[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.source[position:l.position]
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.source[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

/**
 * skipWhitespace 用于跳过所有无意义的空白字符。
 * 在某些词法分析器中，该函数也成为 eatWhitespace 或 consumeWhitespace。
 */
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}
