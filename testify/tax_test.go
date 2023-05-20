package tax_testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	expected := 10.0
	actual := CalculateTax(1000)

	assert.Equal(t, expected, actual)
}
