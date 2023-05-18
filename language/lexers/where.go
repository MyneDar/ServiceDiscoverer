package lexers

import (
	"servicediscoverer/models"
)

// WhereLex is a lexer for WHERE keyword and for the filter arguments and operators
type WhereLex struct{}

// Process is a function that checks if the first word is WHERE and if it is, it returns a token with WHERE name.
// It also checks if the next words are operators or arguments, what will help us to filter.
func (l *WhereLex) Process(command *[]string) (err error, tokens []models.TokenStruct) {
	//E.g. WHERE (ASD=25) AND Aasd = 25
	//E.g.WHERE ((ASD=2.5) AND XYZ=FROM) AND wetw=54

	//We check if the first word is WHERE and if it is, we append it to the tokens slice and remove it from the command.
	//If it isn't, we return nil, tokens, because we don't need to process this command.
	if (*command)[0] == models.WHERE.String() {
		tokens = append(tokens, models.TokenStruct{Name: models.WHERE, Data: (*command)[0]})
		*command = (*command)[1:]
	} else {
		return nil, tokens
	}

	for index, word := range *command {
		//We check if the one of the next words is a keyword, if it is, we remove the processed words from the command and return the tokens.
		if models.IsKeyword(word) {
			*command = (*command)[index:]
			return nil, tokens
		}

		//We check if the word is an operator, if it is, we append it to the tokens slice
		if opIndex, ok := models.IsOperator(word); ok {
			tokens = append(tokens, models.TokenStruct{Name: opIndex, Data: word})
			continue
		}

		var emptySlice []models.TokenStruct
		//We tokenize the word with the RecRuneTokenizer function (It decides that it is a string, int, float etc.)
		// and append them to the tokens slice.
		tokens = append(tokens, models.RecRuneTokenizer(word, emptySlice)...)
	}

	//We set the command to nil, because we don't need it anymore
	*command = nil

	//We return the tokens
	return nil, tokens
}
