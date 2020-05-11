package main

import "fmt"

func main() {
	var x, y float64
	fmt.Scan(&x, &y)
	fmt.Printf("%.3f km/l\n", x/y)
}
