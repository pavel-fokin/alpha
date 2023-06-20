package internal

type Parser struct {
	lexer *Lexer

	curToken  Token
	peekToken Token
}

func NewParser(l *Lexer) *Parser {
	return &Parser{lexer: l}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{}
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
		return p.parseTypeDeclaration()
	case tokens.VAR:
		return p.parseVarDeclaration()
	default:
		return nil
	}
}

func (p *Parser) parseVarDeclaration() *VarDeclaration {
	declaration := &VarDeclaration{Token: p.curToken}

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

func (p *Parser) parseTypeDeclaration() *TypeDeclaration {
	declaration := &TypeDeclaration{Token: p.curToken}

	if !p.expectPeek(tokens.IDENT) {
		return nil
	}

	declaration.Name = p.curToken.Literal
	return declaration
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
