package main

import (
	"fmt"
)

func main() {
	var f float64
	var x int
	fmt.Scanf("%f", &f)
	f = f * 1000
	x = int(f)
	fmt.Print("NOTAS:\n")
	fmt.Printf("%d nota(s) de R$ 100.00\n", x/100000)
	x %= 100000
	fmt.Printf("%d nota(s) de R$ 50.00\n", x/50000)
	x %= 50000
	fmt.Printf("%d nota(s) de R$ 20.00\n", x/20000)
	x %= 20000
	fmt.Printf("%d nota(s) de R$ 10.00\n", x/10000)
	x %= 10000
	fmt.Printf("%d nota(s) de R$ 5.00\n", x/5000)
	x %= 5000
	fmt.Printf("%d nota(s) de R$ 2.00\n", x/2000)
	x %= 2000
	fmt.Print("MOEDAS:\n")
	fmt.Printf("%d moeda(s) de R$ 1.00\n", x/1000)
	x %= 1000
	fmt.Printf("%d moeda(s) de R$ 0.50\n", x/500)
	x %= 500
	fmt.Printf("%d moeda(s) de R$ 0.25\n", x/250)
	x %= 250
	fmt.Printf("%d moeda(s) de R$ 0.10\n", x/100)
	x %= 100
	fmt.Printf("%d moeda(s) de R$ 0.05\n", x/50)
	x %= 50
	fmt.Printf("%d moeda(s) de R$ 0.01\n", x/10)
}
