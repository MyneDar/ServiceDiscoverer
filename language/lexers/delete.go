package lexers

import "servicediscoverer/models"

// DeleteLex is a lexer for DELETE keyword
type DeleteLex struct{}

// Process is a function that checks if the first word is DELETE and if it is, it returns a token with DELETE name.
func (l *DeleteLex) Process(command *[]string) (err error, tokens []models.TokenStruct) {
	word := (*command)[0]

	if word == models.DELETE.String() {
		tokens = append(tokens, models.TokenStruct{Name: models.DELETE, Data: word})
		*command = (*command)[1:]
	}

	return nil, tokens
}
