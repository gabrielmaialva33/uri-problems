package main

import "fmt"

func main() {
	var nome string
	var salariofixo, vendasmes float64
	fmt.Scan(&nome, &salariofixo, &vendasmes)
	fmt.Printf("TOTAL = R$ %.2f\n", salariofixo+(vendasmes*15/100.0))
}
