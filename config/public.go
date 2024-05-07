package config

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-chi/jwtauth"
	"github.com/redis/go-redis/v9"
)

func GetElasticClient() *elasticsearch.TypedClient {
	return elasticClient
}

func GetAppPort() string {
	return appPort
}

func GetJWT() *jwtauth.JWTAuth {
	return jwt
}

func GetRDB() *redis.Client {
	return rdb
}
