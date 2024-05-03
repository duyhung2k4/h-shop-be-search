package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func main() {
	cert, _ := os.ReadFile("cert/http_ca.crt")
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Password: "oouLuchH8ymSYzie_+Fs",
		Username: "elastic",
		CACert:   cert,
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	query := types.Query{
		Bool: &types.BoolQuery{},
	}

	res, _ := es.Search().Index("product_index").Request(&search.Request{
		Query: &query,
	}).Do(context.Background())

	hit := res.Hits.Hits[0]

	dataByte, _ := hit.Source_.MarshalJSON()
	data := map[string]interface{}{}

	json.Unmarshal(dataByte, &data)
	data["_id"] = hit.Id_

	log.Println(data)
}
