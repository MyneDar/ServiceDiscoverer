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

func TestInsertProcessSimpleQuery(t *testing.T) {
	//test initialization
	var insertSimpleQuery = []string{
		"INSERT",
	}

	insertLex := &lexers.InsertLex{}

	//Running of the test
	_, got := insertLex.Process(&insertSimpleQuery)
	if len(got) == 0 {
		t.Errorf("Too many arguments on the test string")
	}
}

//
//
//

func TestInsertProcessManyArgumentsQuery(t *testing.T) {
	//test initialization
	var insertManyArgumentsQuery = []string{
		"INSERT",
		"ASD",
	}

	insertLex := &lexers.InsertLex{}

	//Running of the test
	_, _ = insertLex.Process(&insertManyArgumentsQuery)
	if len(insertManyArgumentsQuery) == 0 {
		t.Errorf("Only one arguments")
	}
}

//
//
// Error cases
//
//
