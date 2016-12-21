package routes

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Service interface {
	AddRoutes(*mux.Router)
}

func GetRouter(services ...Service) http.Handler {
	r := mux.NewRouter()

	// Register routes
	for _, service := range services {
		service.AddRoutes(r)
	}

	router := NewRouter(r)

	// TODO: Add OAUTH2 middleware
	recovery := handlers.RecoveryHandler()
	router.Use(recovery, LoggingMiddleware)

	return router.Handler()
}
