package main

import "fmt"

func main() {
	var v [6] float64
	var m, aux = 0.0, 0.0
	for i := 0; i < len(v); i++ {
		fmt.Scanln(&v[i])
		if v[i] > 0 {
			aux += 1
			m += v[i]
		}
	}
	fmt.Printf("%.0f valores positivos\n%.1f\n", aux, m/aux)
}
