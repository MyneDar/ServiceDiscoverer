package lexers

import "servicediscoverer/models"

type FromLex struct {
}

func (l *FromLex) Process(command *string) (error, []models.Token) {
	return nil, []models.Token{}
}
