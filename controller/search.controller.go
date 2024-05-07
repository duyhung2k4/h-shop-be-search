package controller

import (
	"app/service"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

type searchController struct {
	searchProductService service.SearchService
}

type SearchController interface {
	SearchProduct(w http.ResponseWriter, r *http.Request)
}

// @Summary      Search Product with name
// @Description  Search product with name
// @Tags         Search Product
// @Accept       json
// @Produce      json
// @Param        name   path      string  true  "Name product"
// @Success      200  {object}  Response
// @Router       /product/search [get]
func (c *searchController) SearchProduct(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	nameProduct := params.Get("name")
	sizeString := params.Get("size")

	var size int = 0
	if sizeString == "" {
		size = 20
	} else {
		var errConvertSize error
		size, errConvertSize = strconv.Atoi(sizeString)
		if errConvertSize != nil {
			internalServerError(w, r, errConvertSize)
			return
		}
	}

	products, err := c.searchProductService.SearchFulltextProduct(nameProduct, size)
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    products,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func NewSearchController() SearchController {
	return &searchController{
		searchProductService: service.NewSearchService(),
	}
}
