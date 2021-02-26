package calc

// Calculadora Responsavel por calcular
type Calculadora struct {
	resultado int
}

// Somar soma com o valor informado
func (c *Calculadora) Somar(valor int) {
	c.resultado += valor
}

// Precionar seleciona um numero
func (c *Calculadora) Precionar(valor int) {
	c.resultado += valor
}

// Resultado retorna o resultado
func (c *Calculadora) Resultado() int {
	return c.resultado
}

// Clear limpa calculadora
func (c *Calculadora) Clear() {
	c.resultado = 0
}
