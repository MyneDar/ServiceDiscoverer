package lexers

import "servicediscoverer/models"

// InsertLex is a lexer for INSERT keyword
type InsertLex struct{}

// Process is a function that checks if the first word is INSERT and if it is, it returns a token with INSERT name.
func (l *InsertLex) Process(command *[]string) (err error, tokens []models.TokenStruct) {
	word := (*command)[0]

	if word == models.INSERT.String() {
		tokens = append(tokens, models.TokenStruct{Name: models.INSERT, Data: word})
		*command = (*command)[1:]
	}

	return nil, tokens
}
