package service

import (
	"app/config"
	"app/model"
	"context"
	"encoding/json"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type searchService struct {
	elasticClient *elasticsearch.TypedClient
}

type SearchService interface {
	SearchFulltextProduct(name string, size int) ([]map[string]interface{}, error)
}

func (s *searchService) SearchFulltextProduct(name string, size int) ([]map[string]interface{}, error) {
	var products []map[string]interface{}

	request := search.Request{
		Query: &types.Query{
			Fuzzy: map[string]types.FuzzyQuery{
				"name": {
					Value:     name,
					Fuzziness: 1,
				},
			},
		},
		Size: &size,
	}

	res, err := s.elasticClient.
		Search().
		Index(model.PRODUCT_INDEX).
		Request(&request).Do(context.Background())

	if err != nil {
		return model.EmptyArrayMapData, err
	}

	for _, item := range res.Hits.Hits {
		var convertData map[string]interface{}
		err := json.Unmarshal(item.Source_, &convertData)
		if err != nil {
			return model.EmptyArrayMapData, err
		}

		convertData["_id"] = item.Id_
		products = append(products, convertData)
	}

	log.Println(res.Hits.Total)

	return products, nil
}

func NewSearchService() SearchService {
	return &searchService{
		elasticClient: config.GetElasticClient(),
	}
}
