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

func TestFromProcessSimpleQuery(t *testing.T) {
	//test initialization
	var fromSimpleQuery = []string{
		"FROM",
		"Human.Add",
	}

	fromLex := &lexers.FromLex{}

	//Running of the test
	err, got := fromLex.Process(&fromSimpleQuery)
	if l := len(got); l != 3 {
		t.Errorf("Wrong token number, got: %d, want 3. Err: %s", l, err)
	}
}

//
//
//

func TestFromProcessMoreKeyWordsQuery(t *testing.T) {
	//test initialization
	var fromMoreKeyWordsQuery = []string{
		"FROM",
		"Human.Add",
		"SELECT",
		"INSERT",
	}

	fromLex := &lexers.FromLex{}

	//Running of the test
	_, _ = fromLex.Process(&fromMoreKeyWordsQuery)
	if len(fromMoreKeyWordsQuery) != 2 {
		t.Errorf("Wrong sliced command, got %s", fromMoreKeyWordsQuery)
	}
}

//
//
//

func TestFromProcessOtherCommandWordQuery(t *testing.T) {
	//test initialization
	var fromOtherCommandWordQuery = []string{
		"ASD",
		"Human.Add",
	}

	fromLex := &lexers.FromLex{}

	//Running of the test
	err, got := fromLex.Process(&fromOtherCommandWordQuery)
	if err != nil && len(got) != 0 {
		t.Errorf("No error on wrong command : %s", fromOtherCommandWordQuery)
	}
}

//
//
//

func TestFromProcessWrongTargetQuery(t *testing.T) {
	//test initialization
	var fromWrongTargetQuery = []string{
		"FROM",
		"HumanAdd",
	}

	fromLex := &lexers.FromLex{}

	//Running of the test
	err, got := fromLex.Process(&fromWrongTargetQuery)
	if err != nil && len(got) != 0 {
		t.Errorf("No error on target command : %s", fromWrongTargetQuery)
	}
}

//
//
// Error cases
//
//
