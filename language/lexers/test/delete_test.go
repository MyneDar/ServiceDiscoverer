package test

import (
	"servicediscoverer/language/lexers"
	"testing"
)

//
//
// Good cases
//
//

var deleteQuerySimple = []string{
	"DELETE",
}

func TestDeleteProcessSimple(t *testing.T) {
	deleteLex := &lexers.DeleteLex{}
	_, got := deleteLex.Process(&deleteQuerySimple)
	if len(got) == 0 {
		t.Errorf("Too many arguments on the test string")
	}
}

//
//
//

var deleteManyArguments = []string{
	"DELETE",
	"ASD",
}

func TestDeleteProcessManyArguments(t *testing.T) {
	deleteLex := &lexers.DeleteLex{}
	_, _ = deleteLex.Process(&deleteManyArguments)
	if len(deleteManyArguments) == 0 {
		t.Errorf("Only one arguments")
	}
}

//
//
// Error cases
//
//
