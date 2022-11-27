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

func TestDeleteProcessSimpleQuery(t *testing.T) {
	//test initialization
	var deleteSimpleQuery = []string{
		"DELETE",
	}

	deleteLex := &lexers.DeleteLex{}

	//Running of the test
	_, got := deleteLex.Process(&deleteSimpleQuery)
	if len(got) == 0 {
		t.Errorf("Too many arguments on the test string")
	}
}

//
//
//

func TestDeleteProcessManyArgumentsQuery(t *testing.T) {
	//test initialization
	var deleteManyArgumentsQuery = []string{
		"DELETE",
		"ASD",
	}

	deleteLex := &lexers.DeleteLex{}

	//Running of the test
	_, _ = deleteLex.Process(&deleteManyArgumentsQuery)
	if len(deleteManyArgumentsQuery) == 0 {
		t.Errorf("Only one arguments")
	}
}

//
//
// Error cases
//
//
