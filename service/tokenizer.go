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
		&lexers.WhereLex{},
	}
	return &Tokenizer{tokenizerLexers: appendLexers}
}

func (t *Tokenizer) CommandProcess(command string) (err error, tokens map[models.ServiceToken][]models.TokenStruct) {
	tokens = make(map[models.ServiceToken][]models.TokenStruct)

	commands := strings.Split(command, " ")

	for _, lexer := range t.tokenizerLexers {
		if len(commands) == 0 {
			break
		}
		err, appendableTokens := lexer.Process(&commands)
		if err != nil {
			return err, nil
		}
		if len(appendableTokens) != 0 {
			tokens[appendableTokens[0].Name] = appendableTokens[1:]
		}
	}
	if len(commands) != 0 {
		err = errors.New("not right command list, cant recognize: " + strings.Join(commands, ""))
		return err, nil
	} else {
		return nil, tokens
	}
}
