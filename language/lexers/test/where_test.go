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

func TestWhereProcessSimpleQuery(t *testing.T) {
	//test initialization
	var whereSimpleQuery = []string{
		"WHERE",
		"dogName=\"Buksi\"",
	}

	var whereSimpleOutput = []models.TokenStruct{
		{models.WHERE, "WHERE"},
		{models.IDENT, "dogName"},
		{models.EQL, "="},
		{models.STRING, "Buksi"},
	}

	whereLex := &lexers.WhereLex{}

	expectedQuerySize := 0

	//Running of the test
	err, got := whereLex.Process(&whereSimpleQuery)

	goodOutputCheck(t, err, whereSimpleQuery, expectedQuerySize, whereSimpleOutput, got)
}

//
//
//

func TestWhereProcessParenthesisQuery(t *testing.T) {
	//test initialization
	var whereParenthesisQuery = []string{
		"WHERE",
		"((dogName=\"Buksi\")",
		"AND",
		"age",
		"=",
		"1.5)",
	}

	var whereParenthesisOutput = []models.TokenStruct{
		{models.WHERE, "WHERE"},
		{models.LPAREN, "("},
		{models.LPAREN, "("},
		{models.IDENT, "dogName"},
		{models.EQL, "="},
		{models.STRING, "Buksi"},
		{models.RPAREN, ")"},
		{models.AND, "AND"},
		{models.IDENT, "age"},
		{models.EQL, "="},
		{models.FLOAT, "1.5"},
		{models.RPAREN, ")"},
	}

	whereLex := &lexers.WhereLex{}

	expectedQuerySize := 0

	//Running of the test
	err, got := whereLex.Process(&whereParenthesisQuery)

	goodOutputCheck(t, err, whereParenthesisQuery, expectedQuerySize, whereParenthesisOutput, got)
}

//
//
//

func TestWhereProcessMultipleKeywordQuery(t *testing.T) {
	//test initialization
	var whereMultipleKeywordQuery = []string{
		"WHERE",
		"dogName=\"Buksi\"",
		"FROM",
	}

	var whereMultipleKeywordOutput = []models.TokenStruct{
		{models.WHERE, "WHERE"},
		{models.IDENT, "dogName"},
		{models.EQL, "="},
		{models.STRING, "Buksi"},
	}

	whereLex := &lexers.WhereLex{}

	expectedQuerySize := 1

	//Running of the test
	err, got := whereLex.Process(&whereMultipleKeywordQuery)

	goodOutputCheck(t, err, whereMultipleKeywordQuery, expectedQuerySize, whereMultipleKeywordOutput, got)
}

//
//
//

func TestWhereProcessNotWhereQuery(t *testing.T) {
	//test initialization
	var whereNotWhereQuery = []string{
		"FROM",
	}

	var whereNotWhereOutput []models.TokenStruct

	whereLex := &lexers.WhereLex{}

	expectedQuerySize := len(whereNotWhereQuery)

	//Running of the test
	err, got := whereLex.Process(&whereNotWhereQuery)

	goodNoOutputCheck(t, err, whereNotWhereQuery, expectedQuerySize, whereNotWhereOutput, got)
}

//
//
// Error cases
//
//
