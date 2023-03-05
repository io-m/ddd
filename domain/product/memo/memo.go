package memo

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/io-m/ddd/aggregate"
	"github.com/io-m/ddd/domain/product"
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
	p, ok := mp.products[id]
	if !ok {
		return aggregate.Product{}, errors.New(product.ErrorProductNotFound)
	}
	return p, nil
}

func (mp *MemoProductRepo) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product

	for _, product := range mp.products {
		products = append(products, product)
	}
	return products, nil
}
func (mp *MemoProductRepo) Add(p aggregate.Product) error {
	id := p.GetProductId()
	if _, ok := mp.products[id]; ok {
		return nil
	}
	mp.Lock()
	mp.products[id] = p
	mp.Unlock()
	return nil
}
func (mp *MemoProductRepo) Update(p aggregate.Product) error {
	mp.Lock()
	defer mp.Unlock()
	id := p.GetProductId()
	_, ok := mp.products[id]
	if !ok {
		mp.Add(p)
	}
	mp.products[id] = p
	return nil
}
func (mp *MemoProductRepo) Delete(id uuid.UUID) error {
	mp.Lock()
	defer mp.Unlock()
	if _, ok := mp.products[id]; !ok {
		return nil
	}
	delete(mp.products, id)
	return nil
}
