package lexers

import "servicediscoverer/models"

type WhereLex struct {
}

func (l *WhereLex) Process(command *[]string) (err error, tokens []models.TokenStruct) {
	return nil, tokens
}
