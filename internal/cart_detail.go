package internal

import "github.com/google/uuid"

type CartDetails []CartDetail

type CartDetail struct {
	ID             uuid.UUID
	Qty            int
	Product        Product
	MoneyGuarantee MoneyGuarantee
}

func (c CartDetail) SubTotal() float64 {
	return float64(c.Product.ProductPrice().Price()) * float64(c.Qty)
}

func (c CartDetails) TotalPrice() float64 {

	totalPrice := float64(0)

	for _, detail := range c {
		totalPrice += detail.SubTotal()
	}

	return totalPrice
}

func NewCartDetail(id uuid.UUID, product Product, qty int, moneyGuarantee MoneyGuarantee) CartDetail {
	return CartDetail{
		ID:             id,
		Product:        product,
		Qty:            qty,
		MoneyGuarantee: moneyGuarantee,
	}
}
