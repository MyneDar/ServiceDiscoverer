package interfaces

import "servicediscoverer/models"

type Tokenizer interface {
	CommandProcess(command string) (err error, tokens []models.TokenStruct)
}
