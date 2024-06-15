package config

import (
	"app/model"
	"context"
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

	elasticTypeClient, errElastic = elasticsearch.NewTypedClient(cfg)
	elasticClient, errElastic = elasticsearch.NewClient(cfg)
	initIndex(elasticTypeClient)

	return errElastic
}

func initIndex(elastic *elasticsearch.TypedClient) {
	elastic.Indices.Create(model.PRODUCT_INDEX).Do(context.Background())
}
