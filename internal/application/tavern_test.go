package application

import (
	"github.com/google/uuid"
	"go-ddd-tavern/internal/domain/entity"
	"testing"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	orderApp, err := NewOrderApp(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderApp(orderApp))
	if err != nil {
		t.Error(err)
	}

	customer, err := entity.NewCustomer("StoneRen")
	if err != nil {
		t.Error(err)
	}

	err = orderApp.customers.Add(customer)
	if err != nil {
		t.Error(err)
	}

	buyProductList := []uuid.UUID{
		products[0].GetId(), products[2].GetId(),
	}

	err = tavern.CreateOrder(customer.GetId(), buyProductList)
	if err != nil {
		t.Error(err)
	}

}
