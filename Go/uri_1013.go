package main

import "fmt"
import "math"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Printf("%.0f eh o maior\n", math.Max(float64(a), math.Max(float64(b), float64(c))))
}
