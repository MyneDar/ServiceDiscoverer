package test

import (
	"servicediscoverer/language/lexers"
	"testing"
)

var simpleCommand = []string{
	"WHERE",
	"dogName=\"Buksi\"",
}

func TestSimpleCommand(t *testing.T) {
	whereLex := &lexers.WhereLex{}
	_, got := whereLex.Process(&simpleCommand)
	if len(got) != 4 {
		t.Errorf("Wrong token number: %d", len(got))
	}

}

var parenthesisCommand = []string{
	"WHERE",
	"((dogName=\"Buksi\")",
	"AND",
	"age",
	"=",
	"1.5)",
}

func TestParenthesisCommand(t *testing.T) {
	whereLex := &lexers.WhereLex{}
	_, got := whereLex.Process(&parenthesisCommand)
	if len(got) != 12 {
		t.Errorf("Wrong token number: %d", len(got))
	}

}
