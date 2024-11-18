package util

const (
	USD = "USD"
	CAD = "CAD"
	EUR = "EUR"
	VND = "VND"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, CAD, EUR, VND:
		return true
	}
	return false
}
