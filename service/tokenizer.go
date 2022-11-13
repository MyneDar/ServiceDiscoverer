package service

import (
	"errors"
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
		&lexers.DeleteLex{},
		&lexers.InsertLex{},
		&lexers.UpdateLex{},
		&lexers.SelectLex{},
	}
	return &Tokenizer{tokenizerLexers: appendLexers}
}

func (t *Tokenizer) CommandProcess(command string) (err error, tokens []models.TokenStruct) {
	commands := strings.Split(command, " ")

	for _, lexer := range t.tokenizerLexers {
		if len(commands) == 0 {
			break
		}
		err, appendableTokens := lexer.Process(&commands)
		if err != nil {
			return err, nil
		}
		tokens = append(tokens, appendableTokens...)
	}
	if len(commands) != 0 {
		err = errors.New("not right command list, cant recognize: " + strings.Join(commands, ""))
		return err, nil
	} else {
		return nil, tokens
	}
}
