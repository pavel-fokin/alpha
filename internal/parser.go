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
	case tokens.VAR:
		return p.parseVar()
	case tokens.FUNC:
		return p.parseFunc()
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
	declaration := &Var{Token: p.curToken}

	if !p.expectPeek(tokens.IDENT) {
		return nil
	}

	declaration.Name = p.curToken.Literal

	if !p.expectPeek(tokens.IDENT) {
		return nil
	}

	declaration.Type = string(p.curToken.Literal)
	return declaration
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

	f.Params = p.parseFuncParams()
	f.Returns = p.parseFuncReturns()

	return f
}

func (p *Parser) parseFuncParams() []string {
	params := []string{}

	if p.peekTokenIs(tokens.RPAREN) {
		p.nextToken()
		return params
	}

	p.nextToken()

	params = append(params, p.curToken.Literal)

	for p.peekTokenIs(tokens.COMMA) {
		p.nextToken()
		p.nextToken()
		params = append(params, p.curToken.Literal)
	}

	if !p.expectPeek(tokens.RPAREN) {
		return nil
	}

	return params
}

func (p *Parser) parseFuncReturns() []string {
	returns := []string{}

	if p.peekTokenIs(tokens.IDENT) {
		p.nextToken()
		returns = append(returns, p.curToken.Literal)
	}

	return returns
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
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
