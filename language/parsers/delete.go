package parsers

import "servicediscoverer/models"

type DeleteParser struct {
}

func NewDeleteParser() *DeleteParser {
	return &DeleteParser{}
}

func (l *DeleteParser) Process(tok []models.TokenStruct, information map[string]interface{}) error {
	return nil
}
