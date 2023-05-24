package service

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

type Right string

const (
	None      Right = ""
	User      Right = "user"
	Developer Right = "developer"
	Admin     Right = "admin"
)

func (r Right) String() string {
	return string(r)
}

func BasicAuth(r *http.Request) (bool, Right) {
	// Get the Basic Authentication credentials
	username, password, ok := r.BasicAuth()
	if ok {
		usernameHash := sha256.Sum256([]byte(username))
		passwordHash := sha256.Sum256([]byte(password))
		expectedUserNameHash := sha256.Sum256([]byte("admin"))
		expectedPasswordHash := sha256.Sum256([]byte("admin"))

		usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUserNameHash[:]) == 1
		passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

		if usernameMatch && passwordMatch {
			return true, Admin
		}
	}

	return false, None
}
