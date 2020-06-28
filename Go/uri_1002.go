package main

import (
	"fmt"
	"math"
)

func main() {
	var area, raio float64
	fmt.Scanf("%f", &raio)
	area = 3.14159 * math.Pow(raio, 2)
	fmt.Printf("A=%.4f\n", area)
}
