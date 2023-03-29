package application

import (
	"go-ddd-tavern/internal/domain/repository"
	"go-ddd-tavern/internal/infrastructure/memory"
)

type OrderConfiguration func(app *OrderApp) error

type OrderApp struct {
	customers repository.CustomerRepository
}

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
