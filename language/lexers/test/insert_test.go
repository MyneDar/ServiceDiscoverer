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

func TestInsertProcessSimpleQuery(t *testing.T) {
	//test initialization
	var insertSimpleQuery = []string{
		"INSERT",
	}

	var insertSimpleOutput = []models.TokenStruct{
		{models.INSERT, "INSERT"},
	}

	insertLex := &lexers.InsertLex{}

	expectedQuerySize := 0

	//Running of the test
	err, got := insertLex.Process(&insertSimpleQuery)

	goodOutputCheck(t, err, insertSimpleQuery, expectedQuerySize, insertSimpleOutput, got)
}

//
//
//

func TestInsertProcessManyArgumentsQuery(t *testing.T) {
	//test initialization
	var insertManyArgumentsQuery = []string{
		"INSERT",
		"ASD",
	}

	var insertManyArgumentsOutput = []models.TokenStruct{
		{models.INSERT, "INSERT"},
	}

	insertLex := &lexers.InsertLex{}

	expectedQuerySize := 1

	//Running of the test
	err, got := insertLex.Process(&insertManyArgumentsQuery)

	goodOutputCheck(t, err, insertManyArgumentsQuery, expectedQuerySize, insertManyArgumentsOutput, got)
}

//
//
// Error cases
//
//
