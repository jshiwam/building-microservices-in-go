package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jshiwam/building-microservices-in-go/product-api/data"
)

type Products struct {
	_log *log.Logger
	v    *data.Validation
}

func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
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
