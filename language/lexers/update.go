package lexers

import "servicediscoverer/models"

// UpdateLex is a lexer for UPDATE keyword
type UpdateLex struct{}

// Process is a function that checks if the first word is UPDATE and if it is, it returns a token with UPDATE name.
func (l *UpdateLex) Process(command *[]string) (err error, tokens []models.TokenStruct) {
	word := (*command)[0]

	if word == models.UPDATE.String() {
		tokens = append(tokens, models.TokenStruct{Name: models.UPDATE, Data: word})
		*command = (*command)[1:]
	}

	return nil, tokens
}
