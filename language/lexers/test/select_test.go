package test

import (
	"servicediscoverer/language/lexers"
	"testing"
)

var selectQueryAsterisk = []string{
	"SELECT",
	"*",
}

var selectQueryOneIndentOnly = []string{
	"SELECT",
	"Age",
}

var selectQueryMultipleIndents = []string{
	"SELECT",
	"Age",
	"Name",
}

var selectQueryMultipleKeyWords = []string{
	"INFO",
	"*",
	"SELECT",
	"INSERT",
}

func TestSelectProcessGoodAsterisk(t *testing.T) {
	selectLex := &lexers.SelectLex{}
	err, got := selectLex.Process(&selectQueryAsterisk)
	if l := len(got); l != 2 {
		t.Errorf("Wrong token number, got: %d, want 2. Err: %s", l, err)
	}
}

func TestSelectProcessGoodIdent(t *testing.T) {
	selectLex := &lexers.SelectLex{}
	err, got := selectLex.Process(&selectQueryOneIndentOnly)
	if l := len(got); l != 2 {
		t.Errorf("Wrong token number, got: %d, want 2. Err: %s", l, err)
	}
}

func TestSelectProcessMultipleIndents(t *testing.T) {
	selectLex := &lexers.SelectLex{}
	err, got := selectLex.Process(&selectQueryMultipleIndents)
	if l := len(got); l != 3 {
		t.Errorf("Wrong token number, got: %d, want 3. Err: %s", l, err)
	}
}

func TestSelectProcessKeyWords(t *testing.T) {
	selectLex := &lexers.SelectLex{}
	err, got := selectLex.Process(&selectQueryMultipleKeyWords)
	if lgot, lcommand := len(got), len(selectQueryMultipleKeyWords); lgot != 2 && lcommand == 2 {
		t.Errorf("Wrong token number or remain command length, got: %d and %d, want 2 each. Err: %s", lgot, lcommand, err)
	}
}
