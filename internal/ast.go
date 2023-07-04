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
		out.WriteString("\n")
	}

	return out.String()
}

type Block struct {
	Token        Token
	Declarations []Declaration
}

func (b *Block) declarationNode()     {}
func (b *Block) TokenLiteral() string { return b.Token.Literal }
func (b *Block) String() string {
	var out bytes.Buffer

	out.WriteString("{\n")
	for _, s := range b.Declarations {
		out.WriteString("\t" + s.String())
		out.WriteString("\n")
	}
	out.WriteString("}\n")

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
	Type  string
	Name  string
}

func (vd *Var) declarationNode()     {}
func (vd *Var) TokenLiteral() string { return vd.Token.Literal }
func (vd *Var) String() string {
	var out bytes.Buffer

	out.WriteString(vd.Type + " ")
	out.WriteString(vd.Name)

	return out.String()
}

type Func struct {
	Token   Token
	Name    string
	Args    []*Var
	Returns []string // List of returned types.
}

func (f *Func) declarationNode()     {}
func (f *Func) TokenLiteral() string { return f.Token.Literal }
func (f *Func) String() string {
	var out bytes.Buffer

	args := []string{}

	for _, arg := range f.Args {
		args = append(args, arg.String())
	}

	out.WriteString(f.TokenLiteral() + " ")
	out.WriteString(f.Name + "(")
	out.WriteString(strings.Join(args, ", "))
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

type Return struct {
	Token Token
	Value string
}

func (r *Return) declarationNode()     {}
func (r *Return) TokenLiteral() string { return r.Token.Literal }
func (r *Return) String() string       { return r.Value }
