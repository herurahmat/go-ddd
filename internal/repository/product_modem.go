package repository

import (
	"context"
	"go-ddd/internal"
)

type ProductModem interface {
	Get(ctx context.Context) []internal.Product
}
