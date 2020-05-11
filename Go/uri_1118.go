package main

import (
	"fmt"
	"os"
)

func main() {
	var aux = 0
	var x, y, m = 0.0, 0.0, 0.0
	fmt.Scanf("%f", &x)
	for x < 0 || x > 10 {
		fmt.Println("nota invalida")
		fmt.Scanf("%f", &x)
	}
	fmt.Scanf("%f", &y)
	for y < 0 || y > 10 {
		fmt.Println("nota invalida")
		fmt.Scanf("%f", &y)
	}
	m = (x + y) / 2.0
	fmt.Printf("media = %.2f\n", m)

	for aux != 1 {
		fmt.Println("novo calculo (1-sim 2-nao)")
		fmt.Scanf("%d", &aux)
		if aux == 1 {
			main()
		} else if aux == 2 {
			os.Exit(0)
		}
	}
}
