package service

import (
	"errors"
	"servicediscoverer/interfaces"
	"servicediscoverer/language/lexers"
	"servicediscoverer/models"
	"strings"
)

// Tokenizer is a struct that contains the lexers that are used to tokenize the input.
type Tokenizer struct {
	tokenizerLexers []interfaces.Lexer
}

// NewTokenizer creates a new Tokenizer with the default set of lexers.
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

// CommandProcess is a function that tokenizes the input command.
func (t *Tokenizer) CommandProcess(command string) (err error, tokens map[models.ServiceToken][]models.TokenStruct) {
	tokens = make(map[models.ServiceToken][]models.TokenStruct)

	// Split the command by spaces
	commands := strings.Split(command, " ")

	// Iterate over the lexers and try to tokenize the command
	for _, lexer := range t.tokenizerLexers {
		// If the command is empty, break the loop
		if len(commands) == 0 {
			break
		}

		// Process part of the command with the lexer
		err, appendableTokens := lexer.Process(&commands)
		if err != nil {
			return err, nil
		}

		// If the lexer returned tokens, add them to the tokens map
		if len(appendableTokens) != 0 {
			tokens[appendableTokens[0].Name] = appendableTokens[1:]
		}
	}

	// If the command is not empty after the lexers, return an error
	if len(commands) != 0 {
		err = errors.New("not right command list, cant recognize: " + strings.Join(commands, ""))
		return err, nil
	} else {
		return nil, tokens
	}
}
