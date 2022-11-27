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

func TestInfoProcessAsteriskQuery(t *testing.T) {
	//test initialization
	var infoAsteriskQuery = []string{
		"INFO",
		"*",
	}

	var infoAsteriskOutput = []models.TokenStruct{
		{models.INFO, "INFO"},
		{models.ASTERISK, "*"},
	}

	infoLex := &lexers.InfoLex{}

	expectedQuerySize := 0

	//Running of the test
	err, got := infoLex.Process(&infoAsteriskQuery)

	goodOutputCheck(t, err, infoAsteriskQuery, expectedQuerySize, infoAsteriskOutput, got)
}

//
//
//

func TestInfoProcessIdentDotAsteriskQuery(t *testing.T) {
	//test initialization
	var infoIdentDotAsteriskQuery = []string{
		"INFO",
		"Human.*",
	}

	var infoIdentDotAsteriskOutput = []models.TokenStruct{
		{models.INFO, "INFO"},
		{models.IDENT, "Human"},
		{models.ASTERISK, "*"},
	}

	infoLex := &lexers.InfoLex{}

	expectedQuerySize := 0

	//Running of the test
	err, got := infoLex.Process(&infoIdentDotAsteriskQuery)

	goodOutputCheck(t, err, infoIdentDotAsteriskQuery, expectedQuerySize, infoIdentDotAsteriskOutput, got)
}

//
//
//

func TestInfoProcessGoodIdentDotIdentQuery(t *testing.T) {
	//test initialization
	var infoIdentDotIdentQuery = []string{
		"INFO",
		"Human.Add",
	}

	var infoIdentDotIdentOutput = []models.TokenStruct{
		{models.INFO, "INFO"},
		{models.IDENT, "Human"},
		{models.IDENT, "Add"},
	}

	infoLex := &lexers.InfoLex{}

	expectedQuerySize := 0

	//Running of the test
	err, got := infoLex.Process(&infoIdentDotIdentQuery)

	goodOutputCheck(t, err, infoIdentDotIdentQuery, expectedQuerySize, infoIdentDotIdentOutput, got)
}

//
//
//

func TestInfoProcessMoreKeyWordsQuery(t *testing.T) {
	//test initialization
	var infoMoreKeyWordsQuery = []string{
		"INFO",
		"Human.Add",
		"SELECT",
		"INSERT",
	}

	var infoMoreKeyWordsOutput = []models.TokenStruct{
		{models.INFO, "INFO"},
		{models.IDENT, "Human"},
		{models.IDENT, "Add"},
	}

	infoLex := &lexers.InfoLex{}

	expectedQuerySize := 2

	//Running of the test
	err, got := infoLex.Process(&infoMoreKeyWordsQuery)

	goodOutputCheck(t, err, infoMoreKeyWordsQuery, expectedQuerySize, infoMoreKeyWordsOutput, got)
}

//
//
//

func TestInfoProcessOtherCommandWordQuery(t *testing.T) {
	//test initialization
	var infoOtherCommandWordQuery = []string{
		"ASD",
		"Human.Add",
	}

	var infoOtherCommandWordOutput []models.TokenStruct

	infoLex := &lexers.InfoLex{}

	expectedQuerySize := len(infoOtherCommandWordQuery)

	//Running of the test
	err, got := infoLex.Process(&infoOtherCommandWordQuery)

	goodNoOutputCheck(t, err, infoOtherCommandWordQuery, expectedQuerySize, infoOtherCommandWordOutput, got)
}

//
//
// Error cases
//
//

func TestInfoProcessWrongTargetQuery(t *testing.T) {
	//test initialization
	var infoWrongTargetQuery = []string{
		"INFO",
		"HumanAdd",
	}

	var infoWrongTargetOutput []models.TokenStruct

	infoLex := &lexers.InfoLex{}

	expectedQuerySize := len(infoWrongTargetQuery)

	//Running of the test
	err, got := infoLex.Process(&infoWrongTargetQuery)

	errorCheck(t, err, infoWrongTargetQuery, expectedQuerySize, infoWrongTargetOutput, got)
}

//
//
//
