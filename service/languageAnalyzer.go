package service

import (
	"net/http"
	"servicediscoverer/interfaces"
	"servicediscoverer/language/parsers"
	"servicediscoverer/models"
)

type LanguageAnalyzer struct {
	analyzerParsers map[models.ServiceToken]interfaces.Parser
	information     map[string]interface{}
}

func NewLanguageAnalyzer() *LanguageAnalyzer {
	var analyzerParserMap = map[models.ServiceToken]interfaces.Parser{
		models.FROM:   parsers.NewFromParser(),   //Get paths
		models.INFO:   parsers.NewInfoParser(),   // Info request
		models.DELETE: parsers.NewDeleteParser(), //type decision
		models.INSERT: parsers.NewInsertParser(), //type decision
		models.UPDATE: parsers.NewUpdateParser(), //type decision
		models.SELECT: parsers.NewSelectParser(), //type decision + filter
		models.WHERE:  parsers.NewWhereParser(),  //type decision(?) + filter
	}
	var info = make(map[string]interface{})
	return &LanguageAnalyzer{analyzerParsers: analyzerParserMap, information: info}
}

func (l *LanguageAnalyzer) TokenProcess(tokens map[models.ServiceToken][]models.TokenStruct, json map[string]interface{}) (err error, response *http.Response) {
	//get data to call
	for key, value := range tokens {
		err = l.analyzerParsers[key].Process(value, l.information)
		if err != nil {
			return err, nil
		}

	}

	//endpoint call if needed

	//Filtering if needed

	return nil, nil
}
