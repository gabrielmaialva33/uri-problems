package main

import (
	"fmt"
)

func main() {
	var n = 0
	fmt.Scanln(&n)
	m := make([][3]float64, n)
	var x = make([]float64, n)
	for i := 0; i < len(m); i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&m[i][j])
		}
	}
	for i := 0; i < n; i++ {
		x[i] = (m[i][0] * 2) + (m[i][1] * 3) + (m[i][2] * 5)
		fmt.Printf("%.1f\n", x[i]/10.0)
	}
}
