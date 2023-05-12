package lexers

import (
	"errors"
	"servicediscoverer/models"
	"strings"
)

// FromLex is a lexer for FROM keyword
type FromLex struct{}

// Process is a function that checks if the first word is FROM and if it is, it returns a token with FROM name
func (l *FromLex) Process(command *[]string) (err error, tokens []models.TokenStruct) {
	var split []string

	//a for loop that checks if the first word is FROM and if it is, it returns a token with FROM name
	for index, word := range *command {
		//If the word is FROM, and it is the first word, it appends it to the split slice
		if word == models.FROM.String() && index == 0 {
			split = append(split, word)
			continue
		} else if word != models.FROM.String() && index == 0 {
			break
		}

		//if the word is a keyword, it breaks the loop
		if models.IsKeyword(word) {
			break
		} else {
			split = append(split, word)
		}
	}

	//if the split slice is empty, it returns nil and the tokens slice (which is also empty) and breaks the function
	if len(split) == 0 {
		return nil, tokens
	}

	//It puts the FROM token in the tokens slice
	tokens = append(tokens, models.TokenStruct{Name: models.FROM, Data: split[0]})

	//if the split slice has only one element, it returns nil and the tokens slice and breaks the function
	if len(split) < 2 {
		err = errors.New("Bad From command")
		return err, nil
	}

	//if the target slice has two elements, it puts them in the tokens slice
	target := strings.Split(split[1], models.PERIOD.String())
	if len(target) == 2 {
		tokens = append(tokens, models.TokenStruct{Name: models.IDENT, Data: target[0]})
		tokens = append(tokens, models.TokenStruct{Name: models.IDENT, Data: target[1]})
	} else {
		err = errors.New("Bad Servicename.Endpoint squence")
		return err, nil
	}

	//It removes the processed words from the command slice
	*command = (*command)[len(split):]

	return err, tokens
}
