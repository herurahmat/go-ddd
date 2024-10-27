package internal

import "github.com/google/uuid"

type Cart struct {
	ID          uuid.UUID
	CartDetails CartDetails
}

func (c Cart) NewEmpty() Cart {
	return Cart{
		ID:          c.ID,
		CartDetails: CartDetails{},
	}
}

func (c Cart) SubTotal() float64 {
	return c.CartDetails.TotalPrice()
}

func (c *Cart) AddProduct(product Product, qty int, moneyGuarantee MoneyGuarantee) {
	id := uuid.New()
	c.CartDetails = append(c.CartDetails, NewCartDetail(id, product, qty, moneyGuarantee))
}

func (c *Cart) UpdateProduct(productId int64, quantity int, moneyGuarantee MoneyGuarantee) {

	if len(c.CartDetails) > 0 {
		return
	}

	for i, detail := range c.CartDetails {
		if detail.Product.ID() == productId {
			c.CartDetails[i].Qty = quantity
			c.CartDetails[i].MoneyGuarantee = moneyGuarantee
			return
		}
	}
}

func (c *Cart) RemoveProduct(productId int64) {
	for i, detail := range c.CartDetails {
		if detail.Product.ID() == productId {
			c.CartDetails = append(c.CartDetails[:i], c.CartDetails[i+1:]...)
			return
		}
	}
}

func NewCart(id uuid.UUID, cartDetails CartDetails) Cart {
	return Cart{
		ID:          id,
		CartDetails: cartDetails,
	}
}
