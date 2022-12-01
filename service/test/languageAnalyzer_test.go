package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"servicediscoverer/dev"
	"servicediscoverer/ent"
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
	defer func(LocalClient *ent.Client) {
		errDb := LocalClient.Close()
		assert.Nil(t, errDb, "Db close err should be nil")
	}(dev.LocalClient)

	var analyzer = service.NewLanguageAnalyzer()
	err, response := analyzer.TokenProcess(tokens, nil)

	assert.NotNil(t, err, "There should be an error")
	var notfoundError *ent.NotFoundError
	assert.True(t, errors.As(err, &notfoundError), "There should be a NotFoundError")
	assert.Nil(t, response, "Value should be nil")
}
