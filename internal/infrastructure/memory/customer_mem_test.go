package memory

import (
	"github.com/google/uuid"
	"go-ddd-tavern/internal/domain/entity"
	"go-ddd-tavern/internal/domain/repository"
	"testing"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	customer, err := entity.NewCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}
	id := customer.GetId()

	repo := MemoryCustomerRepository{
		customers: map[uuid.UUID]entity.Customer{
			id: customer,
		},
	}

	testCases := []testCase{
		{
			name:        "No Customer By ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: repository.ErrCustomerNotFound,
		}, {
			name:        "Customer By ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemory_AddCustomer(t *testing.T) {
	type testCase struct {
		name        string
		cust        string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add Customer",
			cust:        "Percy",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := MemoryCustomerRepository{
				customers: map[uuid.UUID]entity.Customer{},
			}

			customer, err := entity.NewCustomer(tc.cust)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(customer)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(customer.GetId())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetId() != customer.GetId() {
				t.Errorf("Expected %v, got %v", customer.GetId(), found.GetId())
			}
		})
	}
}
