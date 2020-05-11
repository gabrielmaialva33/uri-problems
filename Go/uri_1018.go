package main

import "fmt"

func main() {
	var n, aux int
	fmt.Scan(&n)
	fmt.Printf("%d\n", n)
	aux = n / 100
	fmt.Printf("%d nota(s) de R$ 100,00\n", aux)
	n = n % 100
	aux = n / 50
	fmt.Printf("%d nota(s) de R$ 50,00\n", aux)
	n = n % 50
	aux = n / 20
	fmt.Printf("%d nota(s) de R$ 20,00\n", aux)
	n = n % 20
	aux = n / 10
	fmt.Printf("%d nota(s) de R$ 10,00\n", aux)
	n = n % 10
	aux = n / 5
	fmt.Printf("%d nota(s) de R$ 5,00\n", aux)
	n = n % 5
	aux = n / 2
	fmt.Printf("%d nota(s) de R$ 2,00\n", aux)
	n = n % 2
	aux = n / 1
	fmt.Printf("%d nota(s) de R$ 1,00\n", aux)
}
