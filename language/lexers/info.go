package lexers

import "servicediscoverer/models"

type InfoLex struct {
}

func (l *InfoLex) Process(command *string) (error, []models.ServiceToken) {
	return nil, []models.ServiceToken{}
}
