package interfaces

import "servicediscoverer/models"

type Lexer interface {
	Process(command *string) (error, []models.ServiceToken)
}
