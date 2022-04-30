package data

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-hclog"
	protos "github.com/jshiwam/building-microservices-in-go/currency/protos/currency"
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
	Price float64 `json:"price" validate:"required,gt=0"`

	// The SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"required,sku"`
}
type Products []*Product

type ProductsDB struct {
	log hclog.Logger
	cc  protos.CurrencyClient
}

func NewProductsDB(log hclog.Logger, cc protos.CurrencyClient) *ProductsDB {
	return &ProductsDB{log: log, cc: cc}
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func (p *ProductsDB) UpdateProduct(product *Product, id int) error {
	index, err := GetIdIndex(id)
	if err != nil {
		return ErrProductNotFound
	}
	product.ID = id
	productList[index] = product
	return nil
}

func (p *ProductsDB) GetProducts(currency string) (Products, error) {
	if currency == "" {
		return productList, nil
	}

	rate, err := p.getRate(currency)

	if err != nil {
		return nil, err
	}

	np := Products{}

	for _, prod := range productList {
		copy_prod := *prod
		copy_prod.Price = rate * copy_prod.Price
		np = append(np, &copy_prod)
	}

	return np, nil
}

func (p *ProductsDB) getRate(dest string) (float64, error) {

	rateRequest := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value[dest]),
	}
	resp, err := p.cc.GetRate(context.Background(), rateRequest)
	if err != nil {
		p.log.Error("Unable to get rate", "currency", dest, "error", err)
		return -1, err
	}
	p.log.Info(fmt.Sprintf("Currency Response %#v", resp.Rate))
	return resp.Rate, err
}
func (p *ProductsDB) GetProduct(id int, currency string) (*Product, error) {
	index, err := GetIdIndex(id)

	if err != nil {
		return nil, err
	}
	np := *productList[index]
	rate, err := p.getRate(currency)

	if err != nil {
		return nil, err
	}
	np.Price = rate * np.Price
	return &np, nil
}

func (p *ProductsDB) RemoveProduct(id int) error {
	index, err := GetIdIndex(id)

	if err != nil {
		return err
	}
	productList = append(productList[:index], productList[index+1:]...)
	return nil
}
func (p *ProductsDB) AppendProduct(product *Product) {
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
