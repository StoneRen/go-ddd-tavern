package repository

import (
	"errors"
	"github.com/google/uuid"
	"go-ddd-tavern/internal/domain/entity"
)

var (
	ErrProductNotFound     = errors.New("the product was not found")
	ErrProductAlreadyExist = errors.New("the product already exists")
)

type ProductRepository interface {
	GetAll() ([]entity.Product, error)
	GetByID(id uuid.UUID) (entity.Product, error)
	Add(product entity.Product) error
	Update(product entity.Product) error
	Delete(id uuid.UUID) error
}
