package calculadoradesconto

type DescontoFixo struct{}

func (df DescontoFixo) Calcular(total float64, quantidade int) float64 {
	return 10
}
