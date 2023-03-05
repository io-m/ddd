package product

import (
	"github.com/google/uuid"
	"github.com/io-m/ddd/aggregate"
)

type IProductRepo interface {
	GetOne(uuid.UUID) (aggregate.Product, error)
	GetAll() ([]aggregate.Product, error)
	Add(aggregate.Product) error
	Update(aggregate.Product) error
}
