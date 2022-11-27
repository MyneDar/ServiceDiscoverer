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

	//Running of the test
	err, got := whereLex.Process(&whereSimpleQuery)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(whereSimpleOutput), len(got), "Wrong token number")

	for i, value := range whereSimpleOutput {
		assert.Equal(t, value.Name, got[i].Name, "Name doesn't match")
		assert.Equal(t, value.Data, got[i].Data, "Data doesn't match")
	}

	assert.Equal(t, 0, len(whereSimpleQuery), "The number of string is different in the Query")
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

	//Running of the test
	err, got := whereLex.Process(&whereParenthesisQuery)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(whereParenthesisOutput), len(got), "Wrong token number")

	for i, value := range whereParenthesisOutput {
		assert.Equal(t, value.Name, got[i].Name, "Name doesn't match")
		assert.Equal(t, value.Data, got[i].Data, "Data doesn't match")
	}

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

	//Running of the test
	err, got := whereLex.Process(&whereMultipleKeywordQuery)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(whereMultipleKeywordOutput), len(got), "Wrong token number")

	for i, value := range whereMultipleKeywordOutput {
		assert.Equal(t, value.Name, got[i].Name, "Name doesn't match")
		assert.Equal(t, value.Data, got[i].Data, "Data doesn't match")
	}
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

	//Running of the test
	err, got := whereLex.Process(&whereNotWhereQuery)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(whereNotWhereOutput), len(got), "Wrong token number")
}

//
//
// Error cases
//
//
