package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c)
	d = math.Sqrt(math.Pow(b, 2) - 4*a*c)
	if d < 0 || a == 0 || math.IsNaN(d) == true {
		fmt.Printf("Impossivel calcular\n")
	} else {
		fmt.Printf("R1 = %.5f\nR2 = %.5f\n", (-b+d)/(2*a), (-b-d)/(2*a))
	}
}
