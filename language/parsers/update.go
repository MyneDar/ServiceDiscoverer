package lexers

import "servicediscoverer/models"

type UpdateParser struct {
}

func (l *UpdateParser) Process(tok []models.ServiceToken, commandMap *map[string]interface{}) (error, interface{}) {
	return nil, nil
}
