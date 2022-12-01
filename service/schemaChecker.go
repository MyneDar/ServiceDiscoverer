package service

import (
	"errors"
	"github.com/samber/lo"
	"servicediscoverer/models"
)

var ruleMap = map[models.ServiceToken][]models.ServiceToken{
	models.INFO:   {},
	models.FROM:   {models.SELECT, models.DELETE, models.INSERT, models.UPDATE},
	models.SELECT: {models.WHERE},
	models.WHERE:  {},
	models.DELETE: {},
	models.INSERT: {},
	models.UPDATE: {},
}

// INFO FROM külső->
func SchemaCheck(GetTokensMap map[models.ServiceToken][]models.TokenStruct) (err error) {
	keys := lo.Keys[models.ServiceToken, []models.TokenStruct](GetTokensMap)
	initPossibilities := []models.ServiceToken{models.INFO, models.FROM} //Default command begins
	targetState := models.ILLEGAL
	for _, value := range initPossibilities {
		if keys[0] == value {
			targetState = keys[0]
			break
		}
	}
	if targetState != models.ILLEGAL {
		return iterSchemaCheck(keys[1:], ruleMap[targetState])
	} else {
		err = errors.New("bad Schema")
		return err
	}
}

func iterSchemaCheck(keys []models.ServiceToken, possibilities []models.ServiceToken) (err error) {
	err = errors.New("bad Schema")
	if (len(possibilities) == 0 && len(keys) == 0) || (len(possibilities) != 0 && len(keys) == 0) {
		return nil
	} else if len(possibilities) == 0 && len(keys) != 0 {
		return err
	}
	targetState := models.ILLEGAL
	for _, value := range possibilities {
		if keys[0] == value {
			targetState = keys[0]
			break
		}
	}
	if targetState != models.ILLEGAL {
		return iterSchemaCheck(keys[1:], ruleMap[targetState])
	} else {
		return err
	}
}
