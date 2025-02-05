package ast

import (
	"bytes"
	"github.com/newyear/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

// Statement 语句
type Statement interface {
	Node
	statementNode()
}

// Expression 表达式
type Expression interface {
	Node
	expressionNode()
}

// Program 语法分析器生成的每个AST的根结点
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// String 创建一个缓冲区，将每条语句的String方法的返回值写入该缓冲区，最后将缓冲区以字符串形式返回。
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

type LetStatement struct {
	Token token.Token // token.LET词法单元
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ") // 写入 let
	out.WriteString(ls.Name.String())        // 写入标识符
	out.WriteString("=")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";") // 写入结果
	return out.String()
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	var out bytes.Buffer
	if es.Expression != nil {
		out.WriteString(es.Expression.String())
	}

	return out.String()
}
