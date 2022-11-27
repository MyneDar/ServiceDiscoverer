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

func TestSelectProcessAsteriskQuery(t *testing.T) {
	//test initialization
	selectAsteriskQuery := []string{
		"SELECT",
		"*",
	}

	selectAsteriskOutput := []models.TokenStruct{
		{models.SELECT, "SELECT"},
		{models.ASTERISK, "*"},
	}

	selectLex := &lexers.SelectLex{}

	expectedQuerySize := 0

	//Running of the test
	err, got := selectLex.Process(&selectAsteriskQuery)

	goodOutputCheck(t, err, selectAsteriskQuery, expectedQuerySize, selectAsteriskOutput, got)
}

//
//
//

func TestSelectProcessOneIdentOnlyQuery(t *testing.T) {
	//test initialization
	selectOneIndentOnlyQuery := []string{
		"SELECT",
		"Age",
	}

	selectOneIdentOnlyOutput := []models.TokenStruct{
		{models.SELECT, "SELECT"},
		{models.IDENT, "Age"},
	}

	selectLex := &lexers.SelectLex{}

	expectedQuerySize := 0

	//Running of the test
	err, got := selectLex.Process(&selectOneIndentOnlyQuery)

	goodOutputCheck(t, err, selectOneIndentOnlyQuery, expectedQuerySize, selectOneIdentOnlyOutput, got)
}

//
//
//

func TestSelectProcessMultipleIndentsQuery(t *testing.T) {
	//test initialization
	selectMultipleIndentsQuery := []string{
		"SELECT",
		"Age",
		"Name",
	}

	selectMultipleIdentsOutput := []models.TokenStruct{
		{models.SELECT, "SELECT"},
		{models.IDENT, "Age"},
		{models.IDENT, "Name"},
	}

	selectLex := &lexers.SelectLex{}

	expectedQuerySize := 0

	//Running of the test
	err, got := selectLex.Process(&selectMultipleIndentsQuery)

	goodOutputCheck(t, err, selectMultipleIndentsQuery, expectedQuerySize, selectMultipleIdentsOutput, got)
}

//
//
//

func TestSelectProcessMultipleKeyWordsQuery(t *testing.T) {
	//test initialization
	selectMultipleKeyWordsQuery := []string{
		"INFO",
		"*",
		"SELECT",
		"INSERT",
	}

	var selectMultipleKeyWordsOutput []models.TokenStruct

	selectLex := &lexers.SelectLex{}

	expectedQuerySize := len(selectMultipleKeyWordsQuery)

	//Running of the test
	err, got := selectLex.Process(&selectMultipleKeyWordsQuery)

	goodNoOutputCheck(t, err, selectMultipleKeyWordsQuery, expectedQuerySize, selectMultipleKeyWordsOutput, got)
}

//
//
//

func TestSelectProcessUntilKeyWordQuery(t *testing.T) {
	//test initialization
	selectUntilKeyWordQuery := []string{
		"SELECT",
		"*",
		"WHERE",
	}

	selectUntilKeyWordOutput := []models.TokenStruct{
		{models.SELECT, "SELECT"},
		{models.ASTERISK, "*"},
	}

	selectLex := &lexers.SelectLex{}

	expectedQuerySize := 1

	//Running of the test
	err, got := selectLex.Process(&selectUntilKeyWordQuery)

	goodOutputCheck(t, err, selectUntilKeyWordQuery, expectedQuerySize, selectUntilKeyWordOutput, got)
}

//
//
// Error cases
//
//

func TestSelectProcessAsteriskWithIdentQuery(t *testing.T) {
	//test initialization
	selectAsteriskWithIdentQuery := []string{
		"SELECT",
		"*",
		"Age",
	}

	var selectAsteriskWithIdentOutput []models.TokenStruct

	selectLex := &lexers.SelectLex{}

	expectedQuerySize := len(selectAsteriskWithIdentQuery)

	//Running of the test
	err, got := selectLex.Process(&selectAsteriskWithIdentQuery)

	errorCheck(t, err, selectAsteriskWithIdentQuery, expectedQuerySize, selectAsteriskWithIdentOutput, got)
}

//
//
//

func TestSelectProcessOnlySelectQuery(t *testing.T) {
	//test initialization
	selectOnlySelectQuery := []string{
		"SELECT",
	}

	var selectOnlySelectOutput []models.TokenStruct

	selectLex := &lexers.SelectLex{}

	expectedQuerySize := len(selectOnlySelectQuery)

	//Running of the test
	err, got := selectLex.Process(&selectOnlySelectQuery)

	errorCheck(t, err, selectOnlySelectQuery, expectedQuerySize, selectOnlySelectOutput, got)
}

//
//
//
