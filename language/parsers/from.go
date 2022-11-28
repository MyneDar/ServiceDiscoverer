package parsers

import "servicediscoverer/models"

type FromParser struct {
}

func NewFromParser() *FromParser {
	return &FromParser{}
}

func (l *FromParser) Process(tok []models.TokenStruct, information *map[string]interface{}) error {
	return nil
}
