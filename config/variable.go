package config

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-chi/jwtauth"
	"github.com/redis/go-redis/v9"
)

const (
	APP_PORT                = "APP_PORT"
	URL_ELASTIC_SEARCH      = "URL_ELASTIC_SEARCH"
	USER_ELASTIC_SEARCH     = "USER_ELASTIC_SEARCH"
	PASSWORD_ELASTIC_SEARCH = "PASSWORD_ELASTIC_SEARCH"
	URL_REDIS               = "URL_REDIS"
)

var (
	appPort               string
	urlElasticSearch      string
	userElasticSearch     string
	passwordElasticSearch string
	urlRedis              string

	elasticTypeClient *elasticsearch.TypedClient
	elasticClient     *elasticsearch.Client
	rdb               *redis.Client
	jwt               *jwtauth.JWTAuth
)
