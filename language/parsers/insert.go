package parsers

import "servicediscoverer/models"

// InsertParser is a parser for INSERT keyword
type InsertParser struct{}

// NewInsertParser is a function that returns a pointer to a new InsertParser struct
func NewInsertParser() *InsertParser {
	return &InsertParser{}
}

// Process is a function that checks if the call has to be a POST or not
func (l *InsertParser) Process(tok []models.TokenStruct, information map[string]interface{}) error {
	return nil
}
