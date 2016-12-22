package home

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const PATH = "./ui/dist"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) AddRoutes(r *mux.Router) {
	// r.Methods("Get").Path("/").HandlerFunc(s.Home)
	// r.Methods("Get").Path("/").Handler(http.FileServer(http.Dir("./ui/dist")))
	r.Methods("Get").PathPrefix("/").Handler(http.FileServer(http.Dir(PATH)))
}

func (s *Service) Home(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(w, "WELCOME HOME")
}
