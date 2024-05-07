package main

import (
	"app/config"
	"app/routers"
	"log"
	"net/http"
	"sync"
	"time"
)

// @title Swagger Search Product
// @version 1.0
// @description Swagger Search Product

// @host localhost:18888
// @BasePath /search-product/api/v1
func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		server := http.Server{
			Addr:           ":" + config.GetAppPort(),
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
			Handler:        routers.Routers(),
		}

		log.Fatalln(server.ListenAndServe())
		wg.Done()
	}()

	wg.Wait()
}
