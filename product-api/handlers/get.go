package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jshiwam/building-microservices-in-go/product-api/data"
)

// swagger:route GET /products products ListProducts
// Returns a list of products
// responses:
// 	200: productsResponse
//  500: internalServerErrorResponse
func (p *Products) ListProducts(rw http.ResponseWriter, req *http.Request) {
	p._log.Println("ListProducts called")
	product_data := data.GetProducts()
	err := data.ToJson(product_data, rw)
	if err != nil {
		http.Error(rw, "ListProducts cannot convert ToJson ", http.StatusInternalServerError)
		return
	}
}

// swagger:route GET /products/{id} products GetProductById
//
// Fetches the product from the database if the given product ID exists
//
// Responses:
// 	200: productResponse
//	404: errorResponse
//  422: errorValidation

func (p *Products) GetProductById(rw http.ResponseWriter, req *http.Request) {
	p._log.Println("GetProductById called")
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	product_data, err := data.GetProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

	err = data.ToJson(product_data, rw)

	if err != nil {
		http.Error(rw, "Product cannot convert ToJson ", http.StatusInternalServerError)
		return
	}

}
