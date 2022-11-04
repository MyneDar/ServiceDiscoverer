package lexers

import "servicediscoverer/models"

type WhereParser struct {
}

func (l *WhereParser) Process(tok []models.Token, commandMap *map[string]interface{}) (error, interface{}) {
	return nil, nil
}
