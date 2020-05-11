package main

import "fmt"

func main() {
	var v [3] float64
	for i := 0; i < len(v); i++ {
		fmt.Scan(&v[i])
	}
	if v[0]+v[1] > v[2] && v[0]+v[2] > v[1] && v[1]+v[2] > v[0] {
		fmt.Printf("Perimetro = %.1f\n", v[0]+v[1]+v[2])
	} else {
		fmt.Printf("Area = %.1f\n", (v[0]+v[1])*v[2]/2.0)
	}

}
