package memo

import (
	"sync"

	"github.com/google/uuid"
	"github.com/io-m/ddd/aggregate"
)

type MemoProductRepo struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoProductRepo {
	return &MemoProductRepo{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mp *MemoProductRepo) GetOne(id uuid.UUID) (aggregate.Product, error) {

	return aggregate.Product{}, nil
}

func (mp *MemoProductRepo) GetAll() ([]aggregate.Product, error) {
	return []aggregate.Product{}, nil
}
func (mp *MemoProductRepo) Add(p aggregate.Product) error {
	return nil
}
func (mp *MemoProductRepo) Update(p aggregate.Product) error {
	return nil
}
