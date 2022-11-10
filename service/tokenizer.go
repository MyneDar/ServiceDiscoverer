package service

import (
	"servicediscoverer/interfaces"
	"servicediscoverer/language/lexers"
	"servicediscoverer/models"
	"strings"
)

type Tokenizer struct {
	tokenizerLexers []interfaces.Lexer
}

func NewTokenizer() *Tokenizer {
	var appendLexers = []interfaces.Lexer{
		&lexers.FromLex{},
		&lexers.InfoLex{},
	}
	return &Tokenizer{tokenizerLexers: appendLexers}
}

func (t *Tokenizer) CommandProcess(command string) (err error, tokens []models.TokenStruct) {
	commands := strings.Split(command, " ")

	for _, lexer := range t.tokenizerLexers {
		err, appendableTokens := lexer.Process(&commands)
		if err != nil {
			return err, nil
		}
		tokens = append(tokens, appendableTokens...)
	}
	return nil, tokens
}
