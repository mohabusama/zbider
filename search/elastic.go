package search

import (
	"encoding/json"
	"errors"
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	elastic "gopkg.in/olivere/elastic.v5"
	"sync"
)

const (
	elasticSource = "elasticsearch"
)

func getElasticClient(es_url string) (*elastic.Client, error) {
	log.Info("ES URL ", es_url)

	client, err := elastic.NewClient(
		elastic.SetURL(es_url),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	if err != nil {
		return nil, err
	}

	ping, _, err := client.Ping(es_url).Do(context.TODO())

	log.Info(ping.ClusterName)

	return client, nil
}

type Elasticsearch struct {
	mu        sync.Mutex
	connected bool
	url       string
	index     string
	client    *elastic.Client
}

func NewElasticsearch(es_url string, es_index string) *Elasticsearch {
	client, err := getElasticClient(es_url)
	connected := true

	if err != nil {
		connected = false
	}

	return &Elasticsearch{
		connected: connected,
		url:       es_url,
		index:     es_index,
		client:    client,
	}
}

// Are we interested in this search query. Always interested!
func (e *Elasticsearch) Interested(q string) (bool, int64, error) {
	return true, 1000, nil
}

func (e *Elasticsearch) ready() bool {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.connected != true {
		var err error
		e.client, err = getElasticClient(e.url)
		if err != nil {
			return false
		}
	}

	return true
}

func (e *Elasticsearch) Search(q string) (ResultArray, int64, error) {
	if e.ready() != true {
		return nil, 0, errors.New("Elasticsearch client is not connected!")
	}
	// query
	query := elastic.NewSimpleQueryStringQuery(q)

	res, err := e.client.Search(e.index).
		Query(query).
		Do(context.TODO())

	if err != nil {
		return nil, 0, err
	}

	log.Debug("ES query took: ", res.TookInMillis)

	totalHits := res.TotalHits()
	results := ResultArray{}

	for _, hit := range res.Hits.Hits {
		var j Json

		err := json.Unmarshal(*hit.Source, &j)

		if err != nil {
			log.Infof("Failed to unmarshal: ", err)
		}

		item := NewResultItem(j, *hit.Score, elasticSource)

		results = append(results, item)
	}

	return results, totalHits, nil
}
