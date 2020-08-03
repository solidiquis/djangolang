package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/solidiquis/djangolang/config/session"
	"github.com/solidiquis/djangolang/lib"

	"github.com/justinas/alice"
)

var (
	// StandardMiddleware ...middleware for all requests
	StandardMiddleware alice.Chain

	// DynamicMiddleware ...conditional middleware
	DynamicMiddleware alice.Chain
)

func init() {
	StandardMiddleware = alice.New(
		SecureHeaders,
		LogRequest,
		RecoverPanic,
	)

	DynamicMiddleware = alice.New(
		session.Session.Enable,
	)
}

// SecureHeaders ...headers for security
func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")
		next.ServeHTTP(w, r)
	})
}

// LogRequest ...request logs
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

// RecoverPanic ...Delivers 500 error if something goes wrong
func RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				lib.ServerError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
