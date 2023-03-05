package memo

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/io-m/ddd/aggregate"
	"github.com/io-m/ddd/domain/customer"
)

func TestMemo_GetCustomer(t *testing.T) {
	type testCase struct {
		test          string
		id            uuid.UUID
		expectedError error
	}

	// Create a fake customer to add to repository
	c, err := aggregate.NewCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}
	customerId := c.GetCustomerId()
	// Create the repo to use, and add some test Data to it for testing
	// Skip Factory for this
	repo := MemoRepo{
		customers: map[uuid.UUID]aggregate.Customer{
			customerId: c,
		},
	}

	testCases := []testCase{
		{
			test:          "no customer by id",
			id:            uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedError: customer.ErrorCustomerNotFound,
		},
		{
			test:          "customer by id",
			id:            customerId,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		testFunc := func(t *testing.T) {
			_, err := repo.GetOne(tc.id)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}
		}
		t.Run(tc.test, testFunc)
	}
}
