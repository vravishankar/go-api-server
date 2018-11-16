package route

import (
	"net/http"

	"../handler"
	mw "../middleware"
	"github.com/gorilla/mux"
)

// Route - Definition
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Auth        bool
}

// Routes - Definition
type Routes []Route

var routes = Routes{
	Route{"Auth", "POST", "/auth/token", handler.Auth, false},
	Route{"Payments", "GET", "/payments", handler.GetAll, false},
	// Route{"Payment", "GET", "/payments/{id}", handler.GetPayment, true},
}

// NewRouter - Set up routers
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true).PathPrefix("/api/v1").Subrouter()

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		if route.Auth {
			mw.Auth(handler)
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
