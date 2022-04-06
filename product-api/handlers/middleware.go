package handlers

import (
	"context"
	"net/http"

	"github.com/jshiwam/building-microservices-in-go/product-api/data"
)

// Converts the request payload from Json to Product and validates the
// fields in the product
func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		product := &data.Product{}
		err := data.FromJson(product, r.Body)
		if err != nil {
			p._log.Println("[ERROR] deserializing product", err)

			w.WriteHeader(http.StatusBadRequest)
			data.ToJson(&GenericError{Message: err.Error()}, w)
			return
		}
		errs := p.v.Validate(product)
		if errs != nil {
			// http.Error(w, "Invalid product data", http.StatusInternalServerError)
			// if ve, ok := err.(validator.ValidationErrors); ok {
			// 	for _, fe := range ve {
			// 		p._log.Println(fe.Namespace())
			// 		p._log.Println(fe.Field())
			// 		p._log.Println(fe.StructNamespace())
			// 		p._log.Println(fe.StructField())
			// 		p._log.Println(fe.Tag())
			// 		p._log.Println(fe.ActualTag())
			// 		p._log.Println(fe.Kind())
			// 		p._log.Println(fe.Type())
			// 		p._log.Println(fe.Value())
			// 		p._log.Println(fe.Error())
			// 		p._log.Println(fe.Param())
			// 	}
			p._log.Println("[Error] validating product", errs)
			w.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJson(&ValidationError{Messages: errs.Errors()}, w)
			return
		}
		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		req := r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, req)
	})
}
