package session

import (
	"net/http"
	"os"
	"time"

	"github.com/golangcollege/sessions"
)

var (
	// Session ...Session manager
	Session *sessions.Session
)

// InitSession ...initialize Session config
func init() {
	secretA := os.Getenv("SESSIONS_SECRET_ONE")
	secretB := os.Getenv("SESSIONS_SECRET_TWO")
	secretC := os.Getenv("SESSIONS_SECRET_THREE")

	// rotate keys; Session cached
	Session = sessions.New(
		[]byte(secretA),
		[]byte(secretB),
		[]byte(secretC),
	)

	// Session config
	Session.Lifetime = 12 * time.Hour
	Session.Path = "/"
	Session.Persist = false

	// Cookies only sent to our servers if the GET request from other
	// sites cause top level navigation, thus preventing CSRF attacks
	// and eliminating the need for CSRF tokens.
	Session.SameSite = http.SameSiteLaxMode

	if os.Getenv("GO_ENV") == "production" {
		Session.Domain = "djangolang"
		Session.Secure = true
		Session.HttpOnly = true
	} else {
		Session.Domain = "localhost"
		Session.Secure = false
		Session.HttpOnly = false
	}
}
