package service

import (
	"encoding/json"
	"log"
	"net/http"
	"servicediscoverer/interfaces"
	"servicediscoverer/language/lexers"
)

type Service struct {
	deleteLex interfaces.Lexer
	fromLex   interfaces.Lexer
	infoLex   interfaces.Lexer
	insertLex interfaces.Lexer
	selectLex interfaces.Lexer
	updateLex interfaces.Lexer
	whereLex  interfaces.Lexer
}

func (s *Service) init() {
	s.deleteLex = &lexers.DeleteLex{}
	s.fromLex = &lexers.FromLex{}
	s.infoLex = &lexers.InfoLex{}
	s.insertLex = &lexers.InsertLex{}
	s.selectLex = &lexers.SelectLex{}
	s.updateLex = &lexers.UpdateLex{}
	s.whereLex = &lexers.WhereLex{}
}

func (s *Service) getDataHandler(w http.ResponseWriter, r *http.Request) {
	//Authentication

	if r.Method == http.MethodPost {
		//get command from JSON
		var command string

		err := json.NewDecoder(r.Body).Decode(&command)
		if err != nil {
			log.Printf("")
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		//Use lexers, error handling

		err, fromTok := s.fromLex.Process(&command)
		if err != nil {
			log.Printf("")
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err, selectTok := s.selectLex.Process(&command)
		if err != nil {
			log.Printf("")
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err, whereTok := s.whereLex.Process(&command)
		if err != nil {
			log.Printf("")
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		//just fillers to not panic
		s := len(fromTok)
		s = len(selectTok)
		s = len(whereTok)
		print(s)

		//From, select, where parser initialization (maybe not here)

		//Use parsers, error handling

		//Call endpoint and process data

		//Send back data
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
