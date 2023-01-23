package test

import (
	"github.com/stretchr/testify/assert"
	"servicediscoverer/service"
	"testing"
)

//
//
// Good cases
//
//

func TestNewTokenizerWithFromAndDelete(t *testing.T) {
	//test initialization
	var TokenizerTestFromAndDelete = "FROM Human.Add DELETE"
	var tokenizer = service.NewTokenizer()

	//Running of the test
	err, got := tokenizer.CommandProcess(TokenizerTestFromAndDelete)
	if err != nil && got != nil {
		t.Errorf("Something happened on the tokenizing process, %s", err)
	}
}

//
//
//

func TestNewTokenizerIrrationalButCorrect(t *testing.T) {
	//test initialization
	var TokenizerTestIrrationalButCorrectKeywords = "FROM Human.Add INFO Human.Add DELETE INSERT UPDATE SELECT Something"

	var tokenizer = service.NewTokenizer()

	//Running of the test
	err, _ := tokenizer.CommandProcess(TokenizerTestIrrationalButCorrectKeywords)
	if err != nil {
		t.Errorf("Something happened on the tokenizing process, %s", err)
	}
}

//
//
// Error cases
//
//

func TestNewTokenizerWithBadSchema(t *testing.T) {
	//test initialization
	var TokenizerTestFromAndDelete = "FROM Human.Add DELETE"
	var tokenizer = service.NewTokenizer()
	//Running of the test
	err, got := tokenizer.CommandProcess(TokenizerTestFromAndDelete)
	assert.NotNil(t, err, "Error should be nil")
	assert.Nil(t, got, "Got should be nil")
}
