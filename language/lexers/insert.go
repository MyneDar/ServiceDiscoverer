package lexers

import "servicediscoverer/models"

type InsertLex struct {
}

func (l *InsertLex) Process(command *string) (error, []models.ServiceToken) {
	return nil, []models.ServiceToken{}
}
