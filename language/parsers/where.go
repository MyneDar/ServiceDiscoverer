package parsers

import "servicediscoverer/models"

// WhereParser is a parser for WHERE keyword
type WhereParser struct{}

// NewWhereParser is a function that returns a pointer to a new WhereParser struct
func NewWhereParser() *WhereParser {
	return &WhereParser{}
}

// Process is a function that checks what filters should be used for the call
func (l *WhereParser) Process(tok []models.TokenStruct, information map[string]interface{}) error {
	return nil
}
