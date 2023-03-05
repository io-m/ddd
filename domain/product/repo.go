package product

import (
	"github.com/google/uuid"
	"github.com/io-m/ddd/aggregate"
)

var (
	ErrorProductNotFound = "product is not found"
)

type IProductRepo interface {
	GetOne(uuid.UUID) (aggregate.Product, error)
	GetAll() ([]aggregate.Product, error)
	Add(aggregate.Product) error
	Update(aggregate.Product) error
	Delete(uuid.UUID) error
}
