package internal

type Parser struct {
	lexer *Lexer

	curToken  Token
	peekToken Token
}

func NewParser(l *Lexer) *Parser {
	return &Parser{lexer: l}
}

func (p *Parser) Parse() *AST {
	program := &AST{}
	program.Declarations = []Declaration{}

	// We need to fill in curToken and peekToken.
	p.nextToken()
	p.nextToken()

	for p.curToken.Type != tokens.EOF {
		declaration := p.parseDeclaration()
		if declaration != nil {
			program.Declarations = append(program.Declarations, declaration)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseDeclaration() Declaration {
	switch p.curToken.Type {
	case tokens.TYPE:
		return p.parseType()
	case tokens.FUNC:
		return p.parseFunc()
	case tokens.LBRACE:
		return p.parseBlock()
	case tokens.RETURN:
		return p.parseReturn()
	case tokens.IDENT:
		switch p.peekToken.Type {
		case tokens.IDENT:
			return p.parseVar()
		default:
			return nil
		}
	default:
		return nil
	}
}

func (p *Parser) parseType() *Type {
	declaration := &Type{Token: p.curToken}

	if !p.expectPeek(tokens.IDENT) {
		return nil
	}

	declaration.Name = p.curToken.Literal
	return declaration
}

func (p *Parser) parseVar() *Var {
	v := &Var{Token: p.curToken}

	v.Type = p.curToken.Literal

	if !p.expectPeek(tokens.IDENT) {
		return nil
	}

	v.Name = p.curToken.Literal

	return v
}

func (p *Parser) parseBlock() *Block {
	b := &Block{Token: p.curToken}

	p.nextToken()

	for !p.curTokenIs(tokens.RBRACE) && !p.curTokenIs(tokens.EOF) {
		stmt := p.parseDeclaration()
		if stmt != nil {
			b.Declarations = append(b.Declarations, stmt)
		}
		p.nextToken()
	}
	return b
}

func (p *Parser) parseFunc() *Func {
	f := &Func{Token: p.curToken}

	if !p.expectPeek(tokens.IDENT) {
		return nil
	}

	f.Name = p.curToken.Literal

	if !p.expectPeek(tokens.LPAREN) {
		return nil
	}

	f.Args = p.parseFuncParams()
	f.Returns = p.parseFuncReturns()

	return f
}

func (p *Parser) parseFuncParams() []*Var {
	args := []*Var{}

	if p.peekTokenIs(tokens.RPAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()

	args = append(args, p.parseVar())

	for p.peekTokenIs(tokens.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseVar())
	}

	if !p.expectPeek(tokens.RPAREN) {
		return nil
	}

	return args
}

func (p *Parser) parseFuncReturns() []string {
	returns := []string{}

	if p.peekTokenIs(tokens.IDENT) {
		p.nextToken()
		returns = append(returns, p.curToken.Literal)
	}

	return returns
}

func (p *Parser) parseReturn() *Return {
	r := &Return{Token: p.curToken}

	p.nextToken()
	// if !p.expectPeek(tokens.IDENT) {
	// 	return nil
	// }

	r.Value = p.curToken.Literal

	return r
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) curTokenIs(t TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(tokenType TokenType) bool {
	return p.peekToken.Type == tokenType
}

func (p *Parser) expectPeek(tokenType TokenType) bool {
	if p.peekTokenIs(tokenType) {
		p.nextToken()
		return true
	}
	return false
}
