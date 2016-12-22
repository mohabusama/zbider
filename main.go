package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/mohabusama/zbider/home"
	"github.com/mohabusama/zbider/routes"
	"github.com/mohabusama/zbider/search"
	"net/http"
	"os"
)

const (
	version = "0.1-alpha"
)

func main() {
	app := cli.NewApp()
	app.Name = "Zbider"
	app.Version = version

	var verbose bool
	var elasticsearch, elasticsearch_index, bindAddress string

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Verbose output",
			Destination: &verbose,
		},
		cli.StringFlag{
			Name:        "bind, b",
			Usage:       "Bind address",
			Value:       ":8080",
			Destination: &bindAddress,
		},
		cli.StringFlag{
			Name:        "elasticsearch, e",
			Usage:       "Elasticsearch server URL",
			Value:       "http://localhost:9200",
			Destination: &elasticsearch,
		},
		cli.StringFlag{
			Name:        "index, i",
			Usage:       "Elasticsearch Index",
			Value:       "zbider-index-dev",
			Destination: &elasticsearch_index,
		},
	}

	app.Before = func(c *cli.Context) error {
		if verbose {
			log.SetLevel(log.DebugLevel)
			log.Debug("Verbose output")
			log.Debug(app.Name, "-", app.Version)
		}
		return nil
	}

	app.Action = func(c *cli.Context) {
		// Load and register Services routes
		homeService := home.NewService()
		searchService := search.NewService(elasticsearch, elasticsearch_index)

		log.Debug("Registering router")
		router := routes.GetRouter(searchService, homeService)

		log.Infof("Starting %s server: %s", app.Name, bindAddress)
		log.Fatal(http.ListenAndServe(bindAddress, router))
	}

	app.Run(os.Args)
}
