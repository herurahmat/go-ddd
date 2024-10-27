package service

import "go-ddd/internal"

type GetProduct interface {
	GetProduct() []internal.Product
}
