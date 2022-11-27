package test

import (
	"github.com/stretchr/testify/assert"
	"servicediscoverer/language/lexers"
	"servicediscoverer/models"
	"testing"
)

//
//
// Good cases
//
//

func TestUpdateProcessSimpleQuery(t *testing.T) {
	//test initialization
	var updateSimpleQuery = []string{
		"UPDATE",
	}

	var updateSimpleOutput = []models.TokenStruct{
		{models.UPDATE, "UPDATE"},
	}

	insertLex := &lexers.UpdateLex{}

	//Running of the test
	err, got := insertLex.Process(&updateSimpleQuery)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(updateSimpleOutput), len(got), "Wrong token number")

	for i, value := range updateSimpleOutput {
		assert.Equal(t, value.Name, got[i].Name, "Name doesn't match")
		assert.Equal(t, value.Data, got[i].Data, "Data doesn't match")
	}

	assert.Equal(t, 0, len(updateSimpleQuery), "The number of string is different in the Query")
}

//
//
//

func TestUpdateProcessManyArgumentsQuery(t *testing.T) {
	//test initialization
	var updateManyArgumentsQuery = []string{
		"UPDATE",
		"ASD",
	}

	var updateManyArgumentsOutput = []models.TokenStruct{
		{models.UPDATE, "UPDATE"},
	}

	insertLex := &lexers.UpdateLex{}

	//Running of the test
	err, got := insertLex.Process(&updateManyArgumentsQuery)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(updateManyArgumentsOutput), len(got), "Wrong token number")

	for i, value := range updateManyArgumentsOutput {
		assert.Equal(t, value.Name, got[i].Name, "Name doesn't match")
		assert.Equal(t, value.Data, got[i].Data, "Data doesn't match")
	}

	assert.Equal(t, 1, len(updateManyArgumentsQuery), "The number of string is different in the Query")
}

//
//
// Error cases
//
//
