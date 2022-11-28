package service

import (
	"net/http"
	"servicediscoverer/interfaces"
	"servicediscoverer/language/parsers"
	"servicediscoverer/models"
)

type LanguageAnalyzer struct {
	analyzerParsers map[models.ServiceToken]interfaces.Parser
}

func NewLanguageAnalyzer() *LanguageAnalyzer {
	var analyzerParserMap = map[models.ServiceToken]interfaces.Parser{
		models.FROM:   parsers.NewFromParser(),
		models.INFO:   parsers.NewInfoParser(),
		models.DELETE: parsers.NewDeleteParser(),
		models.INSERT: parsers.NewInsertParser(),
		models.UPDATE: parsers.NewUpdateParser(),
		models.SELECT: parsers.NewSelectParser(),
		models.WHERE:  parsers.NewWhereParser(),
	}
	return &LanguageAnalyzer{analyzerParsers: analyzerParserMap}
}

func (l *LanguageAnalyzer) TokenProcess(tokens map[models.ServiceToken][]models.TokenStruct, json map[string]interface{}) (err error, response *http.Response) {
	//get data to call
	for key, value := range tokens {
		err = l.analyzerParsers[key].Process(value)
		if err != nil {
			return err, nil
		}

	}

	//endpoint call if needed

	//Filtering if needed

	return nil, nil
}
