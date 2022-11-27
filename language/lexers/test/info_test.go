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

func TestInfoProcessGoodAsterisk(t *testing.T) {
	//test initialization
	var infoQueryAsterisk = []string{
		"INFO",
		"*",
	}

	infoLex := &lexers.InfoLex{}

	//Running of the test
	err, got := infoLex.Process(&infoQueryAsterisk)
	if l := len(got); l != 2 {
		t.Errorf("Wrong token number, got: %d, want 3. Err: %s", l, err)
	}
}

//
//
//

func TestInfoProcessGoodIdentDotAsterisk(t *testing.T) {
	//test initialization
	var infoQueryIdentDotAsterisk = []string{
		"INFO",
		"Human.*",
	}

	infoLex := &lexers.InfoLex{}

	//Running of the test
	err, got := infoLex.Process(&infoQueryIdentDotAsterisk)
	if l := len(got); l != 3 {
		t.Errorf("Wrong token number, got: %d, want 3. Err: %s", l, err)
	}
}

//
//
//

func TestInfoProcessGoodIdentDotIdentQuery(t *testing.T) {
	//test initialization
	var infoQueryIdentDotIdentQuery = []string{
		"INFO",
		"Human.Add",
	}

	infoLex := &lexers.InfoLex{}

	//Running of the test
	err, got := infoLex.Process(&infoQueryIdentDotIdentQuery)
	if l := len(got); l != 3 {
		t.Errorf("Wrong token number, got: %d, want 3. Err: %s", l, err)
	}
}

//
//
//

func TestInfoProcessMoreKeyWordsQuery(t *testing.T) {
	//test initialization
	var infoQueryMoreKeyWordsQuery = []string{
		"INFO",
		"Human.Add",
		"SELECT",
		"INSERT",
	}

	infoLex := &lexers.InfoLex{}

	//Running of the test
	_, _ = infoLex.Process(&infoQueryMoreKeyWordsQuery)
	if len(infoQueryMoreKeyWordsQuery) != 2 {
		t.Errorf("Wrong sliced command, got %s", infoQueryMoreKeyWordsQuery)
	}
}

//
//
//

func TestInfoProcessOtherCommandWordQuery(t *testing.T) {
	//test initialization
	var infoOtherCommandWordQuery = []string{
		"ASD",
		"Human.Add",
	}

	infoLex := &lexers.InfoLex{}

	//Running of the test
	err, got := infoLex.Process(&infoOtherCommandWordQuery)
	if err != nil && len(got) != 0 {
		t.Errorf("No error on wrong command : %s", infoOtherCommandWordQuery)
	}
}

//
//
//

func TestInfoProcessWrongTargetQuery(t *testing.T) {
	//test initialization
	var infoWrongTargetQuery = []string{
		"INFO",
		"HumanAdd",
	}

	infoLex := &lexers.InfoLex{}

	//Running of the test
	err, got := infoLex.Process(&infoWrongTargetQuery)
	if err == nil && len(got) != 0 {
		t.Errorf("No error on target command : %s", infoWrongTargetQuery)
	}
}

//
//
// Error cases
//
//
