package parsers

import "servicediscoverer/models"

type UpdateParser struct {
}

func NewUpdateParser() *UpdateParser {
	return &UpdateParser{}
}

func (l *UpdateParser) Process(tok []models.TokenStruct, information map[string]interface{}) error {
	return nil
}
