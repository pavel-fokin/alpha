package internal

import (
	"bytes"
	"strings"
)

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

type Func struct {
	Token   Token
	Name    string
	Params  []string
	Returns []string // List of returned types.
}

func (f *Func) declarationNode()     {}
func (f *Func) TokenLiteral() string { return f.Token.Literal }
func (f *Func) String() string {
	var out bytes.Buffer

	out.WriteString(f.TokenLiteral() + " ")
	out.WriteString(f.Name + "(")
	out.WriteString(strings.Join(f.Params, ", "))
	out.WriteString(") ")

	if len(f.Returns) == 0 {
		// Empty returns.
	} else if len(f.Returns) == 1 {
		out.WriteString(f.Returns[0])
	} else {
		out.WriteString(f.Name + "(")
		out.WriteString(strings.Join(f.Returns, ", "))
		out.WriteString(")")
	}

	return out.String()
}
