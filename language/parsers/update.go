package parsers

import "servicediscoverer/models"

// UpdateParser is a parser for update keyword
type UpdateParser struct{}

// NewUpdateParser is a function that returns a pointer to a new UpdateParser struct
func NewUpdateParser() *UpdateParser {
	return &UpdateParser{}
}

// Process is a function that checks if the call has to be a PUT or not
func (l *UpdateParser) Process(tok []models.TokenStruct, information map[string]interface{}) error {
	return nil
}
