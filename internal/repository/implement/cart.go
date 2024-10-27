package implement

import (
	"context"
	"go-ddd/internal"
	"go-ddd/internal/repository"
	"time"

	"github.com/google/uuid"
)

type Cart struct{}

// GetCart implements repository.CartRepository.
func (c *Cart) GetCart(ctx context.Context, id uuid.UUID) internal.Cart {

	newProduct := internal.NewProduct(
		time.Now().Unix(),
		"name",
		"sku",
		"vendir",
	)
	newProduct.SetPrice(float64(100), float64(60))

	return internal.Cart{
		ID: id,
		CartDetails: []internal.CartDetail{
			{
				ID:             uuid.New(),
				Qty:            1,
				Product:        internal.Product{},
				MoneyGuarantee: internal.MONEY_GUARANTEED,
			},
		},
	}
}

func NewCart() repository.CartRepository {
	return &Cart{}
}
