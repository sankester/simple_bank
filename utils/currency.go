package utils

// constans to support all currency
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

func IsSupportedCurrecny(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}

	return false
}
