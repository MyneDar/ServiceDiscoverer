package parsers

import "servicediscoverer/models"

type InsertParser struct {
}

func NewInsertParser() *InsertParser {
	return &InsertParser{}
}

func (l *InsertParser) Process(tok []models.TokenStruct, information map[string]interface{}) error {
	return nil
}
