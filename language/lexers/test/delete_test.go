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

func TestDeleteProcessSimpleQuery(t *testing.T) {
	//test initialization
	var deleteSimpleQuery = []string{
		"DELETE",
	}

	var deleteSimpleOutput = []models.TokenStruct{
		{models.DELETE, "DELETE"},
	}

	deleteLex := &lexers.DeleteLex{}

	expectedQuerySize := 0

	//Running of the test
	err, got := deleteLex.Process(&deleteSimpleQuery)

	goodOutputCheck(t, err, deleteSimpleQuery, expectedQuerySize, deleteSimpleOutput, got)
}

//
//
//

func TestDeleteProcessManyArgumentsQuery(t *testing.T) {
	//test initialization
	var deleteManyArgumentsQuery = []string{
		"DELETE",
		"ASD",
	}

	var deleteManyArgumentsOutput = []models.TokenStruct{
		{models.DELETE, "DELETE"},
	}

	deleteLex := &lexers.DeleteLex{}

	expectedQuerySize := 1

	//Running of the test
	err, got := deleteLex.Process(&deleteManyArgumentsQuery)

	goodOutputCheck(t, err, deleteManyArgumentsQuery, expectedQuerySize, deleteManyArgumentsOutput, got)
}

//
//
// Error cases
//
//
