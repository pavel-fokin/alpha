package internal

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() Token {
	var token Token

	l.skipWhitespace()

	switch l.ch {
	case ',':
		token.Literal = ","
		token.Type = tokens.COMMA
	case '(':
		token.Literal = "("
		token.Type = tokens.LPAREN
	case ')':
		token.Literal = ")"
		token.Type = tokens.RPAREN
	case '{':
		token.Literal = "{"
		token.Type = tokens.LBRACE
	case '}':
		token.Literal = "}"
		token.Type = tokens.RBRACE
	case 0:
		token.Literal = ""
		token.Type = tokens.EOF
	default:
		if isLetter(l.ch) {
			token.Literal = l.readIdentifier()
			token.Type = LookupKeyword(token.Literal)
			return token
		} else {
			token.Literal = string(l.ch)
			token.Type = tokens.ILLEGAL
		}
	}

	l.readChar()
	return token
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

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
