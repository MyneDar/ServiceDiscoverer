package lexers

import (
	"errors"
	"servicediscoverer/models"
	"strings"
)

type FromLex struct {
}

func (l *FromLex) Process(command *[]string) (err error, tokens []models.TokenStruct) {
	var splitted []string
	for index, word := range *command {
		if word == models.FROM.String() && index == 0 {
			splitted = append(splitted, word)
			continue
		} else if word != models.FROM.String() && index == 0 {
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

	//["FROM","service.endpoint"] 2
	tokens = append(tokens, models.TokenStruct{Name: models.FROM, Data: splitted[0]})

	target := strings.Split(splitted[1], models.PERIOD.String())
	if len(target) == 2 {
		tokens = append(tokens, models.TokenStruct{Name: models.IDENT, Data: target[0]})
		tokens = append(tokens, models.TokenStruct{Name: models.IDENT, Data: target[1]})
	} else {
		err = errors.New("Bad Servicename.Endpoint squence")
		return err, nil
	}

	return err, tokens
}
