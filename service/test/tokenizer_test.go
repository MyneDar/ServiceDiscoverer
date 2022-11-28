package test

import (
	"servicediscoverer/dev"
	"servicediscoverer/service"
	"testing"
)

var TokenizerTestFromAndDelete = "FROM Human.Add DELETE"
var TokenizerTestIrrationalButCorrectKeywords = "FROM Human.Add INFO Human.Add DELETE INSERT UPDATE SELECT Something"

func TestNewTokenizerWithFromAndDelete(t *testing.T) {
	err := dev.EntClientInit()
	if err != nil {
		return
	}
	var tokenizer = service.NewTokenizer()
	var analyzer = service.NewLanguageAnalyzer()
	err, got := tokenizer.CommandProcess(TokenizerTestFromAndDelete)
	analyzer.TokenProcess(got, nil)
	if err != nil && got != nil {
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
