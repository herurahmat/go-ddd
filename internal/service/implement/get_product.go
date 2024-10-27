package implement

// this application service. because get data from external

import (
	"context"
	"go-ddd/internal"
	"go-ddd/internal/repository"
	"go-ddd/internal/service"
)

type GetProduct struct {
	productRepository repository.ProductModem
}

// GetProduct implements service.GetProduct.
func (g *GetProduct) GetProduct() []internal.Product {

	return g.productRepository.Get(context.Background())
}

func NewGetProduct() service.GetProduct {
	return &GetProduct{}
}
