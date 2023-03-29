package application

import (
	"github.com/google/uuid"
	"go-ddd-tavern/internal/domain/repository"
	"go-ddd-tavern/internal/infrastructure/memory"
)

type OrderConfiguration func(app *OrderApp) error

type OrderApp struct {
	customers repository.CustomerRepository
}

/**
 * 创建和配置App
 */

func NewOrderApp(cfgs ...OrderConfiguration) (*OrderApp, error) {
	app := &OrderApp{}
	for _, cfg := range cfgs {
		if err := cfg(app); err != nil {
			return nil, err
		}
	}
	return app, nil
}

func WithCustomerRepository(cr repository.CustomerRepository) OrderConfiguration {
	return func(app *OrderApp) error {
		app.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.NewCustomerRepository()
	return WithCustomerRepository(cr)
}

/**
 * 以下是具体业务逻辑
 */

func (app *OrderApp) CreateOrder(customerId uuid.UUID, productIds []uuid.UUID) error {
	_, err := app.customers.Get(customerId)

	if err != nil {
		return err
	}
	return nil
}
