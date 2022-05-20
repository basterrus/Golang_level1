package testify

import (
	"github.com/stretchr/testify/assert"
	"lesson_10/calculator"
	"testing"
)

func TestCalculator(t *testing.T) {

	assert.Equal(t, calculator.Calculator(5.0, 5.0, "*"), 25.0)

	if calculator.Calculator(5.0, 5.0, "*") != 25 {
		t.Errorf("5 умножить на 5 равно 25")
	}

	assert.Equal(t, calculator.Calculator(5.0, 5.0, "+"), 10.0)

	assert.Equal(t, calculator.Calculator(5.0, 5.0, "/"), 1.0)

}
