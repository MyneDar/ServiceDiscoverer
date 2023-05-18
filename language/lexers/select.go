package lexers

import (
	"errors"
	"golang.org/x/exp/slices"
	"servicediscoverer/models"
)

// SelectLex is a lexer for SELECT keyword and for the arguments of the query (which data to select)
type SelectLex struct{}

// Process is a function that checks if the first word is SELECT and get the arguments of the query
func (l *SelectLex) Process(command *[]string) (err error, tokens []models.TokenStruct) {
	var split []string

	for index, word := range *command {
		//if the first word is SELECT, add it to the split slice. If it is not, break the loop
		if word == models.SELECT.String() && index == 0 {
			split = append(split, word)
			continue
		} else if word != models.SELECT.String() && index == 0 {
			break
		}

		//At the next keywords, break the loop, otherwise add the word to the split slice
		if models.IsKeyword(word) {
			break
		} else {
			split = append(split, word)
		}
	}

	//If the split slice is empty we end the function without error
	if len(split) == 0 {
		return nil, tokens
	}

	// we check that the split slice contains asterisk or not
	containAsterisk := slices.Contains(split, models.ASTERISK.String())

	//if there is an asterisk in the split slice that should be the only argument, so if there is more we return an error
	if len(split) > 2 && containAsterisk {
		err = errors.New("incorrect Select syntax: asterix plus more argument")
		return err, nil
	}

	// if there is only the SELECT keyword, we return an error
	if len(split) == 1 {
		err = errors.New("incorrect Select syntax: Only SELECT keyword exist")
		return err, nil
	}

	//We remove everything from the command slice that is in the split slice
	*command = (*command)[len(split):]

	//We tokenize the split slice
	for _, word := range split {

		if word == models.SELECT.String() {
			tokens = append(tokens, models.TokenStruct{Name: models.SELECT, Data: word})
		} else if word == models.ASTERISK.String() {
			tokens = append(tokens, models.TokenStruct{Name: models.ASTERISK, Data: word})
		} else {
			tokens = append(tokens, models.TokenStruct{Name: models.IDENT, Data: word})
		}

	}

	//We return nil error and the tokens slice
	return nil, tokens
}
