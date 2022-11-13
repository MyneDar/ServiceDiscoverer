package test

import (
	"servicediscoverer/service"
	"testing"
)

var TokenizerTestFromAndDelete = "FROM Human.Add DELETE"
var TokenizerTestIrrationalButCorrectKeywords = "FROM Human.Add INFO Human.Add DELETE INSERT UPDATE SELECT Something"

func TestNewTokenizerWithFromAndDelete(t *testing.T) {
	var tokenizer = service.NewTokenizer()
	err, _ := tokenizer.CommandProcess(TokenizerTestFromAndDelete)
	if err != nil {
		t.Errorf("Something happened on the tokenizing process, %s", err)
	}
}

func TestNewTokenizerIrrationalButCorrect(t *testing.T) {
	var tokenizer = service.NewTokenizer()
	err, _ := tokenizer.CommandProcess(TokenizerTestIrrationalButCorrectKeywords)
	if err != nil {
		t.Errorf("Something happened on the tokenizing process, %s", err)
	}
}
