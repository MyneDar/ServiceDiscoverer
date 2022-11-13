package lexers

import (
	"errors"
	"golang.org/x/exp/slices"
	"servicediscoverer/models"
)

type SelectLex struct {
}

func (l *SelectLex) Process(command *[]string) (err error, tokens []models.TokenStruct) {
	var splitted []string
	for index, word := range *command {
		if word == models.SELECT.String() && index == 0 {
			splitted = append(splitted, word)
			continue
		} else if word != models.SELECT.String() && index == 0 {
			break
		}

		if models.IsKeyword(word) {
			break
		} else {
			splitted = append(splitted, word)
		}
	}
	if len(splitted) == 0 {
		return nil, tokens
	}
	containAsterisk := slices.Contains(splitted, models.ASTERISK.String())
	if len(splitted) > 2 && containAsterisk {
		err = errors.New("incorrect Select syntax: asterix plus more argument")
		return err, nil
	}
	if len(splitted) == 1 {
		err = errors.New("incorrect Select syntax: Only SELECT keyword exist")
		return err, nil
	}

	*command = (*command)[len(splitted):]

	for _, word := range splitted {
		if word == models.SELECT.String() {
			tokens = append(tokens, models.TokenStruct{Name: models.SELECT, Data: word})
		} else if word == models.ASTERISK.String() {
			tokens = append(tokens, models.TokenStruct{Name: models.ASTERISK, Data: word})
		} else {
			tokens = append(tokens, models.TokenStruct{Name: models.IDENT, Data: word})
		}
	}
	return nil, tokens
}
