package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	v := [] float64{a, b, c}
	sort.Float64s(v)
	for i := 0; i < 3; i++ {
		v[i] = v[i]
	}
	// A = v[2] B = v[1] C[0]
	if v[2] >= v[1]+v[0] {
		fmt.Print("NAO FORMA TRIANGULO\n")
	} else {
		if math.Pow(v[2], 2) == (math.Pow(v[1], 2) + math.Pow(v[0], 2)) {
			fmt.Print("TRIANGULO RETANGULO\n")
		}
		if math.Pow(v[2], 2) > math.Pow(v[1], 2)+math.Pow(v[0], 2) {
			fmt.Print("TRIANGULO OBTUSANGULO\n")
		}
		if math.Pow(v[2], 2) < math.Pow(v[1], 2)+math.Pow(v[0], 2) {
			fmt.Print("TRIANGULO ACUTANGULO\n")
		}
		if v[2] == v[1] && v[1] == v[0] && v[0] == v[2] {
			fmt.Print("TRIANGULO EQUILATERO\n")
		} else if v[2] == v[1] || v[1] == v[0] || v[2] == v[0] {
			fmt.Print("TRIANGULO ISOSCELES\n")
		}
	}
}
