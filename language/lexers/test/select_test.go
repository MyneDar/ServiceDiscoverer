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

var selectQueryAsterisk = []string{
	"SELECT",
	"*",
}

var selectQueryAsteriskOutput = []models.TokenStruct{
	{models.SELECT, "SELECT"},
	{models.ASTERISK, "*"},
}

func TestSelectProcessGoodAsterisk(t *testing.T) {
	selectLex := &lexers.SelectLex{}
	err, got := selectLex.Process(&selectQueryAsterisk)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(selectQueryAsteriskOutput), len(got), "Wrong token number")

	for i, value := range selectQueryAsteriskOutput {
		assert.Equal(t, value.Name, got[i].Name, "Name doesn't match")
		assert.Equal(t, value.Data, got[i].Data, "Data doesn't match")
	}
}

//
//
//

var selectQueryOneIndentOnly = []string{
	"SELECT",
	"Age",
}

var selectQueryOneIdentOnlyOutput = []models.TokenStruct{
	{models.SELECT, "SELECT"},
	{models.IDENT, "Age"},
}

func TestSelectProcessGoodIdent(t *testing.T) {
	selectLex := &lexers.SelectLex{}
	err, got := selectLex.Process(&selectQueryOneIndentOnly)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(selectQueryOneIdentOnlyOutput), len(got), "Wrong token number")

	for i, value := range selectQueryOneIdentOnlyOutput {
		assert.Equal(t, value.Name, got[i].Name, "Name doesn't match")
		assert.Equal(t, value.Data, got[i].Data, "Data doesn't match")
	}
}

//
//
//

var selectQueryMultipleIndents = []string{
	"SELECT",
	"Age",
	"Name",
}

var selectQueryMultipleIdentsOutput = []models.TokenStruct{
	{models.SELECT, "SELECT"},
	{models.IDENT, "Age"},
	{models.IDENT, "Name"},
}

func TestSelectProcessMultipleIndents(t *testing.T) {
	selectLex := &lexers.SelectLex{}
	err, got := selectLex.Process(&selectQueryMultipleIndents)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(selectQueryMultipleIdentsOutput), len(got), "Wrong token number")

	for i, value := range selectQueryMultipleIdentsOutput {
		assert.Equal(t, value.Name, got[i].Name, "Name doesn't match")
		assert.Equal(t, value.Data, got[i].Data, "Data doesn't match")
	}
}

//
//
//

var selectQueryMultipleKeyWords = []string{
	"INFO",
	"*",
	"SELECT",
	"INSERT",
}

var selectQueryMultipleKeyWordsOutput []models.TokenStruct

func TestSelectProcessMultipleKeyWords(t *testing.T) {
	selectLex := &lexers.SelectLex{}
	commandLength := len(selectQueryMultipleKeyWords)

	err, got := selectLex.Process(&selectQueryMultipleKeyWords)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(selectQueryMultipleKeyWordsOutput), len(got), "Wrong token number")
	assert.Equal(t, commandLength, len(selectQueryMultipleKeyWords), "Doesn't equal with the initial value")
}

//
//
//

var selectQueryUntilKeyWord = []string{
	"SELECT",
	"*",
	"WHERE",
}

var selectQueryUntilKeyWordOutput = []models.TokenStruct{
	{models.SELECT, "SELECT"},
	{models.ASTERISK, "*"},
}

func TestSelectProcessUntilKeyWord(t *testing.T) {
	selectLex := &lexers.SelectLex{}
	err, got := selectLex.Process(&selectQueryUntilKeyWord)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(selectQueryUntilKeyWordOutput), len(got), "Wrong token number")

	for i, value := range selectQueryUntilKeyWordOutput {
		assert.Equal(t, value.Name, got[i].Name, "Name doesn't match")
		assert.Equal(t, value.Data, got[i].Data, "Data doesn't match")
	}
}

//
//
// Error cases
//
//

var selectQueryAsteriskWithIdent = []string{
	"SELECT",
	"*",
	"Age",
}

var selectQueryAsteriskWithIdentOutput []models.TokenStruct

func TestSelectProcessAsteriskWithIdent(t *testing.T) {
	selectLex := &lexers.SelectLex{}
	commandLength := len(selectQueryAsteriskWithIdent)

	err, got := selectLex.Process(&selectQueryAsteriskWithIdent)

	assert.NotNil(t, err, "There should be an error")
	assert.Nil(t, got, "Value should be nil")
	assert.Equal(t, len(selectQueryAsteriskWithIdentOutput), len(got), "Wrong token number")
	assert.Equal(t, commandLength, len(selectQueryAsteriskWithIdent), "Doesn't equal with the initial value")
}

//
//
//

var selectQueryOnlySelect = []string{
	"SELECT",
}

var selectQueryOnlySelectOutput []models.TokenStruct

func TestSelectProcessOnlySelect(t *testing.T) {
	selectLex := &lexers.SelectLex{}
	commandLength := len(selectQueryOnlySelect)

	err, got := selectLex.Process(&selectQueryOnlySelect)

	assert.NotNil(t, err, "There should be an error")
	assert.Nil(t, got, "Value should be nil")
	assert.Equal(t, len(selectQueryOnlySelectOutput), len(got), "Wrong token number")
	assert.Equal(t, commandLength, len(selectQueryOnlySelect), "Doesn't equal with the initial value")
}

//
//
//
