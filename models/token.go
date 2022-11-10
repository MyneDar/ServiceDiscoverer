package models

import "strconv"

type TokenStruct struct {
	Name ServiceToken
	Data string
}

type ServiceToken int

const (
	//Special token
	ILLEGAL ServiceToken = iota
	EOF

	literalBeg
	IDENT  // like 'Human'
	INT    // 12345
	FLOAT  // 123.45
	CHAR   // 'a'
	STRING // "abc"
	literalEnd

	//Commands
	commandBeg
	FROM
	INSERT
	DELETE
	SELECT
	UPDATE
	INFO
	WHERE
	commandEnd

	operatorBeg
	EQL       // =
	LSS       // <
	GTR       // >
	AND       //and
	OR        //or
	NOT       //not
	LEQ       // <=
	GEQ       // >=
	LPAREN    // (
	RPAREN    // )
	SEMICOLON // ;
	PERIOD    // .
	ASTERISK  // *
	operatorEnd
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	IDENT:  "IDENT",
	INT:    "INT",
	FLOAT:  "FLOAT",
	CHAR:   "CHAR",
	STRING: "STRING",

	FROM:   "FROM",
	INSERT: "INSERT",
	DELETE: "DELETE",
	SELECT: "SELECT",
	UPDATE: "UPDATE",
	INFO:   "INFO",
	WHERE:  "HERE",

	EQL:       "=",
	LSS:       "<",
	GTR:       ">",
	AND:       "AND",
	OR:        "OR",
	NOT:       "NOT",
	LEQ:       "<=",
	GEQ:       ">=",
	LPAREN:    "(",
	RPAREN:    ")",
	SEMICOLON: ";",
	PERIOD:    ".",
	ASTERISK:  "*",
}

func (tok ServiceToken) String() string {
	s := ""
	if 0 <= tok && tok < ServiceToken(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

const (
	LowestPrec  = 0 // non-operators
	HighestPrec = 4
)

// Precedence returns the operator precedence of the binary
// operator op. If op is not a binary operator, the result
// is LowestPrecedence.
func (tok ServiceToken) Precedence() int {
	switch tok {
	case EQL, LSS, LEQ, GTR, GEQ:
		return 1
	case OR:
		return 2
	case AND:
		return 3
	}
	return LowestPrec
}

var keywords map[string]ServiceToken

func init() {
	keywords = make(map[string]ServiceToken, commandEnd-(commandBeg+1))
	for i := commandBeg + 1; i < commandEnd; i++ {
		keywords[tokens[i]] = i
	}
}

// Lookup maps an identifier to its keyword token or IDENT (if not a keyword).
func Lookup(ident string) ServiceToken {
	if tok, isKeyword := keywords[ident]; isKeyword {
		return tok
	}
	return IDENT
}

func (tok ServiceToken) IsLiteral() bool { return literalBeg < tok && tok < literalEnd }
func (tok ServiceToken) IsOperator() bool {
	return operatorBeg < tok && tok < operatorEnd
}

func (tok ServiceToken) IsKeyword() bool { return commandBeg < tok && tok < commandEnd }
func IsKeyword(name string) bool {
	_, ok := keywords[name]
	return ok
}

func IsIdentifier(name string) bool {
	return name == "" || IsKeyword(name)
}
