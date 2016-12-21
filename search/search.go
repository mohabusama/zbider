package search

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Service struct {
	es_client *Elasticsearch
}

func NewService(es_url string, index string) *Service {
	return &Service{
		es_client: NewElasticsearch(es_url, index),
	}
}

func (s *Service) AddRoutes(r *mux.Router) {
	r.Methods("Get").Path("/search").HandlerFunc(s.Search)
}

func (s *Service) Search(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")

	if q == "" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	log.Debug("Searching: ", q)

	start := time.Now()

	var totalHits int64 = 0

	searchResult := NewSearchResult()

	// Loop on all *searchers*
	s.es_client.Interested(q)

	// TODO: determine which searcher to use

	// Search!
	results, total, err := s.es_client.Search(q)
	if err != nil {
		panic(err)
	}

	totalHits = totalHits + total

	searchResult.Add(results)

	elapsed := time.Since(start)
	searchResult.Took = elapsed.Seconds()
	searchResult.TotalHits = totalHits

	err = json.NewEncoder(w).Encode(&searchResult)
	if err != nil {
		panic(err)
	}
}
