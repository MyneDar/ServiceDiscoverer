package lexers

import "servicediscoverer/models"

type DeleteLex struct {
}

func (l *DeleteLex) Process(command *string) (error, []models.Token) {
	return nil, []models.Token{}
}
