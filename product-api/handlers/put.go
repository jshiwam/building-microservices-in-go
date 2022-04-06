package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jshiwam/building-microservices-in-go/product-api/data"
)

// swagger:route PUT /products/{id} products UpdateProduct
//
// Updates the product in the database if the given product ID exists
//
// Responses:
// 	201: noContent
//	404: errorResponse
//  422: errorValidation

func (p *Products) UpdateProduct(rw http.ResponseWriter, req *http.Request) {
	p._log.Println("UpdateProduct called")
	product := req.Context().Value(KeyProduct{}).(*data.Product)

	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	err := data.UpdateProduct(product, id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Update ID doesn't exist ", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product Not Found", http.StatusInternalServerError)
		return
	}
}
