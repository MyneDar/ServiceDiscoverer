package interfaces

import (
	"servicediscoverer/models"
)

type Parser interface {
	Process(tok []models.Token, commandMap *map[string]interface{}) (error, interface{})
}
