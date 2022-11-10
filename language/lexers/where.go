package lexers

import "servicediscoverer/models"

type WhereLex struct {
}

func (l *WhereLex) Process(command *string) (error, []models.ServiceToken) {
	return nil, []models.ServiceToken{}
}
