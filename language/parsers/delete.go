package lexers

import "servicediscoverer/models"

type DeleteParser struct {
}

func (l *DeleteParser) Process(tok []models.Token, commandMap *map[string]interface{}) (error, interface{}) {
	return nil, nil
}
