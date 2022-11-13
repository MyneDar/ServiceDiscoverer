package test

import (
	"servicediscoverer/language/lexers"
	"testing"
)

var query = []string{
	"FROM",
	"Human.Add",
}

var queryMoreKeyWords = []string{
	"FROM",
	"Human.Add",
	"SELECT",
	"INSERT",
}

var wrongCommandWord = []string{
	"ASD",
	"Human.Add",
}

var wrongTarget = []string{
	"FROM",
	"HumanAdd",
}

func TestFromProcessGood(t *testing.T) {
	fromLex := &lexers.FromLex{}
	err, got := fromLex.Process(&query)
	if l := len(got); l != 3 {
		t.Errorf("Wrong token number, got: %d, want 3. Err: %s", l, err)
	}
}

func TestFromProcessMoreKeyWords(t *testing.T) {
	fromLex := &lexers.FromLex{}
	_, _ = fromLex.Process(&queryMoreKeyWords)
	if len(queryMoreKeyWords) != 2 {
		t.Errorf("Wrong sliced command, got %s", queryMoreKeyWords)
	}
}

func TestFromProcessWrongCommand(t *testing.T) {
	fromLex := &lexers.FromLex{}
	err, got := fromLex.Process(&wrongCommandWord)
	if err != nil && len(got) != 0 {
		t.Errorf("No error on wrong command : %s", wrongCommandWord)
	}
}
func TestFromProcessWrongTaget(t *testing.T) {
	fromLex := &lexers.FromLex{}
	err, got := fromLex.Process(&wrongTarget)
	if err != nil && len(got) != 0 {
		t.Errorf("No error on target command : %s", wrongTarget)
	}
}
