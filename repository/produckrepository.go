package repository

import "github.com/cecep31/fiberapi/entity"

type ProduckRepository interface {
	Inser(produck entity.Produck)
	FindAll() (produck []entity.Produck)
	DeleteAll()
}

type produckrepositoryimpl struct {
}
