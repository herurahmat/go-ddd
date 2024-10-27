package service

import (
	"go-ddd/internal"

	"github.com/google/uuid"
)

type GetCart interface {
	GetCart(id uuid.UUID) internal.Cart
}
