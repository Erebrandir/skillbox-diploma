package billing

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	name     string
	input    string
	expected BillingData
	ok       bool
}{
	{
		name:  "Тест Billing пример из ТЗ",
		input: "tests/specification.data",
		expected: BillingData{
			CreateCustomer: true,
			Purchase:       true,
			Payout:         false,
			Recurring:      false,
			FraudControl:   true,
			CheckoutPage:   true,
		},
		ok: true,
	},
	{
		name:  "Тест Billing из симулятора",
		input: "tests/simulator.data",
		expected: BillingData{
			CreateCustomer: true,
			Purchase:       true,
			Payout:         true,
			Recurring:      true,
			FraudControl:   false,
			CheckoutPage:   false,
		},
		ok: true,
	},
	{
		name:  "Тест Billing все",
		input: "tests/all.data",
		expected: BillingData{
			CreateCustomer: true,
			Purchase:       true,
			Payout:         true,
			Recurring:      true,
			FraudControl:   true,
			CheckoutPage:   true,
		},
		ok: true,
	},
	{
		name:     "Тест Billing err1",
		input:    "tests/err1.data",
		expected: BillingData{},
		ok:       false,
	},
	{
		name:     "Тест Billing err2",
		input:    "tests/err2.data",
		expected: BillingData{},
		ok:       false,
	},
}

func TestStatusBilling(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, ok := StatusBilling(test.input)
			require.Equal(t, test.expected, result)
			require.Equal(t, test.ok, ok)
		})
	}
}
