package test

import (
	"github.com/stretchr/testify/assert"
	"servicediscoverer/models"
	"testing"
)

//
//
// Good case checkers
//
//

// goodOutputCheck checks the good cases.
func goodOutputCheck(t *testing.T, err error, query []string, expectedQuerySize int, expectedOutput []models.TokenStruct, actualOutput []models.TokenStruct) {
	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(expectedOutput), len(actualOutput), "Wrong token number")

	for i, value := range expectedOutput {
		assert.Equal(t, value.Name, actualOutput[i].Name, "Name doesn't match")
		assert.Equal(t, value.Data, actualOutput[i].Data, "Data doesn't match")
	}

	assert.Equal(t, expectedQuerySize, len(query), "The number of string is different in the Query")
}

//
//
//

// goodNoOutputCheck checks the cases what are good but give back no output.
func goodNoOutputCheck(t *testing.T, err error, query []string, expectedQuerySize int, expectedOutput []models.TokenStruct, actualOutput []models.TokenStruct) {
	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(expectedOutput), len(actualOutput), "Wrong token number")

	assert.Equal(t, expectedQuerySize, len(query), "The number of string is different in the Query")
}

//
//
// Error case checkers
//
//

// errorCheck checks the cases where should be an error.
func errorCheck(t *testing.T, err error, query []string, expectedQuerySize int, expectedOutput []models.TokenStruct, actualOutput []models.TokenStruct) {
	assert.NotNil(t, err, "There should be an error")
	assert.Nil(t, actualOutput, "Value should be nil")
	assert.Equal(t, len(expectedOutput), len(actualOutput), "Wrong token number")

	assert.Equal(t, expectedQuerySize, len(query), "Doesn't equal with the initial value")
}
