package test

import (
	"github.com/stretchr/testify/assert"
	"servicediscoverer/language/lexers"
	"servicediscoverer/models"
	"testing"
)

//
//
// Good cases
//
//

var updateQuerySimple = []string{
	"UPDATE",
}

var updateSimpleOutput = []models.TokenStruct{
	{models.UPDATE, "UPDATE"},
}

func TestUpdateProcessSimple(t *testing.T) {
	insertLex := &lexers.UpdateLex{}
	err, got := insertLex.Process(&updateQuerySimple)

	assert.Nilf(t, err, "Error is not nil: %s", err)
	assert.Equal(t, len(updateSimpleOutput), len(got), "Wrong token number")
}

//
//
//

var updateManyArguments = []string{
	"UPDATE",
	"ASD",
}

func TestUpdateProcessManyArguments(t *testing.T) {
	insertLex := &lexers.UpdateLex{}
	_, _ = insertLex.Process(&updateManyArguments)
	if len(updateManyArguments) == 0 {
		t.Errorf("Only one arguments")
	}
}

//
//
// Error cases
//
//
