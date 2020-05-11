package main

import "fmt"

func main() {
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
}
