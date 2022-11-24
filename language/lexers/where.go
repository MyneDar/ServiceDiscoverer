package lexers

import (
	"servicediscoverer/models"
)

type WhereLex struct {
}

func (l *WhereLex) Process(command *[]string) (err error, tokens []models.TokenStruct) {
	//WHERE (ASD=25) AND Aasd = 25
	//WHERE ((ASD=2.5) AND XYZ=FROM) AND wetw=54

	if (*command)[0] == models.WHERE.String() {
		tokens = append(tokens, models.TokenStruct{Name: models.WHERE, Data: (*command)[0]})
		*command = (*command)[1:]
	} else {
		return nil, tokens
	}
	for index, word := range *command {
		if models.IsKeyword(word) {
			*command = (*command)[index:]
			return nil, tokens
		}
		if opIndex, ok := models.IsOperator(word); ok {
			tokens = append(tokens, models.TokenStruct{Name: opIndex, Data: word})
			continue
		}

		var emptySlice []models.TokenStruct
		tokens = append(tokens, models.RecRuneTokenizer(word, emptySlice)...)
	}
	*command = nil
	return nil, tokens
}
