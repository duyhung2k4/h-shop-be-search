package config

import (
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

func connectElastic() error {
	var errElastic error

	cert, err := os.ReadFile("cert/http_ca.crt")
	if err != nil {
		return err
	}
	cfg := elasticsearch.Config{
		Addresses: []string{
			urlElasticSearch,
		},
		Password: passwordElasticSearch,
		Username: userElasticSearch,
		CACert:   cert,
	}

	elasticClient, errElastic = elasticsearch.NewTypedClient(cfg)

	return errElastic
}
