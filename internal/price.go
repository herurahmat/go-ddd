package internal

// example value object

type Price struct {
	baseValue     float64
	strikethrough float64
}

func (p Price) StrikethroughPrice() float64 {
	return p.strikethrough
}

func (p Price) BaseValue() float64 {
	return p.baseValue
}

func (p Price) Price() float64 {
	if p.baseValue < p.strikethrough {
		return p.baseValue
	}
	return p.strikethrough
}

func NewPrice(baseValue, strikethrough float64) Price {
	return Price{
		baseValue:     baseValue,
		strikethrough: strikethrough,
	}
}
