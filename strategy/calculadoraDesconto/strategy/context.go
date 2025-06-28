package calculadoradesconto

import "fmt"

type Context struct {
	strategia ICalculadoraDesconto
	logger    func(string)
}

func (cd *Context) Calcular(total float64, quantidade int) float64 {
	cd.logger("Iniciando cálculo de desconto...")
	return cd.strategia.Calcular(total, quantidade)
}

func NewCalculadora(tipoDesconto ICalculadoraDesconto) *Context {
	return &Context{
		strategia: tipoDesconto,
		logger:    func(msg string) { fmt.Println("[LOG]", msg) },
	}
}
