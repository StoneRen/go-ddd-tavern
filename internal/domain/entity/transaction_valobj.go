package entity

import (
	"github.com/google/uuid"
	"time"
)

// Transaction 交易
// 所有的key都是小写，因为值对象是不可变的
type Transaction struct {
	amount    uint64
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
