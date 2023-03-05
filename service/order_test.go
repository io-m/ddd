package order

import (
	"testing"

	"github.com/google/uuid"
	"github.com/io-m/ddd/aggregate"
)

func init_products(t *testing.T) []aggregate.Product {
	
	beer, _ := aggregate.NewProduct("beer", "Tasty", 12.45)
	wine, _ := aggregate.NewProduct("wine", "Ruby", 24.45)
	shoes, _ := aggregate.NewProduct("shoes", "Comfy", 122.45)
	return []aggregate.Product{beer, wine, shoes}
}

func TestOrder_NewOrder(t *testing.T) {
	// Create new list of products and insert them into in-memory store WithInitProductMemoRepo(products)
	products := init_products(t)
	os, err := NewOrderService(WithCustomerMemoRepo(), WithInitProductMemoRepo(products), WithStandardLogger())
	if err != nil {
		t.Error(err)
	}
	// Add customer
	c, err := aggregate.NewCustomer("Joseph")
	if err != nil {
		t.Error(err)
	}
	if err = os.customerRepo.Add(c); err != nil {
		t.Error(err)
	}
	// Perform Order for one beer
	productIds := []uuid.UUID{
		products[0].GetProductId(),
	}
	_,  err = os.CreateOrder(c.GetCustomerId(), productIds)
	if err != nil {
		t.Error(err)
	}
}