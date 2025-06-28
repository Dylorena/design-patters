package calculadoradesconto

type DescontoQuantidade struct{}

func (dp DescontoQuantidade) Calcular(total float64, quantidade int) float64 {
	if quantidade >= 5 {
		return 20
	}

	return 0
}
