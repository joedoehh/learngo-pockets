package money

import (
	"testing"
)

func TestParseCurrency(t *testing.T) {
	tt := map[string]struct {
		in          string
		expected    Currency
		expectError bool
	}{
		"100thousandth MGA": {in: "MGA", expected: Currency{code: "MGA", precision: 5}, expectError: false},
		"thousandth BHD":    {in: "BHD", expected: Currency{code: "BHD", precision: 3}, expectError: false},
		"majority EUR":      {in: "EUR", expected: Currency{code: "EUR", precision: 2}, expectError: false},
		"tenth VND":         {in: "VND", expected: Currency{code: "VND", precision: 1}, expectError: false},
		"integer IRR":       {in: "IRR", expected: Currency{code: "IRR", precision: 0}, expectError: false},

		"Eur is wrong code":     {in: "Eur", expectError: true},
		"eur is wrong code":     {in: "eur", expectError: true},
		"INVALID is wrong code": {in: "INVALID", expectError: true},
		"EURA is wrong code":    {in: "EURA", expectError: true},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ParseCurrency(tc.in)

			if tc.expectError {
				if err == nil {
					t.Fatal("expected error, got nothing")
				}
				return
			}

			if err != nil {
				t.Errorf("expected no error, got %s", err.Error())
			}

			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
