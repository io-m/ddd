package service

import (
	"log"

	"github.com/google/uuid"
	"github.com/io-m/ddd/domain/customer"
	"github.com/io-m/ddd/domain/customer/memo"
	"github.com/io-m/ddd/domain/product"
	pm "github.com/io-m/ddd/domain/product/memo"
)

// OrderConfiguration is alias type for function that take any order service and works with it
// Since we work on the passed service we have to pass it as a pointer to effectively change things in it
type OrderConfiguration func(orderService *OrderService) error

type OrderService struct {
	customerRepo customer.ICustomerRepo
	productRepo  product.IProductRepo
	logger       *log.Logger
}

// NewOrderService is factory that accepts any order configuration.
// For example, we may pass WithMemoRepo as configuration function and it will apply all MemoRepo functions for ICustomerRepo interface
// If we apply some other ICustomerRepo interface implementation OrderService will apply those functions (NewOrderService(WithMongoRepo))
func NewOrderService(orderConfigs ...OrderConfiguration) (*OrderService, error) {

	orderService := &OrderService{}

	// loop through all the configs passed as param and apply them
	for _, conf := range orderConfigs {
		// since conf is of type func that accepts order service as attribute,
		// we simply call conf with previously created order service
		if err := conf(orderService); err != nil {
			return nil, err
		}
	}
	return orderService, nil
}

// WithICustomerRepo applies a customer repo interface (any repo that implements it) to the OrderService
func WithICustomerRepo(cr customer.ICustomerRepo) OrderConfiguration {
	return func(os *OrderService) error {
		os.customerRepo = cr
		return nil
	}
}

func WithIProductRepo(pr product.IProductRepo) OrderConfiguration {
	return func(os *OrderService) error {
		os.productRepo = pr
		return nil
	}
}

func WithAnyLogger(l *log.Logger) OrderConfiguration {
	return func(os *OrderService) error {
		os.logger = l
		return nil
	}
}

// WithCustomerMemoRepo applies customer memo repo (which implements ICustomerRepo) to the OrderService
func WithCustomerMemoRepo() OrderConfiguration {
	mr := memo.New()
	return WithICustomerRepo(mr)
}

func WithProductMemoRepo() OrderConfiguration {
	pr := pm.New()
	return WithIProductRepo(pr)
}

func WithStandardLogger() OrderConfiguration {
	sl := log.Default()
	return WithAnyLogger(sl)
}

/* Set of methods on OrderService */

// TODO: continue here
func (os *OrderService) CreateOrder(customerId uuid.UUID, productIds []uuid.UUID) error {
	// Fetch the customer through the customer repo
	c, err := os.customerRepo.GetOne(customerId)
	if err != nil {
		return err
	}
	os.logger.Print(c)
	// Get each product
	p, _ := os.productRepo.GetOne(uuid.New())
	os.logger.Print(p)
	return nil
}
