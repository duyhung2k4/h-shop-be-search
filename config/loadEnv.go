package config

import (
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	appPort = os.Getenv(APP_PORT)
	urlElasticSearch = os.Getenv(URL_ELASTIC_SEARCH)
	userElasticSearch = os.Getenv(USER_ELASTIC_SEARCH)
	passwordElasticSearch = os.Getenv(PASSWORD_ELASTIC_SEARCH)
	urlRedis = os.Getenv(URL_REDIS)

	return nil
}
