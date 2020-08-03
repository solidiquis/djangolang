package controller

import (
	"net/http"

	"github.com/solidiquis/djangolang/config/session"

	// Using later
	_ "github.com/solidiquis/djangolang/lib"
)

// Handler ...test
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n"))
}

// PutSessionData ...write session data
func PutSessionData(w http.ResponseWriter, r *http.Request) {
	session.Session.Put(r, "foo", "bar")
	w.Write([]byte("Data set. You also got a coookie."))
}

// ReadSessionData ...read session data
func ReadSessionData(w http.ResponseWriter, r *http.Request) {
	message := session.Session.GetString(r, "foo")
	w.Write([]byte(message))
}
