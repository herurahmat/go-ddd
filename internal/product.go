package internal

// example value object

type Product struct {
	id        int64
	name      string
	skuNumber string
	vendor    string
	price     Price
}

func (p Product) ID() int64 {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Vendor() string {
	return p.vendor
}

func (p Product) SkuNumber() string {
	return p.skuNumber
}

func (p Product) ProductPrice() Price {
	return p.price
}

func (p *Product) SetPrice(base, strike float64) {
	newPrice := NewPrice(base, strike)
	p.price = newPrice
}

func NewProduct(id int64, name, skuNumber, vendor string) Product {

	return Product{
		id:        id,
		name:      name,
		skuNumber: skuNumber,
		vendor:    vendor,
	}
}
