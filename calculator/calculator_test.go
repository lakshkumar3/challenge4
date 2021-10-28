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
	testCase := TestCase{
		value:    "(1+3)/2",
		expected: "(1+3)/2 = 2.00",
	}
	var err error
	testCase.actual, err = calculator.CalculateExpression(testCase.value)
	if err != nil {
		t.Fail()
	}
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}
