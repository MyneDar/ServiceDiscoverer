package interfaces

import "servicediscoverer/models"

type Lexer interface {
	Process(command *[]string) (err error, tokens []models.TokenStruct)
}
