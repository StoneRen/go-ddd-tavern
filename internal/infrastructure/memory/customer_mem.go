package memory

import (
	"fmt"
	"github.com/google/uuid"
	"go-ddd-tavern/internal/domain/entity"
	"go-ddd-tavern/internal/domain/repository"
	"sync"
)

type MemoryCustomerRepository struct {
	customers map[uuid.UUID]entity.Customer
	mu        sync.Mutex
}

func NewCustomerRepository() *MemoryCustomerRepository {
	return &MemoryCustomerRepository{
		customers: make(map[uuid.UUID]entity.Customer),
	}
}

func (mr *MemoryCustomerRepository) Get(id uuid.UUID) (entity.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}

	return entity.Customer{}, repository.ErrCustomerNotFound
}

func (mr *MemoryCustomerRepository) Add(c entity.Customer) error {
	if mr.customers == nil {
		// Saftey check if customers is not create, shouldn't happen if using the Factory, but you never know
		mr.mu.Lock()
		mr.customers = make(map[uuid.UUID]entity.Customer)
		mr.mu.Unlock()
	}
	// Make sure Customer isn't already in the repository
	if _, ok := mr.customers[c.GetId()]; ok {
		return fmt.Errorf("customer already exists: %w", repository.ErrFailedToAddCustomer)
	}
	mr.mu.Lock()
	mr.customers[c.GetId()] = c
	mr.mu.Unlock()
	return nil
}

// Update will replace an existing customer information with the new customer information
func (mr *MemoryCustomerRepository) Update(c entity.Customer) error {
	if _, ok := mr.customers[c.GetId()]; !ok {
		return fmt.Errorf("customer does not exist: %w", repository.ErrUpdateCustomer)
	}
	mr.mu.Lock()
	mr.customers[c.GetId()] = c
	mr.mu.Unlock()
	return nil
	return nil
}
