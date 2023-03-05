package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/io-m/ddd/entity"
)

var (
	ErrorMissingValues = errors.New("products is missing name or price values")
)

type Product struct {
	// here we embed root entity item
	item     *entity.Item
	price    float64
	quantity int
}

func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrorMissingValues
	}
	return Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 1,
	}, nil
}

/* Getters and setters for Product aggregate */

func (p Product) GetProductId() uuid.UUID {
	return p.item.ID
}

func (p Product) GetProductItem() *entity.Item {
	return p.item
}

func (p Product) GetProductPrice() float64 {
	return p.price
}

func (p Product) GetProductQuantity() int {
	return p.quantity
}
