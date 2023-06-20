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

type AST struct {
	Declarations []Declaration
}

func (a *AST) TokenLiteral() string {
	if len(a.Declarations) > 0 {
		return a.Declarations[0].TokenLiteral()
	} else {
		return ""
	}
}

func (a *AST) String() string {
	var out bytes.Buffer

	for _, s := range a.Declarations {
		out.WriteString(s.String())
	}

	return out.String()
}

type Type struct {
	Token Token
	Name  string
}

func (td *Type) declarationNode()     {}
func (td *Type) TokenLiteral() string { return td.Token.Literal }
func (td *Type) String() string {
	var out bytes.Buffer

	out.WriteString(td.TokenLiteral() + " ")
	out.WriteString(td.Name)

	return out.String()
}

type Var struct {
	Token Token
	Name  string
	Type  string
}

func (vd *Var) declarationNode()     {}
func (vd *Var) TokenLiteral() string { return vd.Token.Literal }
func (vd *Var) String() string {
	var out bytes.Buffer

	out.WriteString(vd.TokenLiteral() + " ")
	out.WriteString(vd.Name + " ")
	out.WriteString(vd.Type)

	return out.String()
}
