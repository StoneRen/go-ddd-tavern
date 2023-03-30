package application

import (
	"github.com/google/uuid"
	"go-ddd-tavern/internal/domain/entity"
	"go-ddd-tavern/internal/domain/repository"
	"go-ddd-tavern/internal/infrastructure/memory"
	"log"
)

type OrderConfiguration func(app *OrderApp) error

type OrderApp struct {
	customers repository.CustomerRepository
	products  repository.ProductRepository
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

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.NewCustomerRepository()
	return func(app *OrderApp) error {
		app.customers = cr
		return nil
	}
}

func WithMongoCustomerRepository(conn string) OrderConfiguration {
	return func(app *OrderApp) error {
		// 示例代码，如果需要的时候开启
		//cr, err := mongo.New(context.Background(), connectionString)
		//if err != nil {
		//	return err
		//}
		//os.customers = cr
		return nil
	}
}

func WithMemoryProductRepository(products []entity.Product) OrderConfiguration {
	return func(app *OrderApp) error {
		productRepository := memory.NewMemoryProductRepository()

		for _, product := range products {
			err := productRepository.Add(product)
			if err != nil {
				return err
			}

		}
		app.products = productRepository
		return nil
	}
}

/**
 * 以下是具体业务逻辑
 */

func (app *OrderApp) CreateOrderSumPrice(customerId uuid.UUID, productIds []uuid.UUID) (float64, error) {
	customer, err := app.customers.Get(customerId)
	if err != nil {
		return 0, err
	}

	var products []entity.Product
	var price float64
	for _, id := range productIds {
		p, err := app.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}
	log.Printf("Customer: %s has ordered %d products ,All price : %f", customer.GetId(), len(products), price)

	return price, nil
}
