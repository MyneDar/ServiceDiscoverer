package service

import (
	json2 "encoding/json"
	"servicediscoverer/interfaces"
	"servicediscoverer/language/parsers"
	"servicediscoverer/models"
)

// LanguageAnalyzer is a struct that contains all the parsers for the language in analyzerParsers map
// and gather information from the parsers in an information map
type LanguageAnalyzer struct {
	analyzerParsers map[models.ServiceToken]interfaces.Parser
	information     map[string]interface{}
}

// NewLanguageAnalyzer is a function that returns a pointer to a new LanguageAnalyzer struct
// and set the parsers in the analyzerParsers map
func NewLanguageAnalyzer() *LanguageAnalyzer {
	var analyzerParserMap = map[models.ServiceToken]interfaces.Parser{
		models.FROM:   parsers.NewFromParser(),   //Get paths
		models.INFO:   parsers.NewInfoParser(),   // Info request
		models.DELETE: parsers.NewDeleteParser(), //type decision
		models.INSERT: parsers.NewInsertParser(), //type decision
		models.UPDATE: parsers.NewUpdateParser(), //type decision
		models.SELECT: parsers.NewSelectParser(), //type decision + filter
		models.WHERE:  parsers.NewWhereParser(),  //filter
	}
	var info = make(map[string]interface{})
	return &LanguageAnalyzer{analyzerParsers: analyzerParserMap, information: info}
}

// TokenProcess is a function that calls the parsers in the analyzerParsers map
// and gather information from the parsers in an information map.
// After that it returns the information that is needed or makes a call and give back the response and the response data if needed.
func (l *LanguageAnalyzer) TokenProcess(tokens map[models.ServiceToken][]models.TokenStruct, json map[string]interface{}) (err error, response []byte) {
	//get data to call
	for key, value := range tokens {
		err = l.analyzerParsers[key].Process(value, l.information)
		if err != nil {
			return err, nil
		}

	}

	//endpoint call if needed
	//TODO automatic response data merge end return it (now only INFO returned)
	//TODO data outsource like "Info_Data" etc
	//Filtering if needed
	respData, err := json2.Marshal(l.information["Info_Data"])
	return err, respData
}
