package parser

import (
	"fmt"
	"github.com/newyear/ast"
	"github.com/newyear/lexer"
	"github.com/newyear/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token // 词法分析器实例的指针，实例上重复调用 NextToken 不断获取输入的下一个词法单元
	peekToken token.Token

	errors []string // 收集报错信息

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	// 读取两个词法单元，相当于初始化 curToken 和 peekToken
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// parseProgram 解析程序
func (p *Parser) parseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

// parseStatement 解析语句
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// parseLetStatement let语句解析
func (p *Parser) parseLetStatement() *ast.LetStatement {
	// 语句
	stmt := &ast.LetStatement{Token: p.curToken}
	// let 后是一个标识符，如果不是代表语法错误 let a = 10;
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// 标识符
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	// let a 后是操作符
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	//TODO:表达式暂不处理
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseReturnStatement return语句解析
func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	if !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// parseExpressionStatement 解析表达式语句
func (p *Parser) parseExpressionStatement() ast.Statement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)

	// 如果下一个是分号，则前移指向分号。如果不是也没关系
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// expectPeek 检查下一个 token 的类型，确保词法单元顺序正确性，当类型正确的情况下，nextToken 移动词法单元
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t) // 收集错误
		return false
	}
}

// peekTokenIs 判断下一个token的type是否是 t
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// curTokenIs 判断当前token的type是否是 t
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// Errors 获取 Parser 收集的 error
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError 收集下一个Token与预期不符的错误
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// registerPrefixParseFn 注册前缀解析函数
func (p *Parser) registerPrefixParseFn(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// registerInfixParseFn 注册中缀解析函数
func (p *Parser) registerInfixParseFn(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseExpression(lowest int) ast.Expression {
	return nil
}
