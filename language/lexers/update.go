package lexers

import "servicediscoverer/models"

type UpdateLex struct {
}

func (l *UpdateLex) Process(command *[]string) (err error, tokens []models.TokenStruct) {
	word := (*command)[0]
	if word == models.UPDATE.String() {
		tokens = append(tokens, models.TokenStruct{Name: models.UPDATE, Data: word})
		*command = (*command)[1:]
	}
	return nil, tokens
}
