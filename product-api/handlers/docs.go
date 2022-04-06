// Package classification Products API
//
// The purpose of this API is to introduce the endpoints with the help of
// which the user is able to fetch, add, update, and delete the products
//
// Schemes: http
// BasePath: /
// Version: 0.0.1
//
//
// Consumes:
// -application/json
//
// Produces:
// -application/json
//
// swagger:meta
package handlers

import "github.com/jshiwam/building-microservices-in-go/product-api/data"

// Returns a list of products from the database
// swagger:response productsResponse
type productsResponseWrapper struct {
	// The list of products
	// in: body
	Body []data.Product
}

// Returns the product with request id if the product exists in database else returns error
// swagger:response productResponse
type productResponseWrapper struct {
	// The product with requested id
	// in: body
	Body data.Product
}

// swagger:parameters DeleteProduct GetProductById UpdateProduct
type productIdRequiredWrapper struct {
	// The unique id of the product
	// required: true
	// minimum: 1
	// in: path
	ID int `json="id"`
}

// No content is returned by this API endpoint
// swagger:response noContent
type noContentResponseWrapper struct {
}

// A ResponseError is an error that is used when no response is received for the given API endpoint.
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// A ValidationError is an error that is used when the required input fails validation.
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A InternalServerError is an error that is used when the some internal computation fails
// swagger:response internalServerErrorResponse
type internalServerErrorResponse struct {
}
