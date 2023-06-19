package internal

import "bytes"

type Node interface {
	TokenLiteral() string
	String() string
}

type Declaration interface {
	Node
	declarationNode()
}

type Program struct {
	Declarations []Declaration
}

func (p *Program) TokenLiteral() string {
	if len(p.Declarations) > 0 {
		return p.Declarations[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Declarations {
		out.WriteString(s.String())
	}

	return out.String()
}

type TypeDeclaration struct {
	Token Token
	Name  string
}

func (td *TypeDeclaration) declarationNode()     {}
func (td *TypeDeclaration) TokenLiteral() string { return td.Token.Literal }
func (td *TypeDeclaration) String() string {
	var out bytes.Buffer

	out.WriteString(td.TokenLiteral() + " ")
	out.WriteString(td.Name)

	return out.String()
}

type VarDeclaration struct {
	Token Token
	Name  string
	Type  string
}

func (vd *VarDeclaration) declarationNode()     {}
func (vd *VarDeclaration) TokenLiteral() string { return vd.Token.Literal }
func (vd *VarDeclaration) String() string {
	var out bytes.Buffer

	out.WriteString(vd.TokenLiteral() + " ")
	out.WriteString(vd.Name + " ")
	out.WriteString(vd.Type)

	return out.String()
}
