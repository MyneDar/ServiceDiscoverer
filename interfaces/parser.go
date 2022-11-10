package interfaces

import (
	"servicediscoverer/models"
)

type Parser interface {
	Process(tok []models.ServiceToken, commandMap *map[string]interface{}) (error, interface{})
}
