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

func TestSelectProcessAsteriskQuery(t *testing.T) {
	//test initialization
	var selectAsteriskQuery = []string{
		"SELECT",
		"*",
	}

	var selectAsteriskOutput = []models.TokenStruct{
		{models.SELECT, "SELECT"},
		{models.ASTERISK, "*"},
	}

	selectLex := &lexers.SelectLex{}

	//Running of the test
	err, got := selectLex.Process(&selectAsteriskQuery)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(selectAsteriskOutput), len(got), "Wrong token number")

	for i, value := range selectAsteriskOutput {
		assert.Equal(t, value.Name, got[i].Name, "Name doesn't match")
		assert.Equal(t, value.Data, got[i].Data, "Data doesn't match")
	}
}

//
//
//

func TestSelectProcessOneIdentOnlyQuery(t *testing.T) {
	//test initialization
	var selectOneIndentOnlyQuery = []string{
		"SELECT",
		"Age",
	}

	var selectOneIdentOnlyOutput = []models.TokenStruct{
		{models.SELECT, "SELECT"},
		{models.IDENT, "Age"},
	}

	selectLex := &lexers.SelectLex{}

	//Running of the test
	err, got := selectLex.Process(&selectOneIndentOnlyQuery)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(selectOneIdentOnlyOutput), len(got), "Wrong token number")

	for i, value := range selectOneIdentOnlyOutput {
		assert.Equal(t, value.Name, got[i].Name, "Name doesn't match")
		assert.Equal(t, value.Data, got[i].Data, "Data doesn't match")
	}
}

//
//
//

func TestSelectProcessMultipleIndentsQuery(t *testing.T) {
	//test initialization
	var selectMultipleIndentsQuery = []string{
		"SELECT",
		"Age",
		"Name",
	}

	var selectMultipleIdentsOutput = []models.TokenStruct{
		{models.SELECT, "SELECT"},
		{models.IDENT, "Age"},
		{models.IDENT, "Name"},
	}

	selectLex := &lexers.SelectLex{}

	//Running of the test
	err, got := selectLex.Process(&selectMultipleIndentsQuery)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(selectMultipleIdentsOutput), len(got), "Wrong token number")

	for i, value := range selectMultipleIdentsOutput {
		assert.Equal(t, value.Name, got[i].Name, "Name doesn't match")
		assert.Equal(t, value.Data, got[i].Data, "Data doesn't match")
	}
}

//
//
//

func TestSelectProcessMultipleKeyWordsQuery(t *testing.T) {
	//test initialization
	var selectMultipleKeyWordsQuery = []string{
		"INFO",
		"*",
		"SELECT",
		"INSERT",
	}

	var selectMultipleKeyWordsOutput []models.TokenStruct

	selectLex := &lexers.SelectLex{}

	commandLength := len(selectMultipleKeyWordsQuery)

	//Running of the test
	err, got := selectLex.Process(&selectMultipleKeyWordsQuery)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(selectMultipleKeyWordsOutput), len(got), "Wrong token number")
	assert.Equal(t, commandLength, len(selectMultipleKeyWordsQuery), "Doesn't equal with the initial value")
}

//
//
//

func TestSelectProcessUntilKeyWordQuery(t *testing.T) {
	//test initialization
	var selectUntilKeyWordQuery = []string{
		"SELECT",
		"*",
		"WHERE",
	}

	var selectUntilKeyWordOutput = []models.TokenStruct{
		{models.SELECT, "SELECT"},
		{models.ASTERISK, "*"},
	}

	selectLex := &lexers.SelectLex{}

	//Running of the test
	err, got := selectLex.Process(&selectUntilKeyWordQuery)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(selectUntilKeyWordOutput), len(got), "Wrong token number")

	for i, value := range selectUntilKeyWordOutput {
		assert.Equal(t, value.Name, got[i].Name, "Name doesn't match")
		assert.Equal(t, value.Data, got[i].Data, "Data doesn't match")
	}
}

//
//
// Error cases
//
//

func TestSelectProcessAsteriskWithIdentQuery(t *testing.T) {
	//test initialization
	var selectAsteriskWithIdentQuery = []string{
		"SELECT",
		"*",
		"Age",
	}

	var selectAsteriskWithIdentOutput []models.TokenStruct

	selectLex := &lexers.SelectLex{}

	commandLength := len(selectAsteriskWithIdentQuery)

	//Running of the test
	err, got := selectLex.Process(&selectAsteriskWithIdentQuery)

	assert.NotNil(t, err, "There should be an error")
	assert.Nil(t, got, "Value should be nil")
	assert.Equal(t, len(selectAsteriskWithIdentOutput), len(got), "Wrong token number")
	assert.Equal(t, commandLength, len(selectAsteriskWithIdentQuery), "Doesn't equal with the initial value")
}

//
//
//

func TestSelectProcessOnlySelectQuery(t *testing.T) {
	//test initialization
	var selectOnlySelectQuery = []string{
		"SELECT",
	}

	var selectOnlySelectOutput []models.TokenStruct

	selectLex := &lexers.SelectLex{}

	commandLength := len(selectOnlySelectQuery)

	//Running of the test
	err, got := selectLex.Process(&selectOnlySelectQuery)

	assert.NotNil(t, err, "There should be an error")
	assert.Nil(t, got, "Value should be nil")
	assert.Equal(t, len(selectOnlySelectOutput), len(got), "Wrong token number")
	assert.Equal(t, commandLength, len(selectOnlySelectQuery), "Doesn't equal with the initial value")
}

//
//
//
