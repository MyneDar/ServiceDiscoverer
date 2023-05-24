package parsers

import "servicediscoverer/models"

// DeleteParser is a parser for DELETE keyword
type DeleteParser struct{}

// NewDeleteParser is a function that returns a pointer to a new DeleteParser struct
func NewDeleteParser() *DeleteParser {
	return &DeleteParser{}
}

// Process is a function that checks if the call has to DELETE or not
func (l *DeleteParser) Process(tok []models.TokenStruct, information map[string]interface{}) error {
	return nil
}
