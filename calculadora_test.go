package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeveSomar(t *testing.T) {
	calc := new(Calculadora)
	calc.Precionar(2)
	calc.Somar(2)

	assert.Equal(t, 4, calc.Resultado())
}
