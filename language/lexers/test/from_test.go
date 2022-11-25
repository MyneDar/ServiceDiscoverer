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

var query = []string{
	"FROM",
	"Human.Add",
}

func TestFromProcessGood(t *testing.T) {
	fromLex := &lexers.FromLex{}
	err, got := fromLex.Process(&query)
	if l := len(got); l != 3 {
		t.Errorf("Wrong token number, got: %d, want 3. Err: %s", l, err)
	}
}

//
//
//

var queryMoreKeyWords = []string{
	"FROM",
	"Human.Add",
	"SELECT",
	"INSERT",
}

func TestFromProcessMoreKeyWords(t *testing.T) {
	fromLex := &lexers.FromLex{}
	_, _ = fromLex.Process(&queryMoreKeyWords)
	if len(queryMoreKeyWords) != 2 {
		t.Errorf("Wrong sliced command, got %s", queryMoreKeyWords)
	}
}

//
//
//

var wrongCommandWord = []string{
	"ASD",
	"Human.Add",
}

func TestFromProcessWrongCommand(t *testing.T) {
	fromLex := &lexers.FromLex{}
	err, got := fromLex.Process(&wrongCommandWord)
	if err != nil && len(got) != 0 {
		t.Errorf("No error on wrong command : %s", wrongCommandWord)
	}
}

//
//
//

var wrongTarget = []string{
	"FROM",
	"HumanAdd",
}

func TestFromProcessWrongTaget(t *testing.T) {
	fromLex := &lexers.FromLex{}
	err, got := fromLex.Process(&wrongTarget)
	if err != nil && len(got) != 0 {
		t.Errorf("No error on target command : %s", wrongTarget)
	}
}

//
//
// Error cases
//
//
