package data

import (
	"fmt"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// The ID for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"`

	// The name of the product
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// The description of the product
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// The price of the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	// The SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"required,sku"`
}
type Products []*Product

var ErrProductNotFound = fmt.Errorf("Product not found")

func UpdateProduct(product *Product, id int) error {
	index, err := GetIdIndex(id)
	if err != nil {
		return ErrProductNotFound
	}
	product.ID = id
	productList[index] = product
	return nil
}
func GetProducts() Products {
	return productList
}

func GetProduct(id int) (*Product, error) {
	index, err := GetIdIndex(id)

	if err != nil {
		return nil, err
	}
	return productList[index], nil
}

func RemoveProduct(id int) error {
	index, err := GetIdIndex(id)

	if err != nil {
		return err
	}
	productList = append(productList[:index], productList[index+1:]...)
	return nil
}
func AppendProduct(product *Product) {
	product.ID = generateID()
	productList = append(productList, product)
}

func GetIdIndex(id int) (int, error) {
	for index, element := range productList {
		if element.ID == id {
			return index, nil
		}
	}
	return -1, ErrProductNotFound
}

func generateID() int {
	pid := productList[len(productList)-1].ID
	return pid + 1
}

var productList = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
	},
}
