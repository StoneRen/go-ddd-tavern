package entity

import "github.com/google/uuid"

type Item struct {
	Id   uuid.UUID
	Name string
	Desc string
}
