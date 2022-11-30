package service

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"servicediscoverer/models"
)

type Service struct {
	tokenizer        *Tokenizer
	languageAnalyzer *LanguageAnalyzer
}

func (s *Service) init() {
	s.tokenizer = NewTokenizer()
	s.languageAnalyzer = NewLanguageAnalyzer()
}

func (s *Service) getDataHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication

	if r.Method == http.MethodPost {
		//get command from JSON
		var command string
		var Json map[string]interface{}

		err := json.NewDecoder(r.Body).Decode(&command)
		if err != nil {
			log.Printf("")
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		//Tokenizer

		err, tokensMap := s.tokenizer.CommandProcess(command)
		if err != nil {
			return
		}

		//Schema check and JSon empty check
		ruleMap := map[models.ServiceToken][]models.ServiceToken{
			models.INFO:   {},
			models.FROM:   {models.SELECT, models.DELETE, models.INSERT, models.UPDATE},
			models.SELECT: {models.WHERE},
			models.WHERE:  {},
			models.DELETE: {},
			models.INSERT: {},
			models.UPDATE: {},
		}
		possibilities := []models.ServiceToken{models.INFO, models.FROM}
		for key, _ := range tokensMap {
			goodSchema := false
			if len(possibilities) == 0 {
				break
			}
			for _, value := range possibilities {
				if key == value {
					goodSchema = true
					possibilities = ruleMap[key]
				}
			}
			if !goodSchema {
				return //Todo: Give back Error message
			}
		}

		//Language analyzer
		err, response := s.languageAnalyzer.TokenProcess(tokensMap, Json)
		if err != nil {
			return
		}

		//Send back data
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return
		}

		_, err = w.Write(body)
		if err != nil {
			return
		}
	} else {
		log.Printf("")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func deleteDataHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication

	//get command from JSON

	//Use lexers, error handling

	//Use parsers, error handling

	//Call delete endpoint

	//Give back the message of the endpoint
}

func insertDataHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication

	//get command from JSON

	//Use lexers, error handling

	//Use parsers, error handling

	//Call insert endpoint

	//Give back the message of the endpoint
}

func updateDataHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication

	//get command from JSON

	//Use lexers, error handling

	//Use parsers, error handling

	//Call update endpoint

	//Give back the message of the endpoint
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication

	//get command from JSON

	//Use lexers, error handling

	//Use parsers, error handling

	//Give back the needed information
}

func registerProviderHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication

	//Process incoming data
	//Under a name only one can be

	//Safety checks

	//Put it into the database and error handling

	//Give back message about success
}

func registerEndpointHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication

	//Process incoming data

	//Safety checks

	//Put it into the database and error handling

	//Give back message about success
}

func updateProviderHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication

	//Process incoming data

	//Safety checks

	//Update it in the database and error handling

	//Give back message about success
}

func updateEndpointHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication

	//Process incoming data

	//Safety checks

	//Update it in the database and error handling

	//Give back message about success
}

func deleteProviderHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication

	//Process incoming data

	//Safety checks

	//Delete it from the database and error handling

	//Give back message about success
}

func deleteEndpointHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication

	//Process incoming data

	//Safety checks

	//Delete it from the database and error handling

	//Give back message about success
}
