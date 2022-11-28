package parsers

import "servicediscoverer/models"

type FromParser struct {
}

func NewFromParser() *FromParser {
	return &FromParser{}
}

func (l *FromParser) Process(tok []models.TokenStruct) error {
	return nil
}
