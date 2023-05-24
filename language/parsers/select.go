package parsers

import "servicediscoverer/models"

// SelectParser is a parser for SELECT keyword
type SelectParser struct{}

// NewSelectParser is a function that returns a pointer to a new SelectParser struct
func NewSelectParser() *SelectParser {
	return &SelectParser{}
}

// Process is a function that checks if the call has to be a GET or not and what filters should it use
func (l *SelectParser) Process(tok []models.TokenStruct, information map[string]interface{}) error {
	return nil
}
