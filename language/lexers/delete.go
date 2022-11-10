package lexers

import "servicediscoverer/models"

type DeleteLex struct {
}

func (l *DeleteLex) Process(command *string) (error, []models.ServiceToken) {
	return nil, []models.ServiceToken{}
}
