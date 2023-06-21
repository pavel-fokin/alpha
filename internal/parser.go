package internal

import "fmt"

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
	p.nextToken()
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

	fmt.Println("p.curToken.Literal", p.curToken.Literal)
	fmt.Println("p.peekToken.Literal", p.peekToken.Literal)

	if !(p.peekTokenIs(tokens.IDENT) || p.peekTokenIs(tokens.LPAREN)) {
		fmt.Println("### 0", !(p.peekTokenIs(tokens.IDENT) || p.peekTokenIs(tokens.LPAREN)))
		return returns
	}

	fmt.Println("### 1")

	if p.peekTokenIs(tokens.IDENT) {
		fmt.Println("#", p.curToken.Literal)
		returns = append(returns, p.curToken.Literal)
		p.nextToken()
		return returns
	}

	fmt.Println("### 2")

	if p.peekTokenIs(tokens.RPAREN) {
		p.nextToken()
		return returns
	}

	p.nextToken()

	returns = append(returns, p.curToken.Literal)

	for p.peekTokenIs(tokens.COMMA) {
		p.nextToken()
		p.nextToken()
		returns = append(returns, p.curToken.Literal)
	}

	if !p.expectPeek(tokens.RPAREN) {
		return nil
	}

	return returns
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
