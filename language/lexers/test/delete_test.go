package test

import (
	"servicediscoverer/language/lexers"
	"testing"
)

var deleteQuerySimple = []string{
	"DELETE",
}

var deleteManyArguments = []string{
	"DELETE",
	"ASD",
}

func TestDeleteProcessSimple(t *testing.T) {
	deleteLex := &lexers.DeleteLex{}
	_, got := deleteLex.Process(&deleteQuerySimple)
	if len(got) == 0 {
		t.Errorf("Too many arguments on the test string")
	}
}

func TestDeleteProcessManyArguments(t *testing.T) {
	deleteLex := &lexers.DeleteLex{}
	_, _ = deleteLex.Process(&deleteManyArguments)
	if len(deleteManyArguments) == 0 {
		t.Errorf("Only one arguments")
	}
}
