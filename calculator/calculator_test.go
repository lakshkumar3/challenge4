package calculator_test

import (
	"challenge1/calculator"
	"testing"
)

type TestCase struct {
	value    string
	expected string
	actual   string
}

func TestCalculateExpression(t *testing.T) {
	testCases := []TestCase{{
		value:    "(1+3)/20",
		expected: "(1+3)/20 = 0.20",
	}, {
		value:    "1/10",
		expected: "1/10 = 0.10",
	}, {
		value:    "(1-5)*(1+2)",
		expected: "(1-5)*(1+2) = -12.00",
	},
		{
			value:    "((1-5)*(1+2))/12",
			expected: "((1-5)*(1+2))/12 = -1.00",
		},
	}
	var err error

	for _, test := range testCases {
		test.actual, err = calculator.CalculateExpression(test.value)
		if err != nil {
			t.Fail()
		}

		if test.actual != test.expected {
			t.Fail()
		}
	}
}

func TestNegativeCalculateExpression(t *testing.T) {
	testCases := []TestCase{{
		value: "(1+3)//20",
	}, {
		value:    "1/10+1/0",
		expected: "1/10 = 0.10",
	}, {
		value: "(1-5)*(1+2)+",
	},
	}
	var err error

	for _, test := range testCases {
		test.actual, err = calculator.CalculateExpression(test.value)

		if err == nil {
			t.Fail()
		}

	}
}
