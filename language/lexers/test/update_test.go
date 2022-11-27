package test

import (
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
	updateSimpleQuery := []string{
		"UPDATE",
	}

	updateSimpleOutput := []models.TokenStruct{
		{models.UPDATE, "UPDATE"},
	}

	insertLex := &lexers.UpdateLex{}

	expectedQuerySize := 0

	//Running of the test
	err, got := insertLex.Process(&updateSimpleQuery)

	goodOutputCheck(t, err, updateSimpleQuery, expectedQuerySize, updateSimpleOutput, got)
}

//
//
//

func TestUpdateProcessManyArgumentsQuery(t *testing.T) {
	//test initialization
	updateManyArgumentsQuery := []string{
		"UPDATE",
		"ASD",
	}

	updateManyArgumentsOutput := []models.TokenStruct{
		{models.UPDATE, "UPDATE"},
	}

	insertLex := &lexers.UpdateLex{}

	expectedQuerySize := 1

	//Running of the test
	err, got := insertLex.Process(&updateManyArgumentsQuery)

	goodOutputCheck(t, err, updateManyArgumentsQuery, expectedQuerySize, updateManyArgumentsOutput, got)
}

//
//
// Error cases
//
//
