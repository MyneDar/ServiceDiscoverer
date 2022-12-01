package test

import (
	"github.com/stretchr/testify/assert"
	"servicediscoverer/models"
	"servicediscoverer/service"
	"testing"
)

func TestGoodSchema(t *testing.T) {
	schemaMap := map[models.ServiceToken][]models.TokenStruct{
		models.FROM:   nil,
		models.SELECT: nil,
	}
	err := service.SchemaCheck(schemaMap)
	assert.Nil(t, err, "Error should be nil")
}

func TestBadSchema(t *testing.T) {
	schemaMap := map[models.ServiceToken][]models.TokenStruct{
		models.SELECT: nil,
		models.FROM:   nil,
	}
	err := service.SchemaCheck(schemaMap)
	assert.NotNil(t, err, "Error should be not nil")
}
