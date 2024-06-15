package service

import (
	"app/config"
	"app/model"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type searchService struct {
	elasticTypeClient *elasticsearch.TypedClient
}

type SearchService interface {
	SearchFulltextProduct(filter map[string]string, size int) ([]map[string]interface{}, error)
}

func (s *searchService) SearchFulltextProduct(filter map[string]string, size int) ([]map[string]interface{}, error) {
	var products []map[string]interface{}

	mustQuery := []types.Query{}
	for key, value := range filter {
		switch key {
		case "name":
			for _, c := range strings.Split(value, " ") {
				if len(c) > 0 {
					q := types.Query{
						Fuzzy: map[string]types.FuzzyQuery{
							"name": {
								Value:     c,
								Fuzziness: 2,
							},
						},
					}
					mustQuery = append(mustQuery, q)
				}
			}
		default:
			q := types.Query{
				Match: map[string]types.MatchQuery{
					key: {
						Query: fmt.Sprint(value),
					},
				},
			}
			mustQuery = append(mustQuery, q)
		}
	}

	request := search.Request{
		Query: &types.Query{
			Bool: &types.BoolQuery{
				Must: mustQuery,
			},
		},
		Size: &size,
	}

	res, err := s.elasticTypeClient.
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

	return products, nil
}

func NewSearchService() SearchService {
	return &searchService{
		elasticTypeClient: config.GetElasticTypeClient(),
	}
}
