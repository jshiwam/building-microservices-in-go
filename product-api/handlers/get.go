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
	p.log.Info("ListProducts called")
	currency := req.URL.Query().Get("currency")
	product_data, err := p.pdb.GetProducts(currency)
	if err != nil {
		http.Error(rw, "Unable to GetProducts", http.StatusInternalServerError)
		return
	}
	err = data.ToJson(product_data, rw)
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
	p.log.Info("GetProductById called")
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	currency := req.URL.Query().Get("currency")
	product_data, err := p.pdb.GetProduct(id, currency)

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
