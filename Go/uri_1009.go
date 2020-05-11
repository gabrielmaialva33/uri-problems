package main

import "fmt"

func main() {
	var nome string
	var salario_fixo, vendas_mes float64
	fmt.Scan(&nome, &salario_fixo, &vendas_mes)
	fmt.Printf("TOTAL = R$ %.2f\n", salario_fixo+(vendas_mes*15/100.0))
}
