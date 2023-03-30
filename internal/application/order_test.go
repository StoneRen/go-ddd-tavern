package application

import (
	"github.com/google/uuid"
	"go-ddd-tavern/internal/domain/entity"
	"testing"
)

func init_products(t *testing.T) []entity.Product {
	beer, err := entity.NewProduct("Beer", "Healthy Beverage", 1.99)
	peenuts, err := entity.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	wine, err := entity.NewProduct("Wine", "Healthy Snacks", 0.22)
	if err != nil {
		t.Error(err)
	}
	products := []entity.Product{
		beer, peenuts, wine,
	}
	return products
}

func TestOrder_NewOrderApp(t *testing.T) {
	products := init_products(t)

	app, err := NewOrderApp(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Error(err)
	}

	// 生成新用户
	customer, err := entity.NewCustomer("StoneRen")
	if err != nil {
		t.Error(err)
	}

	// 添加用户
	err = app.customers.Add(customer)
	if err != nil {
		t.Error(err)
	}

	buyProductList := []uuid.UUID{
		products[0].GetId(), products[2].GetId(),
	}

	price, err := app.CreateOrderSumPrice(customer.GetId(), buyProductList)
	if err != nil {
		t.Error(err)
	}
	exceptedPrice := 1.99 + 0.22
	if price != exceptedPrice {
		t.Errorf("Excepted price is %f, but got %f", exceptedPrice, price)
	}
}
