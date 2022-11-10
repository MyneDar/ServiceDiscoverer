package lexers

import "servicediscoverer/models"

type SelectLex struct {
}

func (l *SelectLex) Process(command *string) (error, []models.ServiceToken) {
	return nil, []models.ServiceToken{}
}
