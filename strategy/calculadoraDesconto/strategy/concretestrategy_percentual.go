package calculadoradesconto

type DescontoPercentual struct{}

const percentual_desconto = 15

func (dp DescontoPercentual) Calcular(total float64, quantidade int) float64 {
	return total * (float64(percentual_desconto) / 100.0)
}
