package application

import (
	"github.com/google/uuid"
	"log"
)

type TavernConfiguration func(t *Tavern) error

type Tavern struct {
	OrderApp   OrderApp
	BillingApp interface{}
}

/**
 * 创建和配置App
 */
func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	t := &Tavern{}
	for _, cfg := range cfgs {
		if err := cfg(t); err != nil {
			return nil, err
		}
	}
	return t, nil
}

func WithOrderApp(app *OrderApp) TavernConfiguration {
	return func(t *Tavern) error {
		t.OrderApp = *app
		return nil
	}
}

func (t *Tavern) CreateOrder(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderApp.CreateOrderSumPrice(customer, products)
	if err != nil {
		return err
	}
	log.Printf("Order price is %0.00f", price)

	return nil
}
