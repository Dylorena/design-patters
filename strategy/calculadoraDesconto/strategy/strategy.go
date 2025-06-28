package calculadoradesconto

type ICalculadoraDesconto interface {
	Calcular(total float64, quantidade int) float64
}
