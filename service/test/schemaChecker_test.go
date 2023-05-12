package test

import (
	"github.com/stretchr/testify/assert"
	"log"
	"servicediscoverer/models"
	"servicediscoverer/service"
	"strconv"
	"testing"
)

//
//
// Good cases
//
//

// goodSchemas is an array of good schemas, where the errors should be always nil.
var goodSchemas = []map[models.ServiceToken][]models.TokenStruct{
	{
		models.FROM:   nil,
		models.SELECT: nil,
	},
	{
		models.FROM:   nil,
		models.SELECT: nil,
	},
	{
		models.FROM:   nil,
		models.SELECT: nil,
	},
}

func TestSchemaCheckGoodFromSelectSchema(t *testing.T) {
	for i := 0; i < len(goodSchemas); i++ {
		err := service.SchemaCheck(goodSchemas[i])
		assert.Nil(t, err, "Error should be nil")
		log.Println("test " + strconv.Itoa(i))
	}
}

//
//
// Error cases
//
//

// badSchemas is an array of bad schemas, where always should be an error
var badSchemas = []map[models.ServiceToken][]models.TokenStruct{
	{
		models.SELECT: nil,
		models.FROM:   nil,
	},
	{
		models.SELECT: nil,
		models.FROM:   nil,
	},
	{
		models.SELECT: nil,
		models.FROM:   nil,
	},
}

func TestSchemaCheckBadSchemas(t *testing.T) {
	for i := 0; i < len(badSchemas); i++ {
		err := service.SchemaCheck(badSchemas[i])
		assert.NotNil(t, err, "Error should be not nil")
		log.Println("test " + strconv.Itoa(i))
	}

}

//
//
//
