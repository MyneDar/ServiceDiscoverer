package lexers

import (
	"errors"
	"servicediscoverer/models"
	"strings"
)

// InfoLex is a lexer for INFO keyword
type InfoLex struct{}

// Process is a function that checks if the first word is INFO and if it is,
// it returns a token with INFO name and the target information( It is tokenized (service.endpoint) or (service.*) or (*))
func (l *InfoLex) Process(command *[]string) (err error, tokens []models.TokenStruct) {
	var split []string

	for index, word := range *command {
		//If the word is INFO, and it is the first word,
		//it appends it to the split slice and continues the loop, if it is another word, it breaks the loop
		if word == models.INFO.String() && index == 0 {
			split = append(split, word)
			continue
		} else if word != models.INFO.String() && index == 0 {
			break
		}

		//if the word is a keyword, it breaks the loop and if not, it appends it to the split slice
		if models.IsKeyword(word) {
			break
		} else {
			split = append(split, word)
		}
	}

	//if the split slice is empty after the loop,
	//it returns nil and the tokens slice (which is also empty) and breaks the function
	if len(split) == 0 {
		return nil, tokens
	}

	//["INFO","service.endpoint"], ["INFO","service.*"], ["INFO","*"]
	//It puts the INFO token in the tokens slice
	tokens = append(tokens, models.TokenStruct{Name: models.INFO, Data: split[0]})

	//We try to split the second element of the split slice by PERIOD
	//(the second part should be the target of information request)
	target := strings.Split(split[1], models.PERIOD.String())

	if len(target) == 2 && target[0] != "" && target[1] != "" {
		//if the targets length is 2 and both of the elements are not empty,
		//we put the first element as an IDENT token, because is sure at this point
		tokens = append(tokens, models.TokenStruct{Name: models.IDENT, Data: target[0]})

		//if the second element is an asterisk, we put an ASTERISK token in the tokens slice if not, we put an IDENT token
		if target[1] == models.ASTERISK.String() {
			tokens = append(tokens, models.TokenStruct{Name: models.ASTERISK, Data: target[1]})
		} else {
			tokens = append(tokens, models.TokenStruct{Name: models.IDENT, Data: target[1]})
		}

		//if the targets length is 1 and the element is an asterisk, we put an ASTERISK token in the tokens slice
	} else if len(target) == 1 && target[0] == models.ASTERISK.String() {
		tokens = append(tokens, models.TokenStruct{Name: models.ASTERISK, Data: target[0]})

		//In every else case, we return an error, because this is a bad sequence
	} else {
		err = errors.New("Bad Servicename.Endpoint squence")
		return err, nil
	}

	//We remove the used elements from the command slice
	*command = (*command)[len(split):]

	//We return nil and the tokens slice
	return err, tokens
}
