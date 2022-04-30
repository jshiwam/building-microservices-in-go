package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/jshiwam/building-microservices-in-go/product-api/data"
)

type Products struct {
	log hclog.Logger
	v   *data.Validation
	pdb *data.ProductsDB
}

func NewProducts(l hclog.Logger, v *data.Validation, pdb *data.ProductsDB) *Products {
	return &Products{l, v, pdb}
}

type KeyProduct struct{}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductId returns a product Id from the url
// Panics if cannot convert the id into an integer
func getProductId(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err)
	}
	return id
}
