package order

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/io-m/ddd/aggregate"
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

// Configuration for initializing products in our in-memory DB
// When we instantiate OrderService in test file we can just call NewOrderService(WithCustomerMemoRepo(), WithInitProductMemoRepo())
func WithInitProductMemoRepo(products []aggregate.Product) OrderConfiguration {
	pr := pm.New()
	for _ , p := range products {
		if err := pr.Add(p); err != nil {
			return nil
		}
	}
	return WithIProductRepo(pr)
}

/* Setting kinda custom logger
 * output will look something like:
	LOGGER :: {
		"level": "INFO ",
		"msg": "Customer c34c1d6d-0263-4641-a505-de41f9b92ea5 ordered 1 products",
		"time": "2023-03-05 15:53:24.705896 +0100 CET m=+0.001130621"
	}
*/ 
func WithStandardLogger(serviceName string) OrderConfiguration {
	sl := log.Default()
	sl.SetPrefix(fmt.Sprintf("%s :: ", serviceName))
	sl.SetFlags(0)
	return WithAnyLogger(sl)
}

/* Set of methods on OrderService */

func (os *OrderService) CreateOrder(customerId uuid.UUID, productIds []uuid.UUID) (float64, error) {
	// Fetch the customer through the customer repo
	c, err := os.customerRepo.GetOne(customerId)
	if err != nil {
		return 0, err
	}
	// Get each product
	var products []aggregate.Product
	var total float64

	for _, id := range productIds {
		p, err := os.productRepo.GetOne(id)
		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetProductPrice()
	}
	
	l, err := logOut("INFO ", fmt.Sprintf("Customer %s ordered %d products", c.GetCustomerId(), len(products)))
	if err != nil {
		return 0, err
	}
	os.logger.Printf(l)
	return total, nil
}

func logOut(level, message string) (string, error){
	logStructure := make(map[string]string)
	logStructure["level"] = level
	logStructure["time"] = time.Now().String()
	logStructure["msg"] = message
	b, err := json.MarshalIndent(logStructure, "", " ")
	if err != nil {
		return "Could not return json format", err
	}
	return string(b), nil
}