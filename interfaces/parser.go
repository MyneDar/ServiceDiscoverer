package interfaces

import (
	"servicediscoverer/models"
)

type Parser interface {
	Process(tok []models.TokenStruct, information *map[string]interface{}) error
}
