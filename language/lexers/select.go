package lexers

import "servicediscoverer/models"

type SelectLex struct {
}

func (l *SelectLex) Process(command *string) (error, []models.Token) {
	return nil, []models.Token{}
}
