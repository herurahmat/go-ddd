package implement

import (
	"context"
	"go-ddd/internal"
	"go-ddd/internal/repository"
	"go-ddd/internal/service"

	"github.com/google/uuid"
)

// domain service, because no external data required here

type GetCart struct {
	cartRepository repository.CartRepository
}

// GetCart implements service.GetCart.
func (g *GetCart) GetCart(id uuid.UUID) internal.Cart {

	data := g.cartRepository.GetCart(context.Background(), id)

	return data

}

func NewGetCart(
	cartRepository repository.CartRepository,
) service.GetCart {
	return &GetCart{cartRepository: cartRepository}
}
