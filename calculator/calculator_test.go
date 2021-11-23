package calculator_test

import (
	"challenge1/calculator"
	"testing"
)

type TestCase struct {
	value  string
	result float64
	actual float64
}

func TestCalculateExpression(t *testing.T) {
	testCases := []TestCase{{
		value:  "(1+3)/20",
		result: 0.20,
	}, {
		value:  "1/10",
		result: 0.10,
	}, {
		value:  "(1-5)*(1+2)",
		result: -12.00,
	},
		{
			value:  "((1-5)*(1+2))/12",
			result: -1.00,
		},
	}
	var err error

	for _, test := range testCases {
		test.actual, err = calculator.CalculateExpression(test.value)
		if err != nil {
			t.Fail()
		}

		if test.actual != test.result {
			t.Fail()
		}
	}
}

func TestNegativeCalculateExpression(t *testing.T) {
	testCases := []TestCase{{
		value: "(1+3)//20",
	}, {
		value: "1/10+1/0",
	}, {
		value: "(1-5)*(1+2)+",
	},
	}
	var err error

	for _, test := range testCases {
		_, err = calculator.CalculateExpression(test.value)

		if err == nil {
			t.Fail()
		}

	}
}
