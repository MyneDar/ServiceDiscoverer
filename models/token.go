package models

import (
	"strconv"
	"strings"
)

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

	oneRuneOperatorBeg
	EQL       // =
	LSS       // <
	GTR       // >
	LPAREN    // (
	RPAREN    // )
	SEMICOLON // ;
	PERIOD    // .
	ASTERISK  // *
	oneRuneOperatorEnd

	multipleRuneOperatorBeg
	AND //and
	OR  //or
	NOT //not
	LEQ // <=
	GEQ // >=
	MultipleRuneOperatorEnd
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
	LPAREN:    "(",
	RPAREN:    ")",
	SEMICOLON: ";",
	PERIOD:    ".",
	ASTERISK:  "*",

	AND: "AND",
	OR:  "OR",
	NOT: "NOT",
	LEQ: "<=",
	GEQ: ">=",
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
	LowestPrec  = 0 // non-oneRuneOperators
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
var oneRuneOperators map[string]ServiceToken
var multipleRuneOperators map[string]ServiceToken

func init() {
	keywords = make(map[string]ServiceToken, commandEnd-(commandBeg+1))
	for i := commandBeg + 1; i < commandEnd; i++ {
		keywords[tokens[i]] = i
	}

	for i := oneRuneOperatorBeg + 1; i < oneRuneOperatorEnd; i++ {
		oneRuneOperators[tokens[i]] = i
	}

	for i := multipleRuneOperatorBeg + 1; i < MultipleRuneOperatorEnd; i++ {
		multipleRuneOperators[tokens[i]] = i
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
	return oneRuneOperatorBeg < tok && tok < MultipleRuneOperatorEnd
}
func IsOperator(word string) (ServiceToken, bool) {
	if indOne, okOne := oneRuneOperators[word]; okOne {
		return indOne, okOne
	}
	if indMult, okMult := multipleRuneOperators[word]; okMult {
		return indMult, okMult
	}
	return ILLEGAL, false
}

func IsOneRuneOperator(r rune) (ServiceToken, bool) {
	index, ok := oneRuneOperators[string(r)]
	return index, ok
}

func (tok ServiceToken) IsKeyword() bool { return commandBeg < tok && tok < commandEnd }
func IsKeyword(name string) bool {
	_, ok := keywords[name]
	return ok
}

func RuneTokenizer(r rune) *TokenStruct {
	if ind, ok := IsOneRuneOperator(r); ok {
		return &TokenStruct{Name: ind, Data: string(r)}
	}
	return nil
}

func RecRuneTokenizer(word string, tokenStruct []TokenStruct) []TokenStruct {
	if len(word) == 0 {
		return tokenStruct
	}
	token := RuneTokenizer(rune(word[0]))
	var newTokenStruct []TokenStruct
	if token != nil {
		newTokenStruct = append(tokenStruct, *token)
		newWord := word[1:]
		return RecRuneTokenizer(newWord, newTokenStruct)
	}
	return RecRuneTokenizerIndex(0, word, tokenStruct)
}

func RecRuneTokenizerIndex(index int, word string, tokenStruct []TokenStruct) []TokenStruct {
	// Where ASD='ASD'
	if len(word) == 0 {
		return tokenStruct
	}
	if index > len(word) {
		token := CheckWord(word)
		return append(tokenStruct, token)
	}

	if index, ok := IsOneRuneOperator(rune(word[index])); ok {
		subWord := word[:index]
		token := CheckWord(subWord)
		newTokenStruct := append(tokenStruct, token, TokenStruct{Name: index, Data: string(word[index])})
		newWord := word[len(subWord)+1:]
		return RecRuneTokenizer(newWord, newTokenStruct)
	}
	return RecRuneTokenizerIndex(index+1, word, tokenStruct)

}

func CheckWord(word string) TokenStruct {
	if subWord := strings.Split(word, "."); len(subWord) == 2 {
		_, errFirst := strconv.Atoi(subWord[0])
		_, errSec := strconv.Atoi(subWord[1])
		if errFirst != nil && errSec != nil {
			return TokenStruct{Name: FLOAT, Data: word}
		} else {
			return TokenStruct{Name: IDENT, Data: word}
		}
	}

	if _, err := strconv.Atoi(word); err == nil {
		return TokenStruct{Name: INT, Data: word}
	}

	if len(word) == 3 && word[0] == '\'' && word[2] == '\'' {
		return TokenStruct{Name: CHAR, Data: word}
	}

	if word[0] == '"' && word[len(word)-1] == '"' {
		return TokenStruct{Name: STRING, Data: word}
	}
	return TokenStruct{Name: IDENT, Data: word}

}

func IsIdentifier(name string) bool {
	return name == "" || IsKeyword(name)
}
