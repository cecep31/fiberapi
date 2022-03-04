package repository

import (
	"github.com/cecep31/fiberapi/entity"
	"gorm.io/gorm"
)

type ProduckRepository interface {
	Insert(produck entity.Produck)
	FindAll() (produck []entity.Produck)
	DeleteAll()
}

func NewProduckRepository(database *gorm.DB) ProduckRepository {
	return &produckrepositoryimpl{
		Db: database,
	}
}

type produckrepositoryimpl struct {
	Db *gorm.DB
}

func (repository *produckrepositoryimpl) Insert(produck entity.Produck) {

}

func (repository *produckrepositoryimpl) FindAll() (produck []entity.Produck) {
	
}

func (repository *produckrepositoryimpl) DeleteAll() {

}
