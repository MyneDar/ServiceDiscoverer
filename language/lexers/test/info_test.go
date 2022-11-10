package test

import (
	"servicediscoverer/language/lexers"
	"testing"
)

var infoQueryAsterisk = []string{
	"INFO",
	"*",
}

var infoQueryIdentDotAsterisk = []string{
	"INFO",
	"Human.*",
}

var infoQueryIdentDotIdent = []string{
	"INFO",
	"Human.Add",
}

var infoQueryMoreKeyWords = []string{
	"INFO",
	"Human.Add",
	"SELECT",
	"INSERT",
}

var infoWrongCommandWord = []string{
	"ASD",
	"Human.Add",
}

var infoWrongTarget = []string{
	"INFO",
	"HumanAdd",
}

func TestInfoProcessGoodAsterisk(t *testing.T) {
	infoLex := &lexers.InfoLex{}
	err, got := infoLex.Process(&infoQueryAsterisk)
	if l := len(got); l != 2 {
		t.Errorf("Wrong token number, got: %d, want 3. Err: %s", l, err)
	}
}

func TestInfoProcessGoodIdentDotAsterisk(t *testing.T) {
	infoLex := &lexers.InfoLex{}
	err, got := infoLex.Process(&infoQueryIdentDotAsterisk)
	if l := len(got); l != 3 {
		t.Errorf("Wrong token number, got: %d, want 3. Err: %s", l, err)
	}
}

func TestInfoProcessGoodIdentDotIdent(t *testing.T) {
	infoLex := &lexers.InfoLex{}
	err, got := infoLex.Process(&infoQueryIdentDotIdent)
	if l := len(got); l != 3 {
		t.Errorf("Wrong token number, got: %d, want 3. Err: %s", l, err)
	}
}

func TestInfoProcessMoreKeyWords(t *testing.T) {
	infoLex := &lexers.InfoLex{}
	_, _ = infoLex.Process(&infoQueryMoreKeyWords)
	if len(infoQueryMoreKeyWords) != 2 {
		t.Errorf("Wrong sliced command, got %s", infoQueryMoreKeyWords)
	}
}

func TestInfoProcessWrongCommand(t *testing.T) {
	infoLex := &lexers.InfoLex{}
	err, got := infoLex.Process(&infoWrongCommandWord)
	if err != nil && len(got) != 0 {
		t.Errorf("No error on wrong command : %s", infoWrongCommandWord)
	}
}
func TestInfoProcessWrongTaget(t *testing.T) {
	infoLex := &lexers.InfoLex{}
	err, got := infoLex.Process(&infoWrongTarget)
	if err != nil && len(got) != 0 {
		t.Errorf("No error on target command : %s", infoWrongTarget)
	}
}
