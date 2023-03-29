package repository

import (
	"errors"
	"github.com/google/uuid"
	"go-ddd-tavern/internal/domain/entity"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	ErrUpdateCustomer      = errors.New("failed to update the customer in the repository")
)

type CustomerRepository interface {
	Get(uuid uuid.UUID) (entity.Customer, error)
	Add(entity.Customer) error
	Update(entity.Customer) error
}
