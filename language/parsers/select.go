package lexers

import "servicediscoverer/models"

type SelectParser struct {
}

func (l *SelectParser) Process(tok []models.ServiceToken, commandMap *map[string]interface{}) (error, interface{}) {
	return nil, nil
}
