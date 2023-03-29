package entity

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrMissingValues = errors.New("missing values")
)

type Product struct {
	item     *Item
	price    float64
	quantity int
}

func NewProduct(name string, desc string, price float64) (Product, error) {
	if name == "" || desc == "" || price == 0 {
		return Product{}, ErrMissingValues
	}

	item := &Item{
		Id:   uuid.New(),
		Name: name,
		Desc: desc,
	}

	return Product{
		item:     item,
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetId() uuid.UUID {
	return p.item.Id
}
func (p Product) GetItem() *Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}
