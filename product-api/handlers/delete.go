package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jshiwam/building-microservices-in-go/product-api/data"
)

// swagger:route DELETE /products/{id} products DeleteProduct
//
// Deletes the product from the database if the given product ID exists
//
// Responses:
// 	201: noContent
//	404: errorResponse
//  500: internalServerErrorResponse

func (p *Products) DeleteProduct(rw http.ResponseWriter, req *http.Request) {
	p.log.Info("DeleteProduct called")
	vars := mux.Vars(req)
	idInt, _ := strconv.Atoi(vars["id"])
	err := p.pdb.RemoveProduct(idInt)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}
