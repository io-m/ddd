package aggregate_test

import (
	"testing"

	"github.com/io-m/ddd/aggregate"
)

/* NewCustomer factory constructor test cases:
* No name provided
* Happy path
 */
func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test          string
		name          string
		expectedError error
	}

	testCases := []testCase{
		{
			test:          "Empty name validation",
			name:          "",
			expectedError: aggregate.ErrorInvalidPerson,
		},
		{
			test:          "Valid name",
			name:          "Josip Miljak",
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		testFunc := func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)
			// Check if the error matches the expected error
			if err != tc.expectedError {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
		}
		t.Run(tc.test, testFunc)
	}

}
