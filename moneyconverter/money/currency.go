package money

import "regexp"

// Currency defines the code of a currency and its decimal precision.
type Currency struct {
	code      string
	precision byte
}

// ErrInvalidCurrencyCode is returned when the currency to parse is not a standard 3-letter code.
const ErrInvalidCurrencyCode = Error("invalid currency code")

var codePattern = regexp.MustCompile("[A-Z][A-Z][A-Z]")

// ParseCurrency returns the currency associated to a name and may return errUnknownCurrency.
func ParseCurrency(code string) (Currency, error) {
	patternMatches := codePattern.MatchString(code)
	lengthMatches := len(code) == 3
	if !patternMatches || !lengthMatches {
		return Currency{}, ErrInvalidCurrencyCode
	}

	switch code {
	case "IRR":
		return Currency{code: code, precision: 0}, nil
	case "MGA", "MRU":
		return Currency{code: code, precision: 5}, nil
	case "CNY", "VND":
		return Currency{code: code, precision: 1}, nil
	case "BHD", "IQD", "KWD", "LYD", "OMR", "TND":
		return Currency{code: code, precision: 3}, nil
	default:
		// All other circulating currencies use a hundredth division.
		return Currency{code: code, precision: 2}, nil
	}
}

// String implements Stringer.
func (c Currency) String() string {
	return c.code
}
