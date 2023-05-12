package service

import (
	"encoding/json"
	"log"
	"net/http"
	"servicediscoverer/dev"
	"servicediscoverer/ent"
)

type Service struct {
	tokenizer        *Tokenizer
	languageAnalyzer *LanguageAnalyzer
}

func (s *Service) Init() {
	err := dev.EntClientInit()
	if err != nil {
		log.Fatal(err)
	}
	s.tokenizer = NewTokenizer()
	s.languageAnalyzer = NewLanguageAnalyzer()
}

func (s *Service) GetDataHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication

	//Check if the method is POST else return error.
	if r.Method == http.MethodPost {
		//get command from JSON
		var command string
		var Json map[string]interface{}

		//Decode JSON
		err := json.NewDecoder(r.Body).Decode(&command)
		if err != nil {
			log.Printf("")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//Tokenize command
		err, tokensMap := s.tokenizer.CommandProcess(command)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//Schema check and JSon empty check
		err = SchemaCheck(tokensMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//Analyze tokens
		err, response := s.languageAnalyzer.TokenProcess(tokensMap, Json)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//Give back response
		_, err = w.Write(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		log.Printf("")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func registerProviderHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication

	if r.Method == http.MethodPost {
		//Process incoming data
		//Under a name, only one can be
		var provider ent.ProviderRegisterData

		err := json.NewDecoder(r.Body).Decode(&provider)
		if err != nil {
			log.Printf("")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//Safety checks

		//Put it into the database and error handling

		//Give back message about success
	} else {
		log.Printf("")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
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
