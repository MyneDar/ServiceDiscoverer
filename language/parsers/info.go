package parsers

import "servicediscoverer/models"

type InfoParser struct {
}

func NewInfoParser() *InfoParser {
	return &InfoParser{}
}

func (l *InfoParser) Process(tok []models.TokenStruct, information map[string]interface{}) error {
	return nil
}
