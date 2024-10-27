package productorbit

import (
	"context"
	"go-ddd/internal"
	"go-ddd/internal/repository"
)

type ProductOrbit struct{}

func (p *ProductOrbit) Get(ctx context.Context) []internal.Product {
	/// you can put http call or anything here

	response := response{
		data: []responseDetail{
			{
				Id:                 1,
				Name:               "A",
				SkuNumber:          "B",
				Vendor:             "C",
				BasePrice:          100,
				StrikethroughPrice: 90,
			},
			{
				Id:                 2,
				Name:               "D",
				SkuNumber:          "E",
				Vendor:             "F",
				BasePrice:          200,
				StrikethroughPrice: 190,
			},
		},
	}

	products := []internal.Product{}

	for _, p := range response.data {

		newProduct := internal.NewProduct(
			p.Id,
			p.Name,
			p.SkuNumber,
			p.Vendor,
		)
		newProduct.SetPrice(p.BasePrice, p.StrikethroughPrice)

		products = append(products, newProduct)
	}

	return products
}

func NewProductOrbit() repository.ProductModem {
	return &ProductOrbit{}
}
