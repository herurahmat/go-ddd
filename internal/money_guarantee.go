package internal

type MoneyGuarantee int

const (
	Undefined MoneyGuarantee = iota
	MONEY_GUARANTEED
	NOT_GUARANTEED
	NOT_OPTIMAL
	NONE
)

func (mg MoneyGuarantee) String() string {
	switch mg {
	case MONEY_GUARANTEED:
		return "Money Guaranteed"
	case NOT_GUARANTEED:
		return "Not Guaranteed"
	case NOT_OPTIMAL:
		return "Not Optimal"
	case NONE:
		return "None"
	default:
		return "Undefined"
	}
}
