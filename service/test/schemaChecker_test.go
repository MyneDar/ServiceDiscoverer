package test

import (
	"github.com/stretchr/testify/assert"
	"log"
	"servicediscoverer/models"
	"servicediscoverer/service"
	"strconv"
	"testing"
)

// generateGoodSchemas generates good schemas for testing.
func generateGoodSchemas() []map[models.ServiceToken][]models.TokenStruct {
	goodStepLists := service.GenerateGoodStepLists()
	var goodSchemas []map[models.ServiceToken][]models.TokenStruct

	for _, goodStepList := range goodStepLists {
		goodSchema := make(map[models.ServiceToken][]models.TokenStruct)
		for _, goodStep := range goodStepList {
			goodSchema[goodStep] = nil
		}

		goodSchemas = append(goodSchemas, goodSchema)
	}

	log.Println("goodSchemas: ", goodSchemas)

	return goodSchemas
}

//
//
// Good cases
//
//

// TestSchemaCheckGoodFromSelectSchema tests good schemas.
func TestSchemaCheckGoodFromSelectSchema(t *testing.T) {
	goodSchemas := generateGoodSchemas()

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

// generateBadSchemas generates bad schemas for testing.
func generateBadSchemas() []map[models.ServiceToken][]models.TokenStruct {
	badStepLists := service.GenerateBadStepLists()
	var badSchemas []map[models.ServiceToken][]models.TokenStruct

	for _, badStepList := range badStepLists {
		badSchema := make(map[models.ServiceToken][]models.TokenStruct)
		for _, badStep := range badStepList {
			badSchema[badStep] = nil
		}

		badSchemas = append(badSchemas, badSchema)
	}

	log.Println("badSchemas: ", badSchemas)

	return badSchemas
}

// TestSchemaCheckBadSchemas tests bad schemas.
func TestSchemaCheckBadSchemas(t *testing.T) {
	badSchemas := generateBadSchemas()

	for i := 0; i < len(badSchemas); i++ {
		err := service.SchemaCheck(badSchemas[i])
		assert.NotNil(t, err, "Error should be not nil")
		log.Println("test " + strconv.Itoa(i))
	}

}

//
//
//
