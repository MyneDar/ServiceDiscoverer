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

var simpleCommand = []string{
	"WHERE",
	"dogName=\"Buksi\"",
}

var simpleCommandOutput = []models.TokenStruct{
	{models.WHERE, "WHERE"},
	{models.IDENT, "dogName"},
	{models.EQL, "="},
	{models.STRING, "Buksi"},
}

func TestWhereProcessSimpleCommand(t *testing.T) {
	whereLex := &lexers.WhereLex{}
	err, got := whereLex.Process(&simpleCommand)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(simpleCommandOutput), len(got), "Wrong token number")

	for i, value := range simpleCommandOutput {
		assert.Equal(t, value.Name, got[i].Name, "Name doesn't match")
		assert.Equal(t, value.Data, got[i].Data, "Data doesn't match")
	}
}

//
//
//

var parenthesisCommand = []string{
	"WHERE",
	"((dogName=\"Buksi\")",
	"AND",
	"age",
	"=",
	"1.5)",
}

var parenthesisCommandOutput = []models.TokenStruct{
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

func TestWhereProcessParenthesisCommand(t *testing.T) {
	whereLex := &lexers.WhereLex{}
	err, got := whereLex.Process(&parenthesisCommand)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(parenthesisCommandOutput), len(got), "Wrong token number")

	for i, value := range parenthesisCommandOutput {
		assert.Equal(t, value.Name, got[i].Name, "Name doesn't match")
		assert.Equal(t, value.Data, got[i].Data, "Data doesn't match")
	}

}

//
//
//

var multipleKeywordCommand = []string{
	"WHERE",
	"dogName=\"Buksi\"",
	"FROM",
}

var multipleKeywordCommandOutput = []models.TokenStruct{
	{models.WHERE, "WHERE"},
	{models.IDENT, "dogName"},
	{models.EQL, "="},
	{models.STRING, "Buksi"},
}

func TestWhereProcessMultipleKeywordCommand(t *testing.T) {
	whereLex := &lexers.WhereLex{}
	err, got := whereLex.Process(&multipleKeywordCommand)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(multipleKeywordCommandOutput), len(got), "Wrong token number")

	for i, value := range multipleKeywordCommandOutput {
		assert.Equal(t, value.Name, got[i].Name, "Name doesn't match")
		assert.Equal(t, value.Data, got[i].Data, "Data doesn't match")
	}
}

//
//
//

var notWhereCommand = []string{
	"FROM",
}

var notWhereCommandOutput []models.TokenStruct

func TestWhereProcessNotWhereCommand(t *testing.T) {
	whereLex := &lexers.WhereLex{}
	err, got := whereLex.Process(&notWhereCommand)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(notWhereCommandOutput), len(got), "Wrong token number")
}

//
//
// Error cases
//
//
