package routes

import (
	"net/http"

	m "github.com/solidiquis/djangolang/config/middleware"
	c "github.com/solidiquis/djangolang/controllers"

	"github.com/bmizerany/pat"
)

// Routes ... app routes
func Routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(c.Handler))
	mux.Get("/foo", m.DynamicMiddleware.ThenFunc(c.PutSessionData))
	mux.Get("/bar", m.DynamicMiddleware.ThenFunc(c.ReadSessionData))

	return m.StandardMiddleware.Then(mux)
}
