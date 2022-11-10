package lexers

import "servicediscoverer/models"

type UpdateLex struct {
}

func (l *UpdateLex) Process(command *string) (error, []models.ServiceToken) {
	return nil, []models.ServiceToken{}
}
