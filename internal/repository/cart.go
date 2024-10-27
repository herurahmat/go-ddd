package repository

import (
	"context"
	"go-ddd/internal"

	"github.com/google/uuid"
)

type CartRepository interface {
	GetCart(ctx context.Context, id uuid.UUID) internal.Cart
}
