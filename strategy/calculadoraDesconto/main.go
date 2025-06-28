package main

import (
	cd "calculadoraDesconto/strategy"
	"fmt"
)

func main() {
	calculadora := cd.NewCalculadora(cd.DescontoPercentual{})

	valor := calculadora.Calcular(300, 3)
	fmt.Printf("💸 Desconto aplicado: R$%.2f\n", valor)
}
