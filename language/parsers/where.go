package parsers

import "servicediscoverer/models"

type WhereParser struct {
}

func NewWhereParser() *WhereParser {
	return &WhereParser{}
}

func (l *WhereParser) Process(tok []models.TokenStruct) error {
	return nil
}
