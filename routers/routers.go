package routers

import (
	"app/controller"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func Routers() http.Handler {
	app := chi.NewRouter()

	app.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	searchProductController := controller.NewSearchController()

	app.Route("/search/api/v1", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			res := map[string]interface{}{
				"mess": "done",
			}
			render.JSON(w, r, res)
		})

		r.Route("/product", func(product chi.Router) {
			product.Get("/search", searchProductController.SearchProduct)
		})
	})

	log.Println("Sevice h-shop-be-search starting success!")

	return app
}
