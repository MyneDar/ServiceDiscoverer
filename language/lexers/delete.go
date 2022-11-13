package lexers

import "servicediscoverer/models"

type DeleteLex struct {
}

func (l *DeleteLex) Process(command *[]string) (err error, tokens []models.TokenStruct) {
	word := (*command)[0]
	if word == models.DELETE.String() {
		tokens = append(tokens, models.TokenStruct{Name: models.DELETE, Data: word})
		*command = (*command)[1:]
	}
	return nil, tokens
}
