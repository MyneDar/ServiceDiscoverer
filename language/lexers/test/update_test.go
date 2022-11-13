package test

import (
	"servicediscoverer/language/lexers"
	"testing"
)

var updateQuerySimple = []string{
	"UPDATE",
}

var updateManyArguments = []string{
	"UPDATE",
	"ASD",
}

func TestUpdateProcessSimple(t *testing.T) {
	insertLex := &lexers.UpdateLex{}
	_, got := insertLex.Process(&updateQuerySimple)
	if len(got) == 0 {
		t.Errorf("Too many arguments on the test string")
	}
}

func TestUpdateProcessManyArguments(t *testing.T) {
	insertLex := &lexers.UpdateLex{}
	_, _ = insertLex.Process(&updateManyArguments)
	if len(updateManyArguments) == 0 {
		t.Errorf("Only one arguments")
	}
}
