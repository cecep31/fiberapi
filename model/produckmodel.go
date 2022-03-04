package model

import "time"

type CreateProduckRequest struct {
	ID        int64     `json:"id"`
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	Quantity  int32     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateProductResponse struct {
	ID        int64     `json:"id"`
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	Quantity  int32     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetProductResponse struct {
	ID        int64     `json:"id"`
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	Quantity  int32     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
