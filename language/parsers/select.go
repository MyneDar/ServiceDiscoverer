package lexers

import "servicediscoverer/models"

type SelectParser struct {
}

func (l *SelectParser) Process(tok []models.Token, commandMap *map[string]interface{}) (error, interface{}) {
	return nil, nil
}
