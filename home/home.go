package home

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) AddRoutes(r *mux.Router) {
	r.Methods("Get").Path("/").HandlerFunc(s.Home)
}

func (s *Service) Home(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(w, "WELCOME HOME")
}
