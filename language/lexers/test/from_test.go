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

func TestFromProcessSimpleQuery(t *testing.T) {
	//test initialization
	var fromSimpleQuery = []string{
		"FROM",
		"Human.Add",
	}

	var fromSimpleOutput = []models.TokenStruct{
		{models.FROM, "FROM"},
		{models.IDENT, "Human"},
		{models.IDENT, "Add"},
	}

	fromLex := &lexers.FromLex{}

	expectedQuerySize := 0

	//Running of the test
	err, got := fromLex.Process(&fromSimpleQuery)

	goodOutputCheck(t, err, fromSimpleQuery, expectedQuerySize, fromSimpleOutput, got)
}

//
//
//

func TestFromProcessMoreKeyWordsQuery(t *testing.T) {
	//test initialization
	var fromMoreKeyWordsQuery = []string{
		"FROM",
		"Human.Add",
		"SELECT",
		"INSERT",
	}

	var fromMoreKeyWordsOutput = []models.TokenStruct{
		{models.FROM, "FROM"},
		{models.IDENT, "Human"},
		{models.IDENT, "Add"},
	}

	fromLex := &lexers.FromLex{}

	expectedQuerySize := 2

	//Running of the test
	err, got := fromLex.Process(&fromMoreKeyWordsQuery)

	goodOutputCheck(t, err, fromMoreKeyWordsQuery, expectedQuerySize, fromMoreKeyWordsOutput, got)
}

//
//
//

func TestFromProcessOtherCommandWordQuery(t *testing.T) {
	//test initialization
	var fromOtherCommandWordQuery = []string{
		"ASD",
		"Human.Add",
	}

	var fromOtherCommandWordOutput []models.TokenStruct

	fromLex := &lexers.FromLex{}

	expectedQuerySize := len(fromOtherCommandWordQuery)

	//Running of the test
	err, got := fromLex.Process(&fromOtherCommandWordQuery)

	goodNoOutputCheck(t, err, fromOtherCommandWordQuery, expectedQuerySize, fromOtherCommandWordOutput, got)
}

//
//
// Error cases
//
//

func TestFromProcessWrongTargetQuery(t *testing.T) {
	//test initialization
	var fromWrongTargetQuery = []string{
		"FROM",
		"HumanAdd",
	}

	var fromWrongTargetOutput []models.TokenStruct

	fromLex := &lexers.FromLex{}

	expectedQuerySize := len(fromWrongTargetQuery)

	//Running of the test
	err, got := fromLex.Process(&fromWrongTargetQuery)

	errorCheck(t, err, fromWrongTargetQuery, expectedQuerySize, fromWrongTargetOutput, got)
}

//
//
//
