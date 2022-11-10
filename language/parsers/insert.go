package lexers

import "servicediscoverer/models"

type InsertParser struct {
}

func (l *InsertParser) Process(tok []models.ServiceToken, commandMap *map[string]interface{}) (error, interface{}) {
	return nil, nil
}
