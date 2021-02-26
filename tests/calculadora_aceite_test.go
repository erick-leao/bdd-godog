package calc_tests

import "github.com/cucumber/godog"

type CalculadoraSuite struct {
	ctx  *godog.ScenarioContext
	calc *Calculadora
}

func (c *CalculadoraSuite) iResultadoDeveSer(arg1 int) (int, error) {
	return c.calc.Resultado(), nil
}

func (c *CalculadoraSuite) iUsuarioPrecionaOValor(valor int) error {
	c.calc.Precionar(valor)
	return nil
}

func (c *CalculadoraSuite) iUsuarioSomaMaisOValor(valor int) error {
	c.calc.Somar(valor)
	return nil
}

func (c *CalculadoraSuite) iqueACalculadoraEstejaLimpo() error {
	c.calc.Clear()
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	calc := &CalculadoraSuite{
		ctx:  ctx,
		calc: new(Calculadora),
	}

	ctx.Step(`^o resultado deve ser (\d+)$`, calc.iResultadoDeveSer)
	ctx.Step(`^o Usuario preciona o valor (\d+)$`, calc.iUsuarioPrecionaOValor)
	ctx.Step(`^o Usuario soma mais o valor (\d+)$`, calc.iUsuarioSomaMaisOValor)
	ctx.Step(`^que a calculadora esteja limpo$`, calc.iqueACalculadoraEstejaLimpo)
}
