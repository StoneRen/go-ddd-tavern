package memory

import (
	"github.com/google/uuid"
	"go-ddd-tavern/internal/domain/entity"
	"go-ddd-tavern/internal/domain/repository"
	"sync"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]entity.Product
	mu       sync.Mutex
}

func NewMemoryProductRepository() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]entity.Product),
	}
}

func (mpr *MemoryProductRepository) GetAll() ([]entity.Product, error) {
	// Collect all Products from map
	var products []entity.Product
	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}

func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (entity.Product, error) {
	if product, ok := mpr.products[uuid.UUID(id)]; ok {
		return product, nil
	}
	return entity.Product{}, repository.ErrProductNotFound
}
func (mpr *MemoryProductRepository) Add(newproduct entity.Product) error {
	mpr.mu.Lock()
	defer mpr.mu.Unlock()

	if _, ok := mpr.products[newproduct.GetId()]; ok {
		return repository.ErrProductAlreadyExist
	}

	mpr.products[newproduct.GetId()] = newproduct

	return nil
}
func (mpr *MemoryProductRepository) Update(upprod entity.Product) error {
	mpr.mu.Lock()
	defer mpr.mu.Unlock()

	if _, ok := mpr.products[upprod.GetId()]; !ok {
		return repository.ErrProductNotFound
	}

	mpr.products[upprod.GetId()] = upprod
	return nil
}

func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.mu.Lock()
	defer mpr.mu.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return repository.ErrProductNotFound
	}
	delete(mpr.products, id)
	return nil
}
