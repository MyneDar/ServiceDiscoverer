package lexers

import (
	"errors"
	"servicediscoverer/models"
	"strings"
)

type InfoLex struct {
}

func (l *InfoLex) Process(command *[]string) (err error, tokens []models.TokenStruct) {
	var splitted []string
	for index, word := range *command {
		if word == models.INFO.String() && index == 0 {
			splitted = append(splitted, word)
			continue
		} else if word != models.INFO.String() && index == 0 {
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

	*command = (*command)[len(splitted):]

	//["INFO","service.endpoint"], ["INFO","service.*"], ["INFO","*"]
	tokens = append(tokens, models.TokenStruct{Name: models.FROM, Data: splitted[0]})

	target := strings.Split(splitted[1], models.PERIOD.String())
	if len(target) == 2 {
		tokens = append(tokens, models.TokenStruct{Name: models.IDENT, Data: target[0]})

		if target[1] == models.ASTERISK.String() {
			tokens = append(tokens, models.TokenStruct{Name: models.ASTERISK, Data: target[1]})
		} else {
			tokens = append(tokens, models.TokenStruct{Name: models.IDENT, Data: target[1]})
		}

	} else if len(target) == 1 && target[0] == models.ASTERISK.String() {
		tokens = append(tokens, models.TokenStruct{Name: models.ASTERISK, Data: target[0]})
	} else {
		err = errors.New("Bad Servicename.Endpoint squence")
		return err, nil
	}

	return err, tokens
}
