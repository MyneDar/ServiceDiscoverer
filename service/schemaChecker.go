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

// getNextStates is a function, which returns the possible next states of the current state.
func getNextStates(state models.ServiceToken) (nextStates []models.ServiceToken) {
	return ruleMap[state]
}

// generateStepLists is a function, which generates all the possible step lists from the current state.
func generateStepLists(state models.ServiceToken, currentList []models.ServiceToken, results *[][]models.ServiceToken) {
	nextStates := getNextStates(state)

	//Base case: If there are no next states, we reached the final state
	if len(nextStates) == 0 {
		// If there is no next state, then return the current list
		*results = append(*results, currentList)
	}

	// special case: if the current state is SELECT, then return the current list
	if state == models.SELECT {
		*results = append(*results, currentList)
	}

	// Recursive case: explore all the next states

	for _, nextState := range nextStates {
		// If the next state is already in the current list, then skip it
		if contains(currentList, nextState) {
			continue
		}

		// Create a new list with the next state appended to the current list
		newList := make([]models.ServiceToken, len(currentList))
		copy(newList, currentList)
		newList = append(newList, nextState)

		generateStepLists(nextState, newList, results)
	}
}

// contains is a function, which checks if the current state is already in the current list.
func contains(states []models.ServiceToken, state models.ServiceToken) bool {
	for _, value := range states {
		if value == state {
			return true
		}
	}
	return false
}

// GenerateGoodStepLists is a function, which generates all the possible good step lists.
func GenerateGoodStepLists() (results [][]models.ServiceToken) {
	startStates := []models.ServiceToken{models.INFO, models.FROM} // INFO and FROM are the only possible start states

	// Generate all the possible step lists from the start states
	for _, startState := range startStates {
		currentList := []models.ServiceToken{startState}
		generateStepLists(startState, currentList, &results)
	}

	return results
}

// GenerateBadStepLists is a function, which generates all the possible bad step lists.
func GenerateBadStepLists() (badStepLists [][]models.ServiceToken) {

	// For each current state, check if the next state is valid from the current state
	for currentState, _ := range ruleMap {
		for nextState, _ := range ruleMap {
			// If the current state is the same as the next state, then skip it
			if currentState == nextState {
				continue
			}

			// If the next state is not valid from the current state, then append the current state and the next state to the bad step lists
			if !isValidStep(currentState, nextState) {
				badStepList := []models.ServiceToken{currentState, nextState}
				badStepLists = append(badStepLists, badStepList)
			}
		}
	}

	return badStepLists
}

// isValidStep is a function, which checks if the next state is valid from the current state.
func isValidStep(currentState, nextState models.ServiceToken) bool {
	ValidNextStates := getNextStates(currentState)

	// Check if the next state is in the valid next states
	for _, validState := range ValidNextStates {
		if validState == nextState {
			return true
		}
	}
	return false
}
