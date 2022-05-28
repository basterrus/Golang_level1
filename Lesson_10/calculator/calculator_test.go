//Реализация табличного теста на примере приложения калькулятор

package calculator

import (
	"testing"
)

type DataForCalculatorFunc struct {
	testName     string
	numOne       float64
	numTwo       float64
	operation    string
	resultOutput float64
}

func TestCalculator(t *testing.T) {
	data := []DataForCalculatorFunc{
		{
			testName:     "sum",
			numOne:       5.0,
			numTwo:       5.0,
			operation:    "+",
			resultOutput: 10.0,
		},
		{
			testName:     "multiplication",
			numOne:       5.0,
			numTwo:       5.0,
			operation:    "*",
			resultOutput: 25.0,
		},
		{
			testName:     "factorial",
			numOne:       5.0,
			numTwo:       5.0,
			operation:    "!",
			resultOutput: 120.0,
		},
	}
	for _, dt := range data {
		dt := dt
		t.Run(dt.testName, func(t *testing.T) {
			doResult := Calculator(dt.numOne, dt.numTwo, dt.operation)

			if doResult != dt.resultOutput {
				t.Errorf("Expected result %v != %v", doResult, dt.resultOutput)
			}
		})
	}
}
