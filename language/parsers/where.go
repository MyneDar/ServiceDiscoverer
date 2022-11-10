package lexers

import "servicediscoverer/models"

type WhereParser struct {
}

func (l *WhereParser) Process(tok []models.ServiceToken, commandMap *map[string]interface{}) (error, interface{}) {
	return nil, nil
}
