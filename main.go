package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	// Tạo một kết nối đến Elasticsearch
	cert, _ := os.ReadFile("http_ca.crt")
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Password: "9LHLW67-CFHol8zF3DG1",
		Username: "elastic",
		CACert:   cert,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	// es.Indices.Create("index_1")
	// dataReq := map[string]interface{}{
	// 	"f1": "hihi haha sas",
	// }
	// dataReqByte, _ := json.Marshal(dataReq)
	// es.Index("index_1", bytes.NewReader(dataReqByte))

	// Tạo một index mới
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"f1": "hihi haha sas",
			},
		},
	}
	json.NewEncoder(&buf).Encode(query)

	res, _ := es.Search(
		es.Search.WithIndex("index_1"),
		es.Search.WithBody(strings.NewReader(buf.String())),
	)

	var r map[string]interface{}
	json.NewDecoder(res.Body).Decode(&r)

	for _, h := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		d := h.(map[string]interface{})
		log.Println(d["_source"])
	}
}
