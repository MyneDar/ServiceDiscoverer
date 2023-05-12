package service

import (
	"errors"
	"github.com/samber/lo"
	"servicediscoverer/models"
)

// The ruleMap is a map, where the key is the current state, and the value is the possible next states.
var ruleMap = map[models.ServiceToken][]models.ServiceToken{
	models.INFO:   {},
	models.FROM:   {models.SELECT, models.DELETE, models.INSERT, models.UPDATE},
	models.SELECT: {models.WHERE},
	models.WHERE:  {},
	models.DELETE: {},
	models.INSERT: {},
	models.UPDATE: {},
}

// SchemaCheck is a function, which checks the schema of the command and returns an error if the schema is bad. It validates by the ruleMap.
func SchemaCheck(GetTokensMap map[models.ServiceToken][]models.TokenStruct) (err error) {

	// Get the keys of the map
	keys := lo.Keys[models.ServiceToken, []models.TokenStruct](GetTokensMap)

	//Default command begins with INFO or FROM
	initPossibilities := []models.ServiceToken{models.INFO, models.FROM}

	// Default targetState is ILLEGAL
	targetState := models.ILLEGAL

	// Check the first token
	for _, value := range initPossibilities {
		if keys[0] == value {
			targetState = keys[0]
			break
		}
	}

	// If the first token is in the possibilities, then check the next token else return an error.
	if targetState != models.ILLEGAL {
		return iterSchemaCheck(keys[1:], ruleMap[targetState])
	} else {
		err = errors.New("bad Schema")
		return err
	}
}

// iterSchemaCheck is a recursive function, which checks the schema of the command and returns an error if the schema is bad. It validates by the ruleMap.
func iterSchemaCheck(keys []models.ServiceToken, possibilities []models.ServiceToken) (err error) {

	err = errors.New("bad Schema")

	if (len(possibilities) == 0 && len(keys) == 0) || (len(possibilities) != 0 && len(keys) == 0) {
		return nil
	} else if len(possibilities) == 0 && len(keys) != 0 {
		return err
	}

	// Default targetState is ILLEGAL
	targetState := models.ILLEGAL

	// Check if the current state is in the possibilities
	for _, value := range possibilities {
		if keys[0] == value {
			targetState = keys[0]
			break
		}
	}

	// If the current state is in the possibilities, then check the next state else return an error.
	if targetState != models.ILLEGAL {
		return iterSchemaCheck(keys[1:], ruleMap[targetState])
	} else {
		return err
	}
}
