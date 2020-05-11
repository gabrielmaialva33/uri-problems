package main

import "fmt"

func main ()  {
	var cod, qtd int
	var price, total float64
	fmt.Scanln(&cod, &qtd, &price)
	total = float64(qtd) * price;
	fmt.Scanln(&cod, &qtd, &price)
	total += float64(qtd) * price
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n",total)
}