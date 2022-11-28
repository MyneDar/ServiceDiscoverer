package parsers

import "servicediscoverer/models"

type SelectParser struct {
}

func NewSelectParser() *SelectParser {
	return &SelectParser{}
}

func (l *SelectParser) Process(tok []models.TokenStruct, information map[string]interface{}) error {
	return nil
}
