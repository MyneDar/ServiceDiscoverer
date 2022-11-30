package test

import (
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
