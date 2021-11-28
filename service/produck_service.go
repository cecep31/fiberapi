package service

import "github.com/cecep31/fiberapi/model"

type ProduckService interface {
	Create(request model.CreateProduckRequest) (response model.CreateProductResponse)
	List() (response []model.GetProductResponse)
}
