package test

import (
	"github.com/stretchr/testify/assert"
	"servicediscoverer/dev"
	"servicediscoverer/models"
	"servicediscoverer/service"
	"testing"
)

//
//
// Good cases
//
//

//
//
//

//
//
// Error cases
//
//

func TestTokenProcessEmptyDatabase(t *testing.T) {
	//test initialization
	tokens := map[models.ServiceToken][]models.TokenStruct{
		models.FROM:   {{Name: models.IDENT, Data: "Human"}, {Name: models.IDENT, Data: "Add"}},
		models.DELETE: {},
	}

	//Running of the test
	err := dev.EntClientInit()
	assert.Nil(t, err, "There should be no error")

	var analyzer = service.NewLanguageAnalyzer()
	err, response := analyzer.TokenProcess(tokens, nil)

	assert.NotNil(t, err, "There should be an error")
	assert.Nil(t, response, "Value should be nil")
}
