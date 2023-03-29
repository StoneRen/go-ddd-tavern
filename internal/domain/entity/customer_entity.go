package entity

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("invalid person")
)

// Customer 聚合
// 可变的为指针类型
// 不可变的为值类型
type Customer struct {
	person       *Person
	products     []*Item
	transactions []Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &Person{
		Name: name,
		Id:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*Item, 0),
		transactions: make([]Transaction, 0),
	}, nil
}

func (c *Customer) GetId() uuid.UUID {
	return c.person.Id
}
func (c *Customer) SetId(id uuid.UUID) {
	if c.person == nil {
		c.person = &Person{}
	}
	c.person.Id = id
}

func (c *Customer) GetName() string {
	return c.person.Name
}
func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &Person{}
	}
	c.person.Name = name
}
