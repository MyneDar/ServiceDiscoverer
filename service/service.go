package service

import (
	"encoding/json"
	"log"
	"net/http"
	"servicediscoverer/dev"
	"servicediscoverer/ent"
	"servicediscoverer/models"
)

// Service is a struct that contains a tokenizer and a language analyzer.
// It handles the incoming data from the clients end and gives back responses.
// It handles the data management of the services and the endpoints as well
// It handles the authentication and the authorization of the clients and the server side users as well.
type Service struct {
	tokenizer        *Tokenizer
	languageAnalyzer *LanguageAnalyzer
}

// Init is a function that returns a new Service
func (s *Service) Init() {
	err := dev.EntClientInit()
	if err != nil {
		log.Fatal(err)
	}
	s.tokenizer = NewTokenizer()
	s.languageAnalyzer = NewLanguageAnalyzer()
}

// GetDataHandler is a function that handles the incoming command and data from the clients and
// gives back responses bases on the command and data.
func (s *Service) GetDataHandler(w http.ResponseWriter, r *http.Request) {
	//Check if the method is POST else return error.
	if r.Method == http.MethodPost {
		//Authentication
		authorized, accessRight := BasicAuth(r)
		if !authorized {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}

		//get command from JSON
		var command string
		var Json map[string]interface{}

		//Decode JSON
		err := json.NewDecoder(r.Body).Decode(&command)
		if err != nil {
			log.Printf("JSON decode error: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//Tokenize command
		err, tokensMap := s.tokenizer.CommandProcess(command)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if tokensMap[models.INFO] != nil && !(accessRight == Admin || accessRight == Developer) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
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

// Todo: Implement
// registerProviderHandler is a function that handles the incoming service registration request
func registerProviderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		//Authentication
		authorized, accessRight := BasicAuth(r)
		if !authorized || accessRight != Admin {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}

		//Process incoming data
		//Under a name, only one can be
		var provider ent.ProviderRegisterData

		err := json.NewDecoder(r.Body).Decode(&provider)
		if err != nil {
			log.Printf("JSON decode error: %v", err)
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

// registerEndpointHandler is a function that handles the incoming endpoint registration request
func registerEndpointHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication
	authorized, accessRight := BasicAuth(r)
	if !authorized || accessRight != Admin {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	//Process incoming data

	//Safety checks

	//Put it into the database and error handling

	//Give back message about success
}

// updateProviderHandler is a function that updates services
func updateProviderHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication
	authorized, accessRight := BasicAuth(r)
	if !authorized || accessRight != Admin {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	//Process incoming data

	//Safety checks

	//Update it in the database and error handling

	//Give back message about success
}

// updateEndpointHandler is a function that updates endpoints
func updateEndpointHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication
	authorized, accessRight := BasicAuth(r)
	if !authorized || accessRight != Admin {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	//Process incoming data

	//Safety checks

	//Update it in the database and error handling

	//Give back message about success
}

// deleteProviderHandler is a function that deletes services
func deleteProviderHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication
	authorized, accessRight := BasicAuth(r)
	if !authorized || accessRight != Admin {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	//Process incoming data

	//Safety checks

	//Delete it from the database and error handling

	//Give back message about success
}

// deleteEndpointHandler is a function that deletes endpoints
func deleteEndpointHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication
	authorized, accessRight := BasicAuth(r)
	if !authorized || accessRight != Admin {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	//Process incoming data

	//Safety checks

	//Delete it from the database and error handling

	//Give back message about success
}
