package interfaces

import (
	"servicediscoverer/models"
)

type LanguageAnalyzer interface {
	TokenProcess(tokens map[models.ServiceToken][]models.TokenStruct, json map[string]interface{}) (err error, response []byte)
}
