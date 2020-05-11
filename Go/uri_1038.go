package main

import "fmt"

func main() {
	v := [5]float64{4.00, 4.50, 5.00, 2.00, 1.50}
	var cod, qtd int
	fmt.Scan(&cod, &qtd)
	fmt.Printf("Total: R$ %.2f\n", v[cod-1]*float64(qtd))
}
