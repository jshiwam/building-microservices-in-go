package handlers

import (
	"net/http"

	"github.com/jshiwam/building-microservices-in-go/product-api/data"
)

// swagger:route POST /products products AddProduct
//
// Adds the product into the database
//
// Responses:
// 	200: productResponse
//	404: errorResponse
//  422: errorValidation

func (p *Products) AddProduct(rw http.ResponseWriter, req *http.Request) {
	p.log.Info("AddProducts called")
	// Read data from request and convert it into a required format
	product := req.Context().Value(KeyProduct{}).(*data.Product)
	p.pdb.AppendProduct(product)
}
