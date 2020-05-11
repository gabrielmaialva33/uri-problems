package main

import "fmt"
import "math"

func main() {
	var raio float64
	fmt.Scan(&raio)
	fmt.Printf("VOLUME = %.3f\n", calc_volume(raio))
}
func calc_volume(x float64) float64 {
	return (4.0 / 3.0) * 3.14159 * math.Pow(x, 3)
}
