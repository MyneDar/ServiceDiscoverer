package test

import (
	"servicediscoverer/language/lexers"
	"testing"
)

var insertQuerySimple = []string{
	"INSERT",
}

var insertManyArguments = []string{
	"INSERT",
	"ASD",
}

func TestInsertProcessSimple(t *testing.T) {
	insertLex := &lexers.InsertLex{}
	_, got := insertLex.Process(&insertQuerySimple)
	if len(got) == 0 {
		t.Errorf("Too many arguments on the test string")
	}
}

func TestInsertProcessManyArguments(t *testing.T) {
	insertLex := &lexers.InsertLex{}
	_, _ = insertLex.Process(&insertManyArguments)
	if len(insertManyArguments) == 0 {
		t.Errorf("Only one arguments")
	}
}
