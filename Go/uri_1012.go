package main

import "fmt"
import "math"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	fmt.Printf("TRIANGULO: %.3f\n"+
		"CIRCULO: %.3f\n"+
		"TRAPEZIO: %.3f\n"+
		"QUADRADO: %.3f\n"+
		"RETANGULO: %.3f\n", (a*c)/2.0, 3.14159*math.Pow(c, 2), (a+b)*c/2.0, b*b, a*b)
}
