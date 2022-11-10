package lexers

import "servicediscoverer/models"

type FromParser struct {
}

func (l *FromParser) Process(tok []models.ServiceToken, commandMap *map[string]interface{}) (error, interface{}) {
	return nil, nil
}
