package parser

import (
	"bytes"
	"fmt"
	"jotlango/internal/lexer"
	"strconv"
	"strings"
)

// Node representa um nó na AST
type Node interface {
	TokenLiteral() string
}

// Statement representa uma declaração
type Statement interface {
	Node
	statementNode()
	String() string
}

// Expression representa uma expressão
type Expression interface {
	Node
	expressionNode()
	String() string
}

// Program representa o programa completo
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out string
	for _, s := range p.Statements {
		out += s.String()
	}
	return out
}

// ExpressionStatement representa uma expressão como declaração
type ExpressionStatement struct {
	Token      lexer.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// IntegerLiteral representa um número inteiro
type IntegerLiteral struct {
	Token lexer.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return strconv.FormatInt(il.Value, 10) }

// FloatLiteral representa um número decimal
type FloatLiteral struct {
	Token lexer.Token
	Value float64
}

func (fl *FloatLiteral) expressionNode()      {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FloatLiteral) String() string       { return strconv.FormatFloat(fl.Value, 'f', -1, 64) }

// StringLiteral representa uma string
type StringLiteral struct {
	Token lexer.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Value }

// Boolean representa um valor booleano
type Boolean struct {
	Token lexer.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return fmt.Sprintf("%t", b.Value) }

// PrefixExpression representa uma expressão prefixada
type PrefixExpression struct {
	Token    lexer.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

// InfixExpression representa uma expressão infixada
type InfixExpression struct {
	Token    lexer.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

// BlockStatement representa um bloco de código
type BlockStatement struct {
	Token      lexer.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// IfExpression representa uma expressão condicional
type IfExpression struct {
	Token       lexer.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

// ReturnStatement representa uma declaração de retorno
type ReturnStatement struct {
	Token       lexer.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// VarStatement representa uma declaração de variável com tipo
type VarStatement struct {
	Token lexer.Token // o token 'var'
	Name  *Identifier
	Type  *Identifier
	Value Expression
}

func (vs *VarStatement) statementNode()       {}
func (vs *VarStatement) TokenLiteral() string { return vs.Token.Literal }
func (vs *VarStatement) String() string {
	var out bytes.Buffer
	out.WriteString(vs.TokenLiteral() + " ")
	out.WriteString(vs.Name.String())
	if vs.Type != nil {
		out.WriteString(": " + vs.Type.String())
	}
	if vs.Value != nil {
		out.WriteString(" = ")
		out.WriteString(vs.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// Identifier representa um identificador
type Identifier struct {
	Token lexer.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// FunctionLiteral representa uma função
type FunctionLiteral struct {
	Token      lexer.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}

// CallExpression representa uma chamada de função
type CallExpression struct {
	Token     lexer.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

// ClassStatement representa uma declaração de classe
type ClassStatement struct {
	Token lexer.Token
	Name  *Identifier
	Body  *BlockStatement
}

func (cs *ClassStatement) expressionNode()      {}
func (cs *ClassStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *ClassStatement) String() string {
	var out bytes.Buffer
	out.WriteString("class ")
	out.WriteString(cs.Name.String())
	out.WriteString(" {")
	out.WriteString(cs.Body.String())
	out.WriteString("}")
	return out.String()
}

// CallStatement representa uma chamada de função
type CallStatement struct {
	Token    lexer.Token
	Function Expression
}

func (cs *CallStatement) expressionNode()      {}
func (cs *CallStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *CallStatement) String() string {
	var out bytes.Buffer
	out.WriteString("call ")
	out.WriteString(cs.Function.String())
	return out.String()
}

// PropertyStatement representa uma declaração de propriedade
type PropertyStatement struct {
	Token lexer.Token
	Name  *Identifier
	Type  *Identifier
}

func (ps *PropertyStatement) statementNode()       {}
func (ps *PropertyStatement) TokenLiteral() string { return ps.Token.Literal }
func (ps *PropertyStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ps.TokenLiteral() + " ")
	out.WriteString(ps.Name.String())
	out.WriteString(": ")
	out.WriteString(ps.Type.String())
	out.WriteString(";")
	return out.String()
}

// FunctionStatement representa uma declaração de função
type FunctionStatement struct {
	Token      lexer.Token
	Name       *Identifier
	Parameters []*Identifier
	ReturnType *Identifier
	Body       *BlockStatement
}

func (fs *FunctionStatement) statementNode()       {}
func (fs *FunctionStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *FunctionStatement) String() string {
	var out bytes.Buffer
	out.WriteString("fn ")
	out.WriteString(fs.Name.String())
	out.WriteString("(")
	for i, param := range fs.Parameters {
		out.WriteString(param.String())
		if i < len(fs.Parameters)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString("): ")
	out.WriteString(fs.ReturnType.String())
	out.WriteString(" ")
	out.WriteString(fs.Body.String())
	return out.String()
}

// InstanceLiteral representa uma instância de classe
type InstanceLiteral struct {
	Token lexer.Token
	Class Expression
}

func (il *InstanceLiteral) expressionNode()      {}
func (il *InstanceLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *InstanceLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("new ")
	out.WriteString(il.Class.String())
	return out.String()
}

// PropertyAccess representa um acesso a propriedade
type PropertyAccess struct {
	Token    lexer.Token
	Object   Expression
	Property *Identifier
}

func (pa *PropertyAccess) expressionNode()      {}
func (pa *PropertyAccess) TokenLiteral() string { return pa.Token.Literal }
func (pa *PropertyAccess) String() string {
	var out bytes.Buffer
	out.WriteString(pa.Object.String())
	out.WriteString(".")
	out.WriteString(pa.Property.String())
	return out.String()
}

// NewExpression representa uma expressão de instanciação de classe
type NewExpression struct {
	Token     lexer.Token  // o token 'new'
	ClassName string       // nome da classe a ser instanciada
	Arguments []Expression // argumentos do construtor
}

func (ne *NewExpression) expressionNode()      {}
func (ne *NewExpression) TokenLiteral() string { return ne.Token.Literal }
func (ne *NewExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ne.Arguments {
		args = append(args, a.String())
	}

	out.WriteString("new ")
	out.WriteString(ne.ClassName)
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

// AssignmentExpression representa uma atribuição
type AssignmentExpression struct {
	Token lexer.Token
	Left  Expression
	Value Expression
}

func (ae *AssignmentExpression) expressionNode()      {}
func (ae *AssignmentExpression) TokenLiteral() string { return ae.Token.Literal }
func (ae *AssignmentExpression) String() string {
	var out bytes.Buffer
	out.WriteString(ae.Left.String())
	out.WriteString(" = ")
	out.WriteString(ae.Value.String())
	return out.String()
}

// Parser representa o analisador sintático
type Parser struct {
	l *lexer.Lexer

	curToken  lexer.Token
	peekToken lexer.Token

	prefixParseFns map[lexer.TokenType]prefixParseFn
	infixParseFns  map[lexer.TokenType]infixParseFn

	errors []string
}

type prefixParseFn func() Expression
type infixParseFn func(Expression) Expression

// Precedências dos operadores
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
	INDEX       // array[index]
)

var precedences = map[lexer.TokenType]int{
	lexer.TokenEQ:       EQUALS,
	lexer.TokenNotEQ:    EQUALS,
	lexer.TokenLT:       LESSGREATER,
	lexer.TokenGT:       LESSGREATER,
	lexer.TokenPlus:     SUM,
	lexer.TokenMinus:    SUM,
	lexer.TokenSlash:    PRODUCT,
	lexer.TokenAsterisk: PRODUCT,
	lexer.TokenLParen:   CALL,
	lexer.TokenDot:      CALL,
}

// NewParser cria um novo parser
func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:              l,
		prefixParseFns: make(map[lexer.TokenType]prefixParseFn),
		infixParseFns:  make(map[lexer.TokenType]infixParseFn),
		errors:         []string{},
	}

	// Registrar funções de parsing de prefixo
	p.registerPrefix(lexer.TokenIdent, p.parseIdentifier)
	p.registerPrefix(lexer.TokenString, p.parseStringLiteral)
	p.registerPrefix(lexer.TokenInt, p.parseIntegerLiteral)
	p.registerPrefix(lexer.TokenFloat, p.parseFloatLiteral)
	p.registerPrefix(lexer.TokenTrue, p.parseBooleanLiteral)
	p.registerPrefix(lexer.TokenFalse, p.parseBooleanLiteral)
	p.registerPrefix(lexer.TokenFunction, p.parseFunctionLiteral)
	p.registerPrefix(lexer.TokenClass, p.parseClassStatement)
	p.registerPrefix(lexer.TokenCall, p.parseCallStatement)
	p.registerPrefix(lexer.TokenPrint, p.parsePrintStatement)
	p.registerPrefix(lexer.TokenLParen, p.parseGroupedExpression)

	// Registrar funções de parsing de infix
	p.registerInfix(lexer.TokenPlus, p.parseInfixExpression)
	p.registerInfix(lexer.TokenMinus, p.parseInfixExpression)
	p.registerInfix(lexer.TokenSlash, p.parseInfixExpression)
	p.registerInfix(lexer.TokenAsterisk, p.parseInfixExpression)
	p.registerInfix(lexer.TokenEQ, p.parseInfixExpression)
	p.registerInfix(lexer.TokenNotEQ, p.parseInfixExpression)
	p.registerInfix(lexer.TokenLT, p.parseInfixExpression)
	p.registerInfix(lexer.TokenGT, p.parseInfixExpression)
	p.registerInfix(lexer.TokenDot, p.parsePropertyAccess)
	p.registerInfix(lexer.TokenLParen, p.parseCallExpression)

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()

	// Ignorar comentários
	if p.curToken.Type == lexer.TokenSlash && p.peekToken.Type == lexer.TokenSlash {
		for !p.curTokenIs(lexer.TokenEOF) && !p.curTokenIs(lexer.TokenNewLine) {
			p.nextToken()
		}
	}
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{
		Statements: []Statement{},
	}

	for p.curToken.Type != lexer.TokenEOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case lexer.TokenVar:
		return p.parseVarStatement()
	case lexer.TokenReturn:
		return p.parseReturnStatement()
	case lexer.TokenClass:
		stmt := p.parseClassStatement()
		if stmt == nil {
			return nil
		}
		return &ExpressionStatement{Token: p.curToken, Expression: stmt}
	case lexer.TokenCall:
		stmt := p.parseCallStatement()
		if stmt == nil {
			return nil
		}
		return &ExpressionStatement{Token: p.curToken, Expression: stmt}
	case lexer.TokenFunction:
		return p.parseFunctionStatement()
	case lexer.TokenPrint:
		stmt := p.parsePrintStatement()
		if stmt == nil {
			return nil
		}
		return &ExpressionStatement{Token: p.curToken, Expression: stmt}
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseClassStatement() Expression {
	stmt := &ClassStatement{
		Token: p.curToken,
		Body:  &BlockStatement{Token: p.curToken},
	}

	if !p.expectPeek(lexer.TokenIdent) {
		return nil
	}

	stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(lexer.TokenLBrace) {
		return nil
	}

	p.nextToken()

	for !p.curTokenIs(lexer.TokenRBrace) && !p.curTokenIs(lexer.TokenEOF) {
		if p.curTokenIs(lexer.TokenSlash) && p.peekTokenIs(lexer.TokenSlash) {
			// Consumir o comentário
			for !p.curTokenIs(lexer.TokenEOF) && !p.curTokenIs(lexer.TokenSemicolon) && !p.curTokenIs(lexer.TokenNewLine) {
				p.nextToken()
			}
			p.nextToken()
			continue
		}

		if p.curTokenIs(lexer.TokenFunction) {
			funcStmt := p.parseFunctionStatement()
			if funcStmt != nil {
				stmt.Body.Statements = append(stmt.Body.Statements, funcStmt)
			}
		} else {
			statement := p.parseStatement()
			if statement != nil {
				stmt.Body.Statements = append(stmt.Body.Statements, statement)
			}
		}
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseIdentifier() Expression {
	return &Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) curTokenIs(t lexer.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t lexer.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t lexer.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) registerPrefix(tokenType lexer.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType lexer.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) noPrefixParseFnError(t lexer.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) peekError(t lexer.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseNewExpression() Expression {
	expression := &NewExpression{Token: p.curToken}

	if !p.expectPeek(lexer.TokenIdent) {
		return nil
	}

	expression.ClassName = p.curToken.Literal

	if !p.expectPeek(lexer.TokenLParen) {
		return nil
	}

	expression.Arguments = p.parseExpressionList(lexer.TokenRParen)

	return expression
}

func (p *Parser) parseExpression(precedence int) Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(lexer.TokenSemicolon) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseExpressionList(end lexer.TokenType) []Expression {
	var list []Expression

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	list = append(list, p.parseExpression(LOWEST))

	for p.peekTokenIs(lexer.TokenComma) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}

func (p *Parser) parseVarStatement() *VarStatement {
	stmt := &VarStatement{Token: p.curToken}

	if !p.expectPeek(lexer.TokenIdent) {
		return nil
	}

	stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if p.peekTokenIs(lexer.TokenColon) {
		p.nextToken() // consume :
		p.nextToken() // move to type
		stmt.Type = &Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	if p.peekTokenIs(lexer.TokenAssign) {
		p.nextToken() // consume =
		p.nextToken() // move to value
		stmt.Value = p.parseExpression(LOWEST)
	}

	if p.peekTokenIs(lexer.TokenSemicolon) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ReturnStatement {
	stmt := &ReturnStatement{Token: p.curToken}

	if !p.expectPeek(lexer.TokenIdent) {
		return nil
	}

	stmt.ReturnValue = p.parseExpression(LOWEST)

	if !p.expectPeek(lexer.TokenSemicolon) {
		return nil
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() Statement {
	stmt := &ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(lexer.TokenSemicolon) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseGroupedExpression() Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(lexer.TokenRParen) {
		return nil
	}

	return exp
}

func (p *Parser) parseStringLiteral() Expression {
	return &StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() Expression {
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit := &IntegerLiteral{
		Token: p.curToken,
		Value: value,
	}

	return lit
}

func (p *Parser) parseFloatLiteral() Expression {
	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit := &FloatLiteral{
		Token: p.curToken,
		Value: value,
	}

	return lit
}

func (p *Parser) parseBooleanLiteral() Expression {
	return &BooleanLiteral{Token: p.curToken, Value: p.curToken.Type == lexer.TokenTrue}
}

func (p *Parser) parseFunctionLiteral() Expression {
	lit := &FunctionLiteral{Token: p.curToken}

	if !p.expectPeek(lexer.TokenLParen) {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(lexer.TokenLBrace) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

func (p *Parser) parseInfixExpression(left Expression) Expression {
	expression := &InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

func (p *Parser) parseCallExpression(function Expression) Expression {
	exp := &CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseExpressionList(lexer.TokenRParen)
	return exp
}

type BooleanLiteral struct {
	Token lexer.Token
	Value bool
}

func (bl *BooleanLiteral) expressionNode()      {}
func (bl *BooleanLiteral) TokenLiteral() string { return bl.Token.Literal }
func (bl *BooleanLiteral) String() string       { return strconv.FormatBool(bl.Value) }

func (p *Parser) parsePropertyAccess(object Expression) Expression {
	exp := &PropertyAccess{
		Token:  p.curToken,
		Object: object,
	}

	p.nextToken()
	exp.Property = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	return exp
}

func (p *Parser) parseFunctionParameters() []*Identifier {
	identifiers := []*Identifier{}

	if p.peekTokenIs(lexer.TokenRParen) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()
	ident := &Identifier{Token: p.curToken, Value: p.curToken.Literal}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(lexer.TokenComma) {
		p.nextToken()
		p.nextToken()
		ident := &Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(lexer.TokenRParen) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseBlockStatement() *BlockStatement {
	block := &BlockStatement{Token: p.curToken}
	block.Statements = []Statement{}

	p.nextToken()

	for !p.curTokenIs(lexer.TokenRBrace) && !p.curTokenIs(lexer.TokenEOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

func (p *Parser) parseCallStatement() Expression {
	stmt := &CallStatement{
		Token: p.curToken,
	}

	if !p.expectPeek(lexer.TokenIdent) {
		return nil
	}

	stmt.Function = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	return stmt
}

func (p *Parser) parseFunctionStatement() Statement {
	stmt := &FunctionStatement{
		Token: p.curToken,
	}

	if !p.expectPeek(lexer.TokenIdent) {
		return nil
	}

	stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(lexer.TokenLParen) {
		return nil
	}

	stmt.Parameters = p.parseFunctionParameters()

	if p.peekTokenIs(lexer.TokenColon) {
		p.nextToken() // consume :
		p.nextToken() // move to return type
		stmt.ReturnType = &Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	if !p.expectPeek(lexer.TokenLBrace) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

func (p *Parser) parsePrintStatement() Expression {
	token := p.curToken

	if !p.expectPeek(lexer.TokenLParen) {
		return nil
	}

	args := p.parseExpressionList(lexer.TokenRParen)
	if args == nil {
		return nil
	}

	return &CallExpression{
		Token:     token,
		Function:  &Identifier{Token: token, Value: "print"},
		Arguments: args,
	}
}
