package lexer

// TokenType representa o tipo de um token
type TokenType string

// Token representa um token do lexer
type Token struct {
	Type    TokenType
	Literal string
}

// Constantes para os tipos de tokens
const (
	TokenIllegal = "ILLEGAL"
	TokenEOF     = "EOF"

	// Identificadores + literais
	TokenIdent  = "IDENT"
	TokenInt    = "INT"
	TokenFloat  = "FLOAT"
	TokenString = "STRING"

	// Operadores
	TokenAssign   = "="
	TokenPlus     = "+"
	TokenMinus    = "-"
	TokenBang     = "!"
	TokenAsterisk = "*"
	TokenSlash    = "/"
	TokenLT       = "<"
	TokenGT       = ">"
	TokenEQ       = "=="
	TokenNotEQ    = "!="
	TokenColon    = ":"

	// Delimitadores
	TokenComma     = ","
	TokenSemicolon = ";"
	TokenLParen    = "("
	TokenRParen    = ")"
	TokenLBrace    = "{"
	TokenRBrace    = "}"
	TokenDot       = "."

	// Palavras-chave
	TokenFunction = "FUNCTION"
	TokenLet      = "LET"
	TokenTrue     = "TRUE"
	TokenFalse    = "FALSE"
	TokenIf       = "IF"
	TokenElse     = "ELSE"
	TokenReturn   = "RETURN"
	TokenClass    = "CLASS"
	TokenProp     = "PROP"
	TokenNew      = "NEW"
)

var keywords = map[string]TokenType{
	"fn":     TokenFunction,
	"class":  TokenClass,
	"prop":   TokenProp,
	"true":   TokenTrue,
	"false":  TokenFalse,
	"if":     TokenIf,
	"else":   TokenElse,
	"return": TokenReturn,
	"new":    TokenNew,
}

// Lexer representa o analisador léxico
type Lexer struct {
	input        string
	position     int  // posição atual no input (aponta para o caractere atual)
	readPosition int  // posição atual de leitura (após o caractere atual)
	ch           byte // caractere atual sendo examinado
}

// NewLexer cria um novo lexer
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar lê o próximo caractere e avança a posição no input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken retorna o próximo token do input
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = Token{Type: TokenEQ, Literal: literal}
		} else {
			tok = newToken(TokenAssign, l.ch)
		}
	case ':':
		tok = newToken(TokenColon, l.ch)
	case '+':
		tok = newToken(TokenPlus, l.ch)
	case '-':
		tok = newToken(TokenMinus, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = Token{Type: TokenNotEQ, Literal: literal}
		} else {
			tok = newToken(TokenBang, l.ch)
		}
	case '*':
		tok = newToken(TokenAsterisk, l.ch)
	case '/':
		tok = newToken(TokenSlash, l.ch)
	case '<':
		tok = newToken(TokenLT, l.ch)
	case '>':
		tok = newToken(TokenGT, l.ch)
	case ',':
		tok = newToken(TokenComma, l.ch)
	case ';':
		tok = newToken(TokenSemicolon, l.ch)
	case '(':
		tok = newToken(TokenLParen, l.ch)
	case ')':
		tok = newToken(TokenRParen, l.ch)
	case '{':
		tok = newToken(TokenLBrace, l.ch)
	case '}':
		tok = newToken(TokenRBrace, l.ch)
	case '.':
		tok = newToken(TokenDot, l.ch)
	case '"':
		tok.Type = TokenString
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = TokenEOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = lookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			return l.readNumber()
		} else {
			tok = newToken(TokenIllegal, l.ch)
		}
	}

	l.readChar()
	return tok
}

// skipWhitespace pula caracteres de espaço em branco
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// readIdentifier lê um identificador e retorna seu valor
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber lê um número e retorna o token apropriado
func (l *Lexer) readNumber() Token {
	position := l.position
	isFloat := false

	for isDigit(l.ch) || l.ch == '.' {
		if l.ch == '.' {
			if isFloat {
				break
			}
			isFloat = true
		}
		l.readChar()
	}

	literal := l.input[position:l.position]
	if isFloat {
		return Token{Type: TokenFloat, Literal: literal}
	}
	return Token{Type: TokenInt, Literal: literal}
}

// readString lê uma string e retorna seu valor
func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

// peekChar retorna o próximo caractere sem avançar a posição
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// newToken cria um novo token
func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

// isLetter verifica se um caractere é uma letra
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit verifica se um caractere é um dígito
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// lookupIdent verifica se um identificador é uma palavra-chave
func lookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return TokenIdent
}
