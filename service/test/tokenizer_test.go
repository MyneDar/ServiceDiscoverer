package test

import (
	"servicediscoverer/service"
	"testing"
)

var TokenizerTest = "FROM Human.Add INFO Human.Add"

func TestNewTokenizer(t *testing.T) {
	var tokenizer = service.NewTokenizer()
	err, tokens := tokenizer.CommandProcess(TokenizerTest)
	if err != nil {
		return
	}
	println(len(tokens))
}
