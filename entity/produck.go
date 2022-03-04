package entity

import (
	"time"
)

type Produck struct {
	Id        int64
	Uuid      string
	Name      string
	Price     int64
	Quantity  int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
