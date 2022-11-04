package lexers

import "servicediscoverer/models"

type InfoParser struct {
}

func (l *InfoParser) Process(tok []models.Token, commandMap *map[string]interface{}) (error, interface{}) {
	return nil, nil
}
