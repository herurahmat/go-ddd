package productorbit

type response struct {
	data []responseDetail
}

type responseDetail struct {
	Id                 int64   `json:"id"`
	Name               string  `json:"name"`
	SkuNumber          string  `json:"skuNumber"`
	Vendor             string  `json:"vendor"`
	BasePrice          float64 `json:"basePrice"`
	StrikethroughPrice float64 `json:"strikethroughPrice"`
}
