package main

import (
	"fmt"
	"math"
)

func main() {
	var raio float64
	fmt.Scan(&raio)
	fmt.Printf("VOLUME = %.3f\n", calcvolume(raio))
}
func calcvolume(x float64) float64 {
	return (4.0 / 3.0) * 3.14159 * math.Pow(x, 3)
}
