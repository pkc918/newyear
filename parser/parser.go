package parser

import (
	"github.com/newyear/ast"
	"github.com/newyear/lexer"
	"github.com/newyear/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token // 词法分析器实例的指针，实例上重复调用 NextToken 不断获取输入的下一个词法单元
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// 读取两个词法单元，相当于初始化 curToken 和 peekToken
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseProgram() *ast.Program {
	return nil
}
